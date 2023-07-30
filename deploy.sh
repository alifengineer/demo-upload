#!/bin/bash

docker stop $IMAGE_NAME || true

docker rm $IMAGE_NAME || true

docker run -d --name $IMAGE_NAME -p 8081:8081 ${{ secrets.DOCKERHUB_USERNAME }}/$IMAGE_NAME:${{ github.sha }}