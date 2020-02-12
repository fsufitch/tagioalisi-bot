// +heroku goVersion go1.12
// +heroku install ./cmd/...

module github.com/fsufitch/discord-boar-bot

require (
	github.com/bwmarrin/discordgo v0.20.2
	github.com/docker/docker v0.7.3-0.20190817195342-4760db040282
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/golang-migrate/migrate/v4 v4.7.0
	github.com/google/wire v0.4.0
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.1
	github.com/lib/pq v1.1.0
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.3.0
	github.com/urfave/cli/v2 v2.0.0
)

go 1.13
