package response

var (
	OK                 = Response{StatusCode: 0, StatusMsg: "OK")}
	ErrLoginQuery      = Response{StatusCode: 101, StatusMsg: "Failed to query user login information"}
	ErrTokenGenerate   = Response{StatusCode: 102, StatusMsg: "Token generation failed"}
	ErrPassword        = Response{StatusCode: 103, StatusMsg: "Password error"}
	ErrPasswordLength  = Response{StatusCode: 104, StatusMsg: "Account name or password length is illegal"}
	ErrPasswordEncrypt = Response{StatusCode: 105, StatusMsg: "Password encryption error"}
	ErrUserCreation    = Response{StatusCode: 106, StatusMsg: "User creation failed"}
	ErrDuplicatedName  = Response{StatusCode: 107, StatusMsg: "User name already exists"}
	ErrParseInt        = Response{StatusCode: 108, StatusMsg: "Integer parsing failed"}
	ErrQueryUserInfo   = Response{StatusCode: 109, StatusMsg: "Failed to query user information"}
)