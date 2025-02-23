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

echo $appk1 > backend/.appkey

fi

read -p '[!] All nodejs process will be killed ctrl-c to abort, any to continue '
sudo pkill node

cd strapi-backend 
if [[ ! -d "node_modules" ]]; then
    echo "[#] Installing strapi dependencies"
    npm ci &>/dev/null
fi

echo "[#] Initializing strapi..."
npm run build &>/dev/null
npm run start &>/dev/null & disown

cd ../frontend
if [[ ! -d "node_modules" ]]; then
    echo "[#] Installing frontend dependencies"
    npm ci &>/dev/null
fi

echo "[#] Transpiling tailwind to css"
npx tailwindcss -i ./src/css/index.css -o ./src/css/style.css &>/dev/null

echo "[#] Building frontend"
BUILD_PATH=../backend/webfiles npm run build &>/dev/null

if [[ $? != 0 ]]; then
    echo "[!] Frontend error" 
    exit
fi

echo "[#] Building backend"
cd ../backend
go build

if [[ $? != 0 ]]; then
    echo "[!] Backend error" 
    exit
fi

echo -e "[#] Starting HTTP server\n"

./koukai
