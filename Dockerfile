FROM alpine:latest

COPY build/buddy_bot /

ENTRYPOINT ["/buddy_bot"]