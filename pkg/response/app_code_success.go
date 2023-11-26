package response

const (
	DataCreated = "SC00001"
	DataFetched = "SC00002"
	DataDeleted = "SC00003"
	DataUpdated = "SC00004"
)

var appResponsesSuccess = map[string]AppResponse{
	DataCreated: {
		HttpStatus: StatusCreated,
		Messages:   "Saved the entry.",
		Code:       DataCreated,
	},
	DataFetched: {
		HttpStatus: StatusOK,
		Messages:   "Data fetched.",
		Code:       DataFetched,
	},
	DataDeleted: {
		HttpStatus: StatusNoContent,
		Messages:   "Data deleted.",
		Code:       DataDeleted,
	},
	DataUpdated: {
		HttpStatus: StatusOK,
		Messages:   "Data updated.",
		Code:       DataUpdated,
	},
}
