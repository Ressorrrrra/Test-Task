FROM golang:latest 

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
CMD ["/cmd/my-app/main"]
