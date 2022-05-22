package sf

type BaseResp struct {
	ApiErrorMsg   string `json:"apiErrorMsg"`
	ApiResponseID string `json:"apiResponseID"`
	ApiResultCode string `json:"apiResultCode"`
	ApiResultData string `json:"apiResultData"`
}

type ApiResultData struct {
	Success   bool        `json:"success"`
	ErrorCode string      `json:"errorCode"`
	ErrorMsg  string      `json:"errorMsg"`
	MsgData   interface{} `json:"msgData"`
}
