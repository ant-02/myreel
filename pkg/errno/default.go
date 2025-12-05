package errno

var (
	Success = NewErrNo(SuccessCode, "ok")

	UserNotFound = NewErrNo(ErrRecordNotFound, "user not found")
	UserIsBaned  = NewErrNo(UserBaned, "user is baned")

	LikeNotFound = NewErrNo(ErrRecordNotFound, "like not found")

	ParamVerifyError  = NewErrNo(ParamVerifyErrorCode, "parameter validation failed")
	ParamMissingError = NewErrNo(ParamMissingErrorCode, "missing parameter")

	AuthInvalid             = NewErrNo(AuthInvalidCode, "authentication failure")
	AuthAccessExpired       = NewErrNo(AuthAccessExpiredCode, "token expiration")
	AuthNoToken             = NewErrNo(AuthNoTokenCode, "lack of token")
	AuthNoOperatePermission = NewErrNo(AuthNoOperatePermissionCode, "No permission to operate")

	InternalServiceError = NewErrNo(InternalServiceErrorCode, "internal server error")
	OSOperationError     = NewErrNo(OSOperateErrorCode, "os operation failed")
	IOOperationError     = NewErrNo(IOOperateErrorCode, "io operation failed")

	UpYunFileError = NewErrNo(UpYunFileErrorCode, "upyun operation failed")
)
