#Go API
FROM golang:latest AS builder
ADD . /app
WORKDIR /app/server
RUN go mod download
RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -ldflags "-w" -a -o /main .

#Build React application
FROM node:alpine AS node_builder
COPY --from=builder /app/client ./
RUN npm install -g npm
RUN npm run build

#Final Stage, this will be container
FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /main ./server/
COPY --from=node_builder /build ./client/build
WORKDIR /server
ENV PORT=8080

RUN chmod +x ./main
EXPOSE 8080

CMD ./main