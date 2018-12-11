FROM golang:1.11.0-stretch as builder

WORKDIR /secretmessages

COPY . ./

# Building using -mod=vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -a -installsuffix cgo -ldflags="-w -s" -o application

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /secretmessages/application .
COPY --from=builder /secretmessages/internal/pkg/views/layouts internal/pkg/views/layouts
COPY --from=builder /secretmessages/internal/pkg/views/pages internal/pkg/views/pages

EXPOSE 5000

CMD ["./application"]