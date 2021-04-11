FROM golang:latest

ENV GO111MODULE=on

WORKDIR /app

COPY ./go.mod .

RUN go mod download

COPY . .

# Build the Go app
RUN go install github.com/applinh/kaepora

EXPOSE 5000

CMD ["kaepora", "server"]                                                                        
