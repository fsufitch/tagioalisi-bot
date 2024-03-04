package security

import "github.com/google/wire"

// ProvideSecurity is a wire set containing code necessary to manage security
var ProvideSecurity = wire.NewSet(
	wire.Struct(new(JWTSupport), "*"),
	wire.Struct(new(AESSupport), "*"),
)
