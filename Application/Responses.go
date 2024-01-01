package Application

import (
	"github.com/bykovme/gotrans"
)

func (req Request) Ok(body interface{}) {
	req.Resonse(200, buildResponse(body, gotrans.T("ok"), 200, nil))
}

func (req Request) Created(body interface{}) {
	req.Resonse(201, buildResponse(body, gotrans.T("created"), 201, nil))
}

func (req Request) NotAuth() {
	req.Resonse(401, buildResponse(nil, gotrans.T("you_are_not_auth"), 401, nil))
}

func (req Request) BadRequest(err interface{}) {
	req.Resonse(422, buildResponse(nil, gotrans.T("somthing_wrong_in_your_request"), 422, err))
}

func (req Request) UserNotFound() {
	req.Resonse(404, buildResponse(nil, gotrans.T("we_not_foun_you_in_our_system"), 404, nil))
}

func buildResponse(payload interface{}, message string, code int, err interface{}) map[string]interface{} {
	response := make(map[string]interface{})
	response["payload"] = payload
	response["message"] = message
	response["code"] = code
	response["errors"] = err
	return response
}
