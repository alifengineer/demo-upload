docker pull ${{ secrets.DOCKERHUB_USERNAME }}/go-app:${{ github.sha }}

docker tag ${{ secrets.DOCKERHUB_USERNAME }}/go-app:${{ github.sha }} ${{ secrets.DOCKERHUB_USERNAME }}/go-app:latest

docker stop go-app || true

docker rm go-app || true

docker run -d --name go-app -p 8080:8080 ${{ secrets.DOCKERHUB_USERNAME }}/go-app:latest