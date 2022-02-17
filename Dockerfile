# env GOOS=linux GOARCH=amd64 go build
FROM alpine:latest
#WORKDIR /root/
#COPY --from=builder /app/client-go .
COPY ./client-go ./client-go
ENTRYPOINT [ "./client-go" ]