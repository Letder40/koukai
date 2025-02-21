#!/bin/bash

dependencies=(npm go)

dependency_check=true
for dep in ${dependencies[@]}; do
    which $dep &>/dev/null
    if [[ $? != 0 ]]; then
        echo "[!] $dep is a dependency required to deploy koukai"
        dependency_check=false
        exit
    fi
done

if [[ ! $dependency_check ]]; then
    exit
fi

if [[ ! -f "./strapi-backend/.env" ]]; then 
    read -p "Strapi panel addr - default [127.0.0.1]: " addr
    if [[ -z "$addr" ]]; then addr="127.0.0.1"; fi
    
    read -p "Strapi panel port - default [1337]: " port
    if [[ -z "$port" ]]; then port="1337"; fi

    appk1=$(openssl rand -base64 32)
    appk2=$(openssl rand -base64 32)
    api_token_salt=$(openssl rand -base64 32)
    admin_jwt_secret=$(openssl rand -base64 32)
    transfer_token_salt=$(openssl rand -base64 32)
    jwt_secret=$(openssl rand -base64 32)

    cat <<EOF > ./strapi-backend/.env
HOST=$addr
PORT=$port
APP_KEYS="$appk1,$appk2"
API_TOKEN_SALT=$api_token_salt
ADMIN_JWT_SECRET=$admin_jwt_secret
TRANSFER_TOKEN_SALT=$transfer_token_salt
JWT_SECRET=$jwt_secret
EOF

fi

read -p '[!] All nodejs process will be killed ctrl-c to abort, any to continue '
sudo pkill node

if [[ ! -f "./strapi-backend/package-lock.json" ]]; then
    echo "[#] Installing strapi dependencies"
    cd strapi-backend && npm install
fi

cd strapi-backend\
&& echo "[#] Initializing strapi..."\
&& npm run build &>/dev/null\
&& npm run start &>/dev/null & disown\
&& echo "[#] Strapi successfully started" 

if [[ ! -f "./frontend/package-lock.json" ]]; then
    echo "[#] Installing frontend dependencies"
    cd frontend && npm install
fi

cd frontend && BUILD_PATH=../backend/webfiles npm run build &>/dev/null\
&& echo "[#] Frontend built" 

cd ../backend\
&& go build\
&& ./koukai

npx tailwindcss -i ./frontend/src/css/index.css -o ./frontend/src/css/style.css || npm exec tailwindcss -i ./frontend/src/css/index.css -o ./frontend/src/css/style.css
