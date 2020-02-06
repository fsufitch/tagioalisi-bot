FROM golang:1.13 AS deps

WORKDIR /bot
ADD ./go.mod /bot
ADD ./go.sum /bot
RUN go mod download

#####
FROM deps

ADD . /bot
RUN go build ./cmd/discord-boar-bot
RUN go build ./cmd/boarbot-migrate

CMD [ "./discord-boar-bot" ]