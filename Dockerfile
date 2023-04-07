FROM golang:1.18

RUN mkdir /app   
COPY . /app   
WORKDIR /app   
RUN go get ./ && go build && go mod download
EXPOSE 8080

CMD ["go", "run", "main.go"]