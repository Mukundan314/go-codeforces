package codeforces

import (
	"strconv"
)

// GetBlogEntryComments returns a list of comments to the specified blog entry.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#blogEntry.comments
func (c *Client) GetBlogEntryComments(blogEntryID int) ([]Comment, error) {
	params := make(map[string][]string)
	params["blogEntryId"] = []string{strconv.FormatInt(int64(blogEntryID), 10)}

	var res []Comment
	err := c.makeAPICall("blogEntry.comments", params, &res)

	return res, err
}

// GetBlogEntryComments returns a list of comments to the specified blog entry.
//
// GetBlogEntryComments is a wrapper around DefaultClient.GetBlogEntryComments
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#blogEntry.comments
func GetBlogEntryComments(blogEntryID int) ([]Comment, error) {
	return DefaultClient.GetBlogEntryComments(blogEntryID)
}

// GetBlogEntry returns blog entry.
//
// Codeforces Api docs: https://codeforces.com/apiHelp/methods#blogEntry.view
func (c *Client) GetBlogEntry(blogEntryID int) (BlogEntry, error) {
	params := make(map[string][]string)
	params["blogEntryId"] = []string{strconv.FormatInt(int64(blogEntryID), 10)}

	var res BlogEntry
	err := c.makeAPICall("blogEntry.view", params, &res)

	return res, err
}

// GetBlogEntry returns blog entry.
//
// GetBlogEntry is a wrapper around DefaultClient.GetBlogEntry
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#blogEntry.view
func GetBlogEntry(blogEntryID int) (BlogEntry, error) {
	return DefaultClient.GetBlogEntry(blogEntryID)
}
