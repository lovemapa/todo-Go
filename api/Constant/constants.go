package constant

var Database = "todoGin"

const (
	SuccessCode                 = 200
	SuccessFlag                 = 1
	SuccessMsg                  = "Success"
	FailureCode                 = 400
	FailureFlag                 = 0
	Response                    = "data"
	UnauthorizedCode            = 401
	UnauthorizedStatus          = 2
	UnauthorizedMsg             = "Unauthorized"
	FailureMsg                  = "Failure"
	UsersCollection             = "users"
	CreatedSuccssfully          = "Account created succssfully"
	AdvertiseCreatedSuccssfully = "advertisement created succssfully"
	TodoCreatedSuccess          = "Todo created Successfully"
	ListFetchedSuccess          = "List fetched successfully"
	DeletedSuccessfully         = "Deleted Successfully"
	AccountAlreadyExists        = "Account already exists"
	UpdatedSuccessfully         = "Upated Successfully"
	EmptyReqBody                = "Request Body empty"
	FileMissing                 = "File is missing"
	LoggedInSuccess             = "Logged in successfully"
	AccountNotExists            = "Account doesn't exists"
	IncorrectPassword           = "Password incorrect"
)
