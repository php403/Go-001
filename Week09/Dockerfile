FROM  golang:latest
WORKDIR /home/pg
COPY . .
RUN go build .
EXPOSE 8888

RUN go run main.go

