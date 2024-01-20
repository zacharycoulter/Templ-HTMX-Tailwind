FROM alpine:3.19.0 AS generate
WORKDIR /app
COPY . .
RUN ./bin/tailwindcss/tailwindcss_linux_arm64 -o ./assets/global/style.css --minify && \
    ./bin/templ/templ_linux_arm64 generate

FROM golang:1.21.5-alpine3.19 AS build
WORKDIR /app
COPY --from=generate /app .
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x && \
    CGO_ENABLED=0 GOOS=linux go build -o server

FROM alpine:3.19.0 as final
WORKDIR /app
COPY --from=generate /app/assets assets/
COPY --from=build /app/server server
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT [ "./server" ]
