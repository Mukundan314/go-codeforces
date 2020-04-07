package codeforces

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	apiKey     *string
	apiSecret  *string
	locale     *string
	httpClient *http.Client
}

type apiResponse struct {
	Status  string          `json:"status"`
	Comment string          `json:"comment,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

var DefaultClient = NewClient()

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

func (c *Client) getApiSig(method string, params map[string][]string) (string, error) {
	q := url.Values{}
	for k, v := range params {
		q.Set(k, strings.Join(v, ";"))
	}

	r := fmt.Sprintf("%06d", rand.Intn(1000000))

	sum := sha512.Sum512([]byte(fmt.Sprintf("%s/%s?%s#%s", r, method, q.Encode(), *c.apiSecret)))
	apiSig := hex.EncodeToString(sum[:])

	return r + apiSig, nil
}

func (c *Client) makeApiCall(method string, params map[string][]string, v interface{}) error {
	u, err := url.Parse("http://codeforces.com/api/")
	if err != nil {
		return err
	}

	if c.locale != nil {
		params["lang"] = []string{*c.locale}
	}

	if (c.apiKey != nil) && (c.apiSecret != nil) {
		params["time"] = []string{strconv.FormatInt(time.Now().Unix(), 10)}
		params["apiKey"] = []string{*c.apiKey}

		apiSig, err := c.getApiSig(method, params)
		if err != nil {
			return err
		}

		params["apiSig"] = []string{apiSig}
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, strings.Join(v, ";"))
	}

	u.Path = path.Join(u.Path, method)
	u.RawQuery = q.Encode()

	resp, err := c.httpClient.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var res apiResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return err
	}

	if res.Status == "FAILED" {
		return errors.New(res.Comment)
	}

	return json.Unmarshal(res.Result, v)
}

func (c *Client) SetApiKey(apiKey, apiSecret string) {
	c.apiKey = &apiKey
	c.apiSecret = &apiSecret
}

func (c *Client) SetLocale(locale string) {
	c.locale = &locale
}
