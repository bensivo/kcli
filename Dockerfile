FROM golang:1.20.2-alpine3.17 as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./internal ./internal
COPY ./cmd ./cmd
COPY ./*.go ./

RUN go build -o /app/kcli

FROM alpine:3.17
WORKDIR /app
COPY --from=build /app/kcli /app/kcli

ENTRYPOINT [ "./kcli" ] 
