// +heroku goVersion go1.14
// +heroku install ./cmd/...

module github.com/fsufitch/tagioalisi-bot

go 1.14

require (
	cgt.name/pkg/go-mwclient v1.0.3
	github.com/antonholmquist/jason v1.0.0
	github.com/bwmarrin/discordgo v0.20.2
	github.com/golang-migrate/migrate/v4 v4.9.1
	github.com/golang/protobuf v1.3.4 // indirect
	github.com/google/uuid v1.1.1
	github.com/google/wire v0.4.0
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/lib/pq v1.3.0
	github.com/mrjones/oauth v0.0.0-20190623134757-126b35219450 // indirect
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.5.1
	github.com/urfave/cli/v2 v2.2.0
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
)
