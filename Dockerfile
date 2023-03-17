FROM golang:1.17.3-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./internal ./internal
COPY ./cmd ./cmd
COPY ./*.go ./

RUN go build

ENTRYPOINT [ "./kcli" ] 
