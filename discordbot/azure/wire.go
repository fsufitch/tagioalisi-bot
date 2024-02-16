package azure

import "github.com/google/wire"

// AzureProviderSet is a Wire provider set for Azure functionality
var AzureProviderSet = wire.NewSet(
	wire.Struct(new(BingNewsSearch), "*"),
)
