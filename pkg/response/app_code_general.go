package response

const (
	InternalServerError = "GER00001"
	BadRequest          = "GER00002"
	MethodNotAllowed    = "GER00003"
)

var appResponsesGeneral = map[string]AppResponse{
	InternalServerError: {
		HttpStatus: StatusInternalServerError,
		Messages:   StatusText(StatusInternalServerError),
		Code:       InternalServerError,
	},
	BadRequest: {
		HttpStatus: StatusBadRequest,
		Messages:   StatusText(StatusBadRequest),
		Code:       BadRequest,
	},
	MethodNotAllowed: {
		HttpStatus: StatusMethodNotAllowed,
		Messages:   StatusText(StatusMethodNotAllowed),
		Code:       MethodNotAllowed,
	},
}
