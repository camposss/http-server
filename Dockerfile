FROM golang:alpine as builder

RUN mkdir /app
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM scratch

COPY --from=builder /app/main /app/main

ENTRYPOINT [ "app/main" ]