// +heroku goVersion go1.12
// +heroku install ./cmd/...

module github.com/fsufitch/discord-boar-bot

require (
	github.com/bwmarrin/discordgo v0.20.2
	github.com/golang-migrate/migrate/v4 v4.7.0
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/google/go-cmp v0.3.0 // indirect
	github.com/google/wire v0.4.0
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.1
	github.com/lib/pq v1.1.0
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.3.0
	github.com/urfave/cli/v2 v2.0.0
	golang.org/x/net v0.0.0-20190613194153-d28f0bde5980 // indirect
	golang.org/x/sys v0.0.0-20190801041406-cbf593c0f2f3 // indirect
)

go 1.13
