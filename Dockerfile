FROM golang:1.22.12-alpine3.20 AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /out/server \
    .

FROM alpine:3.20.6

RUN addgroup -S app && adduser -S -G app app

WORKDIR /app

COPY --from=build --chown=app:app /out/server ./server
COPY --from=build --chown=app:app /src/templates ./templates
COPY --from=build --chown=app:app /src/static ./static

ENV PORT=8989
EXPOSE 8989

USER app

CMD ["./server"]
