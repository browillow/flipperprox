# 1) Build the go web server

FROM golang:1.9-alpine as build-go
WORKDIR /go/src/browillow/flipper/flipperprox
COPY . /go/src/browillow/flipper/flipperprox
RUN go install

# 2) Build the final image

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app/server/
COPY --from=build-go /go/bin /app/server
ENTRYPOINT /app/server/flipperprox
EXPOSE 80