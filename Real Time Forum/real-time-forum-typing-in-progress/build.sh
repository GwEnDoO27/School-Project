	#!/bin/bash
# Build Docker image
sudo docker build -t real-time-forum .
# Run Docker container
sudo docker run -p 8080:8080 real-time-forum
# Remove Docker image
sudo docker rmi -f real-time-forum