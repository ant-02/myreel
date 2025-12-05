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

// VideoService
const (
	VideoLikeGravity = 2
	VideoCommentGarvity = 3
	VideoVisitGarvity = 1
	VideoCreatedAtGarvity = 1.8
)
