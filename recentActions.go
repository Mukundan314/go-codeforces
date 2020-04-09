package codeforces

import (
	"strconv"
)

// Returns recent actions.
// Codeforces API docs: https://codeforces.com/apiHelp/methods#recentActions
func (c *Client) GetRecentActions(maxCount int) ([]RecentAction, error) {
	params := make(map[string][]string)
	params["maxCount"] = []string{strconv.FormatInt(int64(maxCount), 10)}

	var res []RecentAction
	err := c.makeApiCall("recentActions", params, &res)

	return res, err
}

// Returns recent actions.
// GetRecentActions is a wrapper around DefaultClient.GetRecentActions
// Codeforces API docs: https://codeforces.com/apiHelp/methods#recentActions
func GetRecentActions(maxCount int) ([]RecentAction, error) {
	return DefaultClient.GetRecentActions(maxCount)
}
