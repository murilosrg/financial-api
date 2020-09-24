FROM golang:1.15-alpine AS build

WORKDIR /go/src/github.com/murilosrg/financial-api

ENV CC=gcc
ENV GIN_MODE=release

COPY . .

RUN apk add --no-cache gcc musl-dev \
    && go build ./cmd/financial && mv financial /go/bin

FROM golang:1.15-alpine AS runtime

COPY --from=build /go/bin/financial /usr/local/bin
COPY --from=build /go/src/github.com/murilosrg/financial-api /financial
COPY configuration.example.yml configuration.yml

WORKDIR /financial

CMD "financial" "-init"