package errno

// User
const (
	ServiceWrongPassword = 1000 + iota
	ServiceUserExist
	ServiceUserNotExist
	AddressNotExist

	ErrRecordNotFound
	UserLogOut
	UserBaned
	UserAlreadyLogin
)