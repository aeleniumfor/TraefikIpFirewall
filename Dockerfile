FROM golang:1.19-alpine3.18 AS build

WORKDIR /go/src/github.com/TraefikIpFirewall/
ADD ./ ./
RUN go build -o ./main ./main.go 


FROM alpine:latest
COPY --from=build /go/src/github.com/TraefikIpFirewall/main /app/main
CMD [ "/app/main" ]