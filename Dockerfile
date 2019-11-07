FROM alpine:latest

COPY api /usr/local/bin/api
CMD ["api"]
