package jsonModel

type (
	//JSONResponse - for response api
	JSONResponse struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	//JSONResponsePost - for response api with update or post data
	JSONResponsePost struct {
		Status   int    `json:"status"`
		Messages string `json:"message"`
	}

	//JSONResponseBadRequest - for response api with response 400 (badrequest)
	JSONResponseBadRequest struct {
		Status     int         `json:"status"`
		Validation interface{} `json:"validation"`
	}

	//JSONErrorResponse response api error
	JSONErrorResponse struct {
		Status   int    `json:"status"`
		Messages string `json:"message"`
		Target   string `json:"target"`
	}
)
