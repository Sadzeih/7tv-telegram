FROM golang:buster as build

WORKDIR /app

COPY . .
RUN go build -o bot

FROM gcr.io/distroless/base

COPY --from=build /app/bot /

ENTRYPOINT ["/bot"]
