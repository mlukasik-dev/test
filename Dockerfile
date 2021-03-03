FROM golang:1.16 AS build

WORKDIR /build

COPY . .

RUN make build

FROM debian:buster-slim

WORKDIR /app

COPY --from=build /build/main .

CMD [ "./main" ]