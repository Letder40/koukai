package routing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"koukai/requests"
	"koukai/utils"
	"net/http"
	"strings"
	"sync"
)

// Message struct for reading messages
type Message struct {
	DocumentId string   `json:"documentId"`
	Body       string   `json:"body"`
	UserData   UserData `json:"sent_by"`
}

// Structs for saving messages
type connection struct {
	UserDocumentId string `json:"connect"`
}

type messageData struct {
	Body   string     `json:"body"`
	SentBy connection `json:"sent_by"`
}

type payload struct {
	Data messageData `json:"data"`
}

// Converts a Message to Strapi-compatible payload
func payloadFromMessage(message Message) payload {
	return payload{
		Data: messageData{
			Body:   message.Body,
			SentBy: connection{UserDocumentId: message.UserData.DocumentId},
		},
	}
}

// Server struct
type Server struct {
	DocumentId string    `json:"documentId"`
	Name       string    `json:"name"`
	Messages   []Message `json:"messages"`
	clients    map[http.ResponseWriter]bool
	mu         sync.Mutex
}

// Initialize global server
func (server *Server) InitGlobalServer() error {
	type ServerResponse struct {
		Data []Server `json:"data"`
	}

	uri := "/api/servers?populate[0]=messages&populate[1]=messages.sent_by&filter[name]=public"
	serverData, err := requests.ServerStrapiRequest("GET", uri, nil)
	if err != nil {
		fmt.Println("Error fetching server data:", err)
		return err
	}

	var resp ServerResponse
	if err := json.Unmarshal([]byte(serverData), &resp); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return err
	}

	if len(resp.Data) > 0 {
		*server = resp.Data[0]
	} else {
		fmt.Println("No server data found")
		return fmt.Errorf("no server data available")
	}

	server.clients = make(map[http.ResponseWriter]bool)
	fmt.Println("Initialized server:", server.DocumentId)
	return nil
}

// Save a message and broadcast it
func (server *Server) saveMessage(message *Message, userData UserData) {
	message.UserData = userData
	payload := payloadFromMessage(*message)

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling payload:", err)
		return
	}

	respBody, err := requests.ServerStrapiRequest("POST", "/api/messages", bytes.NewReader(payloadJson))
	if err != nil {
		fmt.Println("Error creating message:", err)
		return
	}

	type strapiResponse struct {
		Data struct {
			DocumentId string `json:"documentId"`
		} `json:"data"`
	}

	var response strapiResponse
	if err := json.Unmarshal([]byte(respBody), &response); err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	if response.Data.DocumentId == "" {
		fmt.Println("No documentId returned from Strapi")
		return
	}

	uri := fmt.Sprintf("/api/servers/%s", server.DocumentId)
	jsonStr := fmt.Sprintf(`{"data":{"messages":{"connect":["%s"]}}}`, response.Data.DocumentId)

	_, err = requests.ServerStrapiRequest("PUT", uri, strings.NewReader(jsonStr))
	if err != nil {
		fmt.Println("Error updating server:", err)
	}
}

// Broadcast message to all clients
func (server *Server) BroadcastMessage(message Message, userData UserData) {
	server.mu.Lock()
	defer server.mu.Unlock()

    message.UserData = userData
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error marshaling message:", err)
		return
	}

	for client := range server.clients {
		fmt.Fprintf(client, "data: %s\n\n", jsonData)
		client.(http.Flusher).Flush()
	}
}

// Retrieve user data from token
func getUserData(w http.ResponseWriter, jwt string) (UserData, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://127.0.0.1:8000/api/user", nil)
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return UserData{}, err
	}

	req.Header.Set("Authorization", "Bearer "+jwt)

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Request failed", http.StatusInternalServerError)
		return UserData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		http.Error(w, "Server error: "+string(body), http.StatusInternalServerError)
		return UserData{}, fmt.Errorf("server error: %s", string(body))
	}

	var userData UserData
	if err := json.NewDecoder(resp.Body).Decode(&userData); err != nil {
		http.Error(w, "Failed to parse user data", http.StatusInternalServerError)
		return UserData{}, err
	}

	return userData, nil
}

// Handle posting a message
func (server *Server) HandlePostMessage(w http.ResponseWriter, r *http.Request) {
	jwt, err := requests.GetUserJWT(r)
	if err != nil {
		http.Error(w, "Missing authentication token", http.StatusForbidden)
		return
	}

	userData, err := getUserData(w, jwt)
	if err != nil {
		return
	}

	var message Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	server.saveMessage(&message, userData)

	server.BroadcastMessage(message, userData)

	w.Header().Set("Content-Type", "application/json")
	utils.SendJsonFrom(w, message)
}

// Handle SSE client connections
func (server *Server) ListenChannel(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	_, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	// Register client
	server.mu.Lock()
	server.clients[w] = true
	server.mu.Unlock()

	// Keep connection open until client disconnects
	<-r.Context().Done()

	// Remove client from the list
	server.mu.Lock()
	delete(server.clients, w)
	server.mu.Unlock()
}
