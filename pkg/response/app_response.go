package response

var appResponses = make(map[string]AppResponse)

// AppResponseCode
func init() {
	for k, v := range appResponsesGeneral {
		appResponses[k] = v
	}
	for k, v := range appResponsesSuccess {
		appResponses[k] = v
	}
	for k, v := range appResponsesError {
		appResponses[k] = v
	}
}
