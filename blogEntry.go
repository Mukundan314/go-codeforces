package codeforces

import (
	"strconv"
)

// Returns a list of comments to the specified blog entry.
// Codeforces API docs: https://codeforces.com/apiHelp/methods#blogEntry.comments
func (c *Client) GetBlogEntryComments(blogEntryId int) ([]Comment, error) {
	params := make(map[string][]string)
	params["blogEntryId"] = []string{strconv.FormatInt(int64(blogEntryId), 10)}

	var res []Comment
	err := c.makeApiCall("blogEntry.comments", params, &res)

	return res, err
}

// Returns a list of comments to the specified blog entry.
// GetBlogEntryComments is a wrapper around DefaultClient.GetBlogEntryComments
// Codeforces API docs: https://codeforces.com/apiHelp/methods#blogEntry.comments
func GetBlogEntryComments(blogEntryId int) ([]Comment, error) {
	return DefaultClient.GetBlogEntryComments(blogEntryId)
}

// Returns blog entry.
// Codeforces Api docs: https://codeforces.com/apiHelp/methods#blogEntry.view
func (c *Client) GetBlogEntry(blogEntryId int) (BlogEntry, error) {
	params := make(map[string][]string)
	params["blogEntryId"] = []string{strconv.FormatInt(int64(blogEntryId), 10)}

	var res BlogEntry
	err := c.makeApiCall("blogEntry.view", params, &res)

	return res, err
}

// Returns blog entry.
// GetBlogEntry is a wrapper around DefaultClient.GetBlogEntry
// Codeforces API docs: https://codeforces.com/apiHelp/methods#blogEntry.view
func GetBlogEntry(blogEntryId int) (BlogEntry, error) {
	return DefaultClient.GetBlogEntry(blogEntryId)
}
