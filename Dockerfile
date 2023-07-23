FROM golang:1.20

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app stddevapi/cmd

EXPOSE 8080

ENV RANDOMORG_API_KEY="098c2e44-7ac0-4258-b8f6-9a117039e177"

CMD ["app"]
