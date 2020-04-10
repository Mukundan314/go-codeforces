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

// DefaultClient is the default Client and is used by GetBlogEntryComments,
// GetBlogEntry, GetContestHacks, GetContestList, GetContestRatingChanges,
// GetContestStandings, GetContestStatus, GetProblemsetProblems,
// GetProblemsetRecentStatus, GetRecentActions, GetUserBlogEntries,
// GetUserFriends, GetUserInfo, GetUserRatedList, GetUserRating and
// GetUserStatus
var DefaultClient = NewClient()

// NewClient creates a new Client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

func (c *Client) getAPISig(method string, params map[string][]string) (string, error) {
	q := url.Values{}
	for k, v := range params {
		q.Set(k, strings.Join(v, ";"))
	}

	r := fmt.Sprintf("%06d", rand.Intn(1000000))

	sum := sha512.Sum512([]byte(fmt.Sprintf("%s/%s?%s#%s", r, method, q.Encode(), *c.apiSecret)))
	apiSig := hex.EncodeToString(sum[:])

	return r + apiSig, nil
}

func (c *Client) makeAPICall(method string, params map[string][]string, v interface{}) error {
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

		apiSig, err := c.getAPISig(method, params)
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

// SetAPIKey sets apiKey and apiSecret of a client
func (c *Client) SetAPIKey(apiKey, apiSecret string) {
	c.apiKey = &apiKey
	c.apiSecret = &apiSecret
}

// SetLocale sets locale of a client
func (c *Client) SetLocale(locale string) {
	c.locale = &locale
}
