cd artifacts/channel

./create-artifacts.sh

cd ..

docker-compose up -d

cd ..

./createChannel.sh

sleep 3

./deployChaincode.sh

sleep 2

cd api

node app.js