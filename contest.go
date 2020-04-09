package codeforces

import (
	"strconv"
)

// Returns list of hacks in the specified contests. Full information about
// hacks is available only after some time after the contest end. During the
// contest user can see only own hacks.
// Codeforces API docs: https://codeforces.com/apiHelp/methods#contest.hacks
func (c *Client) GetContestHacks(contestId int) ([]Hack, error) {
	params := make(map[string][]string)
	params["contestId"] = []string{strconv.FormatInt(int64(contestId), 10)}

	var res []Hack
	err := c.makeApiCall("contest.hacks", params, &res)

	return res, err
}

// Returns list of hacks in the specified contests. Full information about
// hacks is available only after some time after the contest end. During the
// contest user can see only own hacks.
// GetContestHacks is a wrapper around DefaultClient.GetContestHacks
// Codeforces API docs: https://codeforces.com/apiHelp/methods#contest.hacks
func GetContestHacks(contestId int) ([]Hack, error) {
	return DefaultClient.GetContestHacks(contestId)
}

// Returns information about all available contests.
// Codeforces API docs: https://codeforces.com/apiHelp/methods#contest.list
func (c *Client) GetContestList(gym bool) ([]Contest, error) {
	params := make(map[string][]string)
	params["gym"] = []string{strconv.FormatBool(gym)}

	var res []Contest
	err := c.makeApiCall("contest.list", params, &res)

	return res, err
}

// Returns information about all available contests.
// GetContestList is a wrapper around DefaultClient.GetContestList
// Codeforces API docs: https://codeforces.com/apiHelp/methods#contest.list
func GetContestList(gym bool) ([]Contest, error) {
	return DefaultClient.GetContestList(gym)
}

// Returns rating changes after the contest.
// Codeforces API docs: https://codeforces.com/apiHelp/methods#contest.ratingChanges
func (c *Client) GetContestRatingChanges(contestId int) ([]RatingChange, error) {
	params := make(map[string][]string)
	params["contestId"] = []string{strconv.FormatInt(int64(contestId), 10)}

	var res []RatingChange
	err := c.makeApiCall("contest.ratingChanges", params, &res)

	return res, err
}

// Returns rating changes after the contest.
// GetContestRatingChanges is a wrapper around DefaultClient.GetContestRatingChanges
// Codeforces API docs: https://codeforces.com/apiHelp/methods#contest.ratingChanges
func GetContestRatingChanges(contestId int) ([]RatingChange, error) {
	return DefaultClient.GetContestRatingChanges(contestId)
}

// Set count to 0 for infinite count.
// Set handles to a empty list to not filter by handles.
// Set room to 0 to show all participants.
// Returns the description of the contest and the requested part of the standings.
// Returns object with three fields: "contest", "problems" and "rows".
// Field "contest" contains a Contest object.
// Field "problems" contains a list of Problem objects.
// Field "rows" contains a list of RanklistRow objects.
// Codeforces API docs: https://codeforces.com/apiHelp/methods#contest.standings
func (c *Client) GetContestStandings(contestId, from, count int, handles []string, room int, showUnofficial bool) (Contest, []Problem, []RanklistRow, error) {
	params := make(map[string][]string)

	params["contestId"] = []string{strconv.FormatInt(int64(contestId), 10)}
	params["from"] = []string{strconv.FormatInt(int64(from), 10)}
	params["showUnofficial"] = []string{strconv.FormatBool(showUnofficial)}

	if count != 0 {
		params["count"] = []string{strconv.FormatInt(int64(count), 10)}
	}
	if len(handles) > 0 {
		params["handles"] = handles
	}
	if room != 0 {
		params["room"] = []string{strconv.FormatInt(int64(room), 10)}
	}

	var res struct {
		Contest  Contest       `json:"contest"`
		Problems []Problem     `json:"problems"`
		Rows     []RanklistRow `json:"rows"`
	}

	err := c.makeApiCall("contest.standings", params, &res)

	return res.Contest, res.Problems, res.Rows, err
}

// Set count to 0 for infinite count.
// Set handles to a empty list to not filter by handles.
// Set room to 0 to show all participants.
// Returns the description of the contest and the requested part of the standings.
// Returns object with three fields: "contest", "problems" and "rows".
// Field "contest" contains a Contest object.
// Field "problems" contains a list of Problem objects.
// Field "rows" contains a list of RanklistRow objects.
// GetContestStandings is a wrapper around DefaultClient.GetContestStandings
// Codeforces API docs: https://codeforces.com/apiHelp/methods#contest.standings
func GetContestStandings(contestId, from, count int, handles []string, room int, showUnofficial bool) (Contest, []Problem, []RanklistRow, error) {
	return DefaultClient.GetContestStandings(contestId, from, count, handles, room, showUnofficial)
}

// Set handle to a empty string to get status for all handles.
// Set count to 0 for infinite count.
// Returns submissions for specified contest. Optionally can return submissions
// of specified user.
// Codeforces API docs: https://codeforces.com/apiHelp/methods#contest.status
func (c *Client) GetContestStatus(contestId int, handle string, from, count int) ([]Submission, error) {
	params := make(map[string][]string)

	params["contestId"] = []string{strconv.FormatInt(int64(contestId), 10)}
	params["from"] = []string{strconv.FormatInt(int64(from), 10)}

	if count != 0 {
		params["count"] = []string{strconv.FormatInt(int64(count), 10)}
	}
	if handle != "" {
		params["handle"] = []string{handle}
	}

	var res []Submission
	err := c.makeApiCall("contest.status", params, &res)

	return res, err
}

// Set handle to a empty string to get status for all handles.
// Set count to 0 for infinite count.
// Returns submissions for specified contest. Optionally can return submissions
// of specified user.
// GetContestStatus is a wrapper around DefaultClient.GetContestStatus
// Codeforces API docs: https://codeforces.com/apiHelp/methods#contest.status
func GetContestStatus(contestId int, handle string, from, count int) ([]Submission, error) {
	return DefaultClient.GetContestStatus(contestId, handle, from, count)
}
