// +heroku goVersion go1.14
// +heroku install ./cmd/...

module github.com/fsufitch/tagialisi-bot

go 1.14

require (
	github.com/bwmarrin/discordgo v0.20.2
	github.com/golang-migrate/migrate/v4 v4.9.1
	github.com/google/wire v0.4.0
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/lib/pq v1.3.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.5.1
	github.com/urfave/cli/v2 v2.1.1
)
