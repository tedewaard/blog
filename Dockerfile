FROM golang:1.26.1-alpine AS builder

RUN apk add --no-cache git make curl libstdc++ libgcc

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/a-h/templ/cmd/templ@v0.3.1001

COPY . .

# Download Tailwind CSS standalone CLI (after COPY to avoid being overwritten)
# Needed the musl version because alpine doesn't have glibc
RUN curl -L https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64-musl -o tailwindcss \
    && chmod +x tailwindcss

RUN make all

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/server .

COPY --from=builder /app/src/static ./src/static
COPY --from=builder /app/content ./content

ENV PORT=3000
ENV ENVIRONMENT=production
ENV CONTENT_DIR=./content/blog
ENV STATIC_DIR=./src/static

EXPOSE 3000

CMD ["./server"]
