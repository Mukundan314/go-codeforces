package codeforces

import (
	"strconv"
)

// GetContestHacks returns list of hacks in the specified contests. Full
// information about hacks is available only after some time after the contest
// end. During the contest user can see only own hacks.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#contest.hacks
func (c *Client) GetContestHacks(contestID int) ([]Hack, error) {
	params := make(map[string][]string)
	params["contestId"] = []string{strconv.FormatInt(int64(contestID), 10)}

	var res []Hack
	err := c.makeAPICall("contest.hacks", params, &res)

	return res, err
}

// GetContestHacks returns list of hacks in the specified contests. Full
// information about hacks is available only after some time after the contest
// end. During the contest user can see only own hacks.
//
// GetContestHacks is a wrapper around DefaultClient.GetContestHacks.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#contest.hacks
func GetContestHacks(contestID int) ([]Hack, error) {
	return DefaultClient.GetContestHacks(contestID)
}

// GetContestList returns information about all available contests.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#contest.list
func (c *Client) GetContestList(gym bool) ([]Contest, error) {
	params := make(map[string][]string)
	params["gym"] = []string{strconv.FormatBool(gym)}

	var res []Contest
	err := c.makeAPICall("contest.list", params, &res)

	return res, err
}

// GetContestList returns information about all available contests.
//
// GetContestList is a wrapper around DefaultClient.GetContestList.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#contest.list
func GetContestList(gym bool) ([]Contest, error) {
	return DefaultClient.GetContestList(gym)
}

// GetContestRatingChanges returns rating changes after the contest.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#contest.ratingChanges
func (c *Client) GetContestRatingChanges(contestID int) ([]RatingChange, error) {
	params := make(map[string][]string)
	params["contestId"] = []string{strconv.FormatInt(int64(contestID), 10)}

	var res []RatingChange
	err := c.makeAPICall("contest.ratingChanges", params, &res)

	return res, err
}

// GetContestRatingChanges returns rating changes after the contest.
//
// GetContestRatingChanges is a wrapper around
// DefaultClient.GetContestRatingChanges.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#contest.ratingChanges
func GetContestRatingChanges(contestID int) ([]RatingChange, error) {
	return DefaultClient.GetContestRatingChanges(contestID)
}

// GetContestStandings returns the description of the contest and the requested
// part of the standings.
//
// Set count to 0 for infinite count.
// Set handles to a empty list to not filter by handles.
// Set room to 0 to show all participants.
//
// Returns object with three fields: "contest", "problems" and "rows".
// Field "contest" contains a Contest object.
// Field "problems" contains a list of Problem objects.
// Field "rows" contains a list of RanklistRow objects.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#contest.standings
func (c *Client) GetContestStandings(contestID, from, count int, handles []string, room int, showUnofficial bool) (Contest, []Problem, []RanklistRow, error) {
	params := make(map[string][]string)

	params["contestId"] = []string{strconv.FormatInt(int64(contestID), 10)}
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

	err := c.makeAPICall("contest.standings", params, &res)

	return res.Contest, res.Problems, res.Rows, err
}

// GetContestStandings returns the description of the contest and the requested
// part of the standings.
//
// Set count to 0 for infinite count.
// Set handles to a empty list to not filter by handles.
// Set room to 0 to show all participants.
//
// Returns object with three fields: "contest", "problems" and "rows".
// Field "contest" contains a Contest object.
// Field "problems" contains a list of Problem objects.
// Field "rows" contains a list of RanklistRow objects.
//
// GetContestStandings is a wrapper around DefaultClient.GetContestStandings.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#contest.standings
func GetContestStandings(contestID, from, count int, handles []string, room int, showUnofficial bool) (Contest, []Problem, []RanklistRow, error) {
	return DefaultClient.GetContestStandings(contestID, from, count, handles, room, showUnofficial)
}

// GetContestStatus returns submissions for specified contest. Optionally can
// return submissions of specified user.
//
// Set handle to a empty string to get status for all handles.
// Set count to 0 for infinite count.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#contest.status
func (c *Client) GetContestStatus(contestID int, handle string, from, count int) ([]Submission, error) {
	params := make(map[string][]string)

	params["contestId"] = []string{strconv.FormatInt(int64(contestID), 10)}
	params["from"] = []string{strconv.FormatInt(int64(from), 10)}

	if count != 0 {
		params["count"] = []string{strconv.FormatInt(int64(count), 10)}
	}
	if handle != "" {
		params["handle"] = []string{handle}
	}

	var res []Submission
	err := c.makeAPICall("contest.status", params, &res)

	return res, err
}

// GetContestStatus returns submissions for specified contest. Optionally can
// return submissions of specified user.
//
// Set handle to a empty string to get status for all handles.
// Set count to 0 for infinite count.
//
// GetContestStatus is a wrapper around DefaultClient.GetContestStatus.
//
// Codeforces API docs: https://codeforces.com/apiHelp/methods#contest.status
func GetContestStatus(contestID int, handle string, from, count int) ([]Submission, error) {
	return DefaultClient.GetContestStatus(contestID, handle, from, count)
}
