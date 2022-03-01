# Build the Go Binary
FROM golang:1.17 as build_rumor-api
ENV CGO_ENABLED 0
ARG BUILD_REF

COPY . /service

WORKDIR /service/app/services/rumor-api
RUN go build -ldflags "-X main.build=${BUILD_REF}" -o ./rumor-api


# Run the Go Binary
FROM alpine:3.15

ARG BUILD_DATE
ARG BUILD_REF

COPY --from=build_rumor-api /service/app/services/rumor-api/rumor-api /service/rumor-api

WORKDIR /service
CMD ["./rumor-api"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="rumor-api" \
      org.opencontainers.image.authors="Jan Duursma <me@janduursma.com>" \
      org.opencontainers.image.source="https://github.com/janduursma/rumor-api/app/services/rumor-api" \
      org.opencontainers.image.revision="${BUILD_REF}" \
      org.opencontainers.image.vendor="Jan Duursma"