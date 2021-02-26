FROM golang:1.16 as build

WORKDIR /app

COPY . .

RUN make build

FROM ubuntu:20.04

WORKDIR /app

COPY --from=build /app/main .

EXPOSE 8000

CMD [ "./main" ]