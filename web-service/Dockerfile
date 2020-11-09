# Compile golang app
FROM golang:1.15-alpine as build_base
RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /app/server
COPY . .
ENV GO11MODULE=on
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server

# Use Compiled app 
FROM alpine
RUN apk add --no-cache ca-certificates apache2-utils tzdata gettext
WORKDIR /app/cmd/app
COPY --from=build_base /server server

CMD ./server
