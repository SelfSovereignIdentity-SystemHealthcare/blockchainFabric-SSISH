#!/usr/bin/env bash

# Script used to stop the development environment.

# Stop API containers
cd ./ccapi; sudo docker-compose down; cd ..

# Stop fabric containers
cd ./fabric; sudo ./stopFabric.sh; cd ..

# Remove chaincode vendor folder if not needed
if [ -d "chaincode/vendor" ]; then
    sudo rm -rf ./chaincode/vendor
fi

# Remove any stopped containers and unused images and volumes
sudo docker container prune -f
sudo docker image prune -f
sudo docker volume prune -f

echo 'Development environment stopped successfully.'
