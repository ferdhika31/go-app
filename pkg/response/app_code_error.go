package response

const (
	Unauthorized       = "ER00001"
	FailedConnectDb    = "ER00002"
	SessionExpired     = "ER00003"
	InvalidCredentials = "ER00004"
	LockedAccount      = "ER00005"
	ValidationFailed   = "ER00006"
	LogoutFailed       = "ER00007"
	RouteNotFound      = "ER00008"
	DuplicateData      = "ER00009"
	DataNotFound       = "ER00010"
	FailedFetchData    = "ER00011"
)

var appResponsesError = map[string]AppResponse{
	Unauthorized: {
		HttpStatus: StatusUnauthorized,
		Messages:   "You don't have access permission",
		Code:       Unauthorized,
	},
	FailedConnectDb: {
		HttpStatus: StatusBadRequest,
		Messages:   "Failed to connect to the database, contact the system administrator.",
		Code:       FailedConnectDb,
	},
	SessionExpired: {
		HttpStatus: StatusUnauthorized,
		Messages:   "The session expired, login to the application again.",
		Code:       SessionExpired,
	},
	InvalidCredentials: {
		HttpStatus: StatusUnauthorized,
		Messages:   "Enter valid credentials.",
		Code:       InvalidCredentials,
	},
	LockedAccount: {
		HttpStatus: StatusForbidden,
		Messages:   "If your account is locked, contact your administrator.",
		Code:       LockedAccount,
	},
	ValidationFailed: {
		HttpStatus: StatusUnprocessableEntity,
		Messages:   "Validation Failed.",
		Code:       ValidationFailed,
	},
	LogoutFailed: {
		HttpStatus: StatusBadRequest,
		Messages:   "Logout failed, try again.",
		Code:       LogoutFailed,
	},
	RouteNotFound: {
		HttpStatus: StatusNotFound,
		Messages:   "Route Not Found",
		Code:       RouteNotFound,
	},
	DuplicateData: {
		HttpStatus: StatusConflict,
		Messages:   "Data already exists.",
		Code:       DuplicateData,
	},
	DataNotFound: {
		HttpStatus: StatusNotFound,
		Messages:   "Data does not exist.",
		Code:       DataNotFound,
	},
	FailedFetchData: {
		HttpStatus: StatusBadRequest,
		Messages:   "Failed to fetch data.",
		Code:       FailedFetchData,
	},
}
