package sf

import (
	`crypto/md5`
	`encoding/base64`
	`encoding/json`
	`fmt`
	`io/ioutil`
	`net/http`
	`net/url`
	`time`

	`github.com/google/uuid`
	`github.com/pkg/errors`
)

type Client struct {
	URL       string `json:"url"`
	PartnerID string `json:"partnerID"`
	CheckWord string `json:"check_word"`
	*http.Client
}

func NewClient(url, partnerID, checkWord string) *Client {
	return &Client{
		URL:       url,
		PartnerID: partnerID,
		CheckWord: checkWord,
		Client:    &http.Client{},
	}
}

func (c *Client) do(msgData []byte, serviceCode string) (int, []byte, error) {
	// first md5, then base64
	timestamp := time.Now().Unix()
	str := url.QueryEscape(fmt.Sprintf("%s%d%s", string(msgData), timestamp, c.CheckWord))
	md := md5.New()
	_, err := md.Write([]byte(str))
	if err != nil {
		err = errors.Wrap(err, "md5 err")
		return http.StatusInternalServerError, nil, err
	}
	digest := base64.StdEncoding.EncodeToString(md.Sum(nil))
	// generate request id
	uuid, err := uuid.NewUUID()
	if err != nil {
		err = errors.Wrap(err, "generate uuid err")
		return http.StatusInternalServerError, nil, err
	}

	postData := url.Values{}
	postData.Add("partnerID", c.PartnerID)
	postData.Add("requestID", uuid.String())
	postData.Add("serviceCode", serviceCode)
	postData.Add("timestamp", fmt.Sprintf("%d", timestamp))
	postData.Add("msgDigest", digest)
	postData.Add("msgData", string(msgData))

	resp, err := c.PostForm(c.URL, postData)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "read response body err")
		return resp.StatusCode, nil, err
	}
	return resp.StatusCode, respData, nil
}

func (c *Client) Do(data interface{}, serviceCode string) (map[string]interface{}, error) {
	msgData, err := json.Marshal(data)
	if err != nil {
		err = errors.Wrap(err, "marshal data err")
		return nil, err
	}
	code, respBytes, err := c.do(msgData, serviceCode)
	if err != nil {
		err = errors.Wrap(err, "do request err")
		return nil, err
	}
	if code != http.StatusOK {
		return nil, errors.New(http.StatusText(code))
	}
	resp := &BaseResp{}
	err = json.Unmarshal(respBytes, resp)
	if err != nil {
		err = errors.Wrap(err, "unmarshal response err")
		return nil, err
	}

	if resp.ApiResultCode != serviceApiResultCodeA1000 {
		return nil, errors.New(fmt.Sprintf("errorCode: %s errorMsg: %s", resp.ApiResultCode, resp.ApiResultCode))
	}

	apiData := &ApiResultData{}
	err = json.Unmarshal([]byte(resp.ApiResultData), apiData)
	if err != nil {
		err = errors.Wrap(err, "unmarshal response err")
		return nil, err
	}

	if !apiData.Success {
		return nil, errors.New(fmt.Sprintf("errorCode: %s errorMsg: %s", apiData.ErrorCode, apiData.ErrorMsg))
	}

	return apiData.MsgData.(map[string]interface{}), nil
}
