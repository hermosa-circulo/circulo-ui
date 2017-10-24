sudo apt-get install -y nodejs npm

npm cache clean
n stable

sudo ln -sf /usr/local/bin/node /usr/bin/node

node -v

apt-get purge -y nodejs npm

sudo ln -sf /usr/local/bin/npm /usr/bin/npm

cd ui-ng
npm install

cd ..
docker-compose up -d

/resources/node_modules/.bin/ng serve --host 0.0.0.0  --port 80 --open --public [hostname]
