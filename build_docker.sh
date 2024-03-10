docker build -t dt/vpn/bot/app:0.0.1 .
#apple silicon for other platform
#docker buildx build --platform=linux/amd64 -t dt/vpn/bot/app:0.0.1 .

#run it
#docker run -d -p 8080:8080 dt/vpn/bot/app:0.0.1
#and check localhost:8080