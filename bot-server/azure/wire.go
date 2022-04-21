package azure

import "github.com/google/wire"

// AzureProviderSet is a Wire provider set for Azure functionality
var AzureProviderSet = wire.NewSet(
	ProvideOnlineNewsSearch,
	wire.Bind(new(NewsSearch), new(*OnlineNewsSearch)),
)
