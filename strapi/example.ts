const element: HTMLElement = document.getElementById("example")!;
const new_element = document.createElement("p");
new_element.className = "myelementclass";
new_element.id = "myelementid";
new_element.innerHTML = "Esto es una prueba"
element.appendChild(new_element);
