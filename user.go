package codeforces

import (
	"strconv"
)

// GetUserBlogEntries returns a list of all user's blog entries.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#user.blogEntries
func (c *Client) GetUserBlogEntries(handle string) ([]BlogEntry, error) {
	params := make(map[string][]string)
	params["handle"] = []string{handle}

	var res []BlogEntry
	err := c.makeAPICall("user.blogEntries", params, &res)

	return res, err
}

// GetUserBlogEntries returns a list of all user's blog entries.
//
// GetUserBlogEntries is a wrapper around DefaultClient.GetUserBlogEntries.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#user.blogEntries
func GetUserBlogEntries(handle string) ([]BlogEntry, error) {
	return DefaultClient.GetUserBlogEntries(handle)
}

// GetUserFriends Returns authorized user's friends. Using this method requires
// authorization.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#user.friends
func (c *Client) GetUserFriends(onlyOnline bool) ([]string, error) {
	params := make(map[string][]string)
	params["onlyOnline"] = []string{strconv.FormatBool(onlyOnline)}

	var res []string
	err := c.makeAPICall("user.friends", params, &res)

	return res, err
}

// GetUserFriends returns authorized user's friends. Using this method requires
// authorization.
//
// GetUserFriends is a wrapper around DefaultClient.GetUserFriends.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#user.friends
func GetUserFriends(onlyOnline bool) ([]string, error) {
	return DefaultClient.GetUserFriends(onlyOnline)
}

// GetUserInfo returns information about one or several users.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#user.info
func (c *Client) GetUserInfo(handles []string) ([]User, error) {
	params := make(map[string][]string)
	params["handles"] = handles

	var res []User
	err := c.makeAPICall("user.info", params, &res)

	return res, err
}

// GetUserInfo returns information about one or several users.
//
// GetUserInfo is a wrapper around DefaultClient.GetUserInfo.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#user.info
func GetUserInfo(handles []string) ([]User, error) {
	return DefaultClient.GetUserInfo(handles)
}

// GetUserRatedList returns the list users who have participated in at least
// one rated contest.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#user.ratedList
func (c *Client) GetUserRatedList(activeOnly bool) ([]User, error) {
	params := make(map[string][]string)
	params["activeOnly"] = []string{strconv.FormatBool(activeOnly)}

	var res []User
	err := c.makeAPICall("user.ratedList", params, &res)

	return res, err
}

// GetUserRatedList returns the list users who have participated in at least
// one rated contest.
//
// GetUserRatedList is a wrapper around DefaultClient.GetUserRatedList.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#user.ratedList
func GetUserRatedList(activeOnly bool) ([]User, error) {
	return DefaultClient.GetUserRatedList(activeOnly)
}

// GetUserRating returns rating history of the specified user.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#user.rating
func (c *Client) GetUserRating(handle string) ([]RatingChange, error) {
	params := make(map[string][]string)
	params["handle"] = []string{handle}

	var res []RatingChange
	err := c.makeAPICall("user.rating", params, &res)

	return res, err
}

// GetUserRating returns rating history of the specified user.
//
// GetUserRating is a wrapper around DefaultClient.GetUserRating.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#user.rating
func GetUserRating(handle string) ([]RatingChange, error) {
	return DefaultClient.GetUserRating(handle)
}

// GetUserStatus returns submissions of specified user.
//
// Set count to 0 for infinite count.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#user.status
func (c *Client) GetUserStatus(handle string, from, count int) ([]Submission, error) {
	params := make(map[string][]string)
	params["handle"] = []string{handle}
	params["from"] = []string{strconv.FormatInt(int64(from), 10)}
	if count != 0 {
		params["count"] = []string{strconv.FormatInt(int64(count), 10)}
	}

	var res []Submission
	err := c.makeAPICall("user.status", params, &res)

	return res, err
}

// GetUserStatus returns submissions of specified user.
//
// Set count to 0 for infinite count.
//
// GetUserStatus is a wrapper around DefaultClient.GetUserStatus.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#user.status
func GetUserStatus(handle string, from, count int) ([]Submission, error) {
	return DefaultClient.GetUserStatus(handle, from, count)
}
