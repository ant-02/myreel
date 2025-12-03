package constants

// Service Name
const (
	GatewayServiceName = "gateway"
	UserServiceName    = "user"
	VideoServiceName   = "video"
)

// UserService
const (
	UserMaximumPasswordLength      = 72 // DO NOT EDIT (ref: bcrypt.GenerateFromPassword)
	UserMinimumPasswordLength      = 5
	UserDefaultEncryptPasswordCost = 10
)
