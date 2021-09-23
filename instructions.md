docker build -f .\cmd\client\Dockerfile -t websocket-client .
docker run -it websocket-client

docker build -f .\cmd\server\Dockerfile -t websocket-server .
docker run -p 8000:8000 -it websocket-server

