FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY build/buddy_bot /buddy_bot
COPY resources /resources

RUN chmod +x /buddy_bot

ENTRYPOINT ["/buddy_bot"]