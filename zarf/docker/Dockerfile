# Build the Go Binary
FROM golang:1.17 as build_basic_service
ENV CGO_ENABLED 0
ARG BUILD_REF

COPY . /service

WORKDIR /service
RUN go build -ldflags "-X main.build=${BUILD_REF}" -o ./service


# Run the Go Binary
FROM alpine:3.15

ARG BUILD_DATE
ARG BUILD_REF

COPY --from=build_basic_service /service /service

WORKDIR /service
CMD ["./service"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="service" \
      org.opencontainers.image.authors="Jan Duursma <me@janduursma.com>" \
      org.opencontainers.image.revision="${BUILD_REF}"