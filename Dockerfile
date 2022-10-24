FROM golang:1-alpine AS builder
COPY Linux/tinyurl .
CMD ["./tinyurl"]