cd artifacts

docker-compose down

sleep 2

docker volume prune

sleep 2

docker network prune

sleep 7

systemctl restart docker