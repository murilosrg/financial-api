FROM golang:1.15.2-alpine AS ci

RUN apk update && apk upgrade && apk add git gcc g++
RUN apk update && apk upgrade && apk add zip unzip openjdk8-jre git gcc musl-dev
RUN wget https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-4.4.0.2170-linux.zip
RUN unzip sonar-scanner-cli-4.4.0.2170-linux.zip
RUN mv ./sonar-scanner-4.4.0.2170-linux /opt/sonar-scanner
RUN echo "sonar.host.url=https://sonarcloud.io/" >> /opt/sonar-scanner/conf/sonar-scanner.properties
RUN sed -i 's/use_embedded_jre=true/use_embedded_jre=false/g' /opt/sonar-scanner/bin/sonar-scanner
ENV PATH=/opt/sonar-scanner/bin:$PATH

COPY . .

ENV GOPATH=/

RUN GO111MODULE=on go get github.com/golang/mock/mockgen@v1.4.4
RUN go generate ./...
RUN go test -v ./... -coverprofile=coverage.out
RUN sonar-scanner -D"sonar.login=$SONAR_TOKEN"

FROM golang:1.15-alpine AS build

WORKDIR /go/src/github.com/murilosrg/financial-api

COPY . .

RUN apk add --no-cache gcc musl-dev g++ \
    && go build ./cmd/financial && mv financial /go/bin

FROM golang:1.15-alpine AS runtime

COPY --from=build /go/bin/financial /usr/local/bin
COPY --from=build /go/src/github.com/murilosrg/financial-api /financial
COPY configuration.example.yml configuration.yml

WORKDIR /financial

CMD "financial" "-init"