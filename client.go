package sf

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
)

const (
	defaultURL       = "https://sfapi-sbox.sf-express.com/std/service"
	defaultPartnerID = "LBLERiwZfH"
	defaultCheckWord = "oNv0SZTAyeUwRCTUNGvswI1D5oQPHU90"
)

type Opt func(client *Client) error

type Client struct {
	url       string
	partnerID string
	checkWord string
	client    *http.Client
}

func NewClientWithOpts(opts ...Opt) (*Client, error) {
	c := &Client{
		url:       defaultURL,
		partnerID: defaultPartnerID,
		checkWord: defaultCheckWord,
		client:    &http.Client{},
	}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func WithURL(u string) Opt {
	return func(c *Client) error {
		if _, err := url.Parse(u); err != nil {
			return errors.New("url invalid")
		}
		c.url = u
		return nil
	}
}

func WithPartnerID(partnerID string) Opt {
	return func(c *Client) error {
		if partnerID == "" {
			return errors.New("partnerID invalid")
		}
		c.partnerID = partnerID
		return nil
	}
}

func WithCheckWord(checkWord string) Opt {
	return func(c *Client) error {
		if checkWord == "" {
			return errors.New("checkWord invalid")
		}
		c.checkWord = checkWord
		return nil
	}
}

func WithHTTPClient(client *http.Client) Opt {
	return func(c *Client) error {
		c.client = client
		return nil
	}
}

func (c *Client) do(msgData []byte, serviceCode string) (int, []byte, error) {
	// first md5, then base64
	timestamp := time.Now().Unix()
	str := url.QueryEscape(fmt.Sprintf("%s%d%s", string(msgData), timestamp, c.checkWord))
	md := md5.New()
	_, err := md.Write([]byte(str))
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	digest := base64.StdEncoding.EncodeToString(md.Sum(nil))
	// generate request id
	uuID, err := uuid.NewUUID()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	postData := url.Values{}
	postData.Add("partnerID", c.partnerID)
	postData.Add("requestID", uuID.String())
	postData.Add("serviceCode", serviceCode)
	postData.Add("timestamp", fmt.Sprintf("%d", timestamp))
	postData.Add("msgDigest", digest)
	postData.Add("msgData", string(msgData))

	resp, err := c.client.PostForm(c.url, postData)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	return resp.StatusCode, respData, nil
}

func (c *Client) Do(data interface{}, serviceCode string) (interface{}, error) {
	msgData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	code, respBytes, err := c.do(msgData, serviceCode)
	if err != nil {
		return nil, err
	}
	if code != http.StatusOK {
		return nil, errors.New(http.StatusText(code))
	}
	resp := &BaseResp{}
	err = json.Unmarshal(respBytes, resp)
	if err != nil {
		return nil, err
	}

	if resp.ApiResultCode != serviceApiResultCodeA1000 {
		return nil, errors.New(fmt.Sprintf("errorCode: %s errorMsg: %s", resp.ApiResultCode, resp.ApiResultCode))
	}

	apiData := &ApiResultData{}
	err = json.Unmarshal([]byte(resp.ApiResultData), apiData)
	if err != nil {
		return nil, err
	}

	if !apiData.Success || apiData.ErrorCode != errorCodeS0000 {
		return nil, errors.New(fmt.Sprintf("errorCode: %s errorMsg: %s", apiData.ErrorCode, apiData.ErrorMsg))
	}

	return apiData.MsgData, nil
}
