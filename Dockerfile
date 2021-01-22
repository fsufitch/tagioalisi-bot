FROM golang:1.13 AS deps

WORKDIR /bot
ADD ./go.mod /bot
ADD ./go.sum /bot
RUN go mod download

#####
FROM deps

ADD . /bot
RUN go build -v ./cmd/tagi-bot
RUN go build -v ./cmd/tagi-migrate

CMD [ "./tagi-bot" ]