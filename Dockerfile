FROM golang:1.21-alpine AS builder
WORKDIR /build
RUN apk add gcc g++ --no-cache
COPY . .
RUN cd /build/wireguard && go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o app -a -ldflags '-w -extldflags "-static"'  /build/wireguard/cmd/main.go

FROM alpine
WORKDIR /app
RUN apk update && apk add sudo && apk add iptables && apk add -U wireguard-tools
RUN echo "net.ipv4.ip_forward=1" >> /etc/sysctl.conf
RUN echo "net.ipv4.conf.all.src_valid_mark=1" >> /etc/sysctl.conf
COPY --from=builder /build/app /app/app
ENTRYPOINT ["sysctl -p && /app/app"]
EXPOSE 5454
EXPOSE 51820