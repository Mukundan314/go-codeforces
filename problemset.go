package codeforces

import (
	"strconv"
)

// Leave problemsetName empty for default problemset
// Returns all problems from problemset. Problems can be filtered by tags.
// Codeforces API docs: https://codeforces.com/apiHelp/methods#problemset.problems
func (c *Client) GetProblemsetProblems(tags []string, problemsetName string) ([]Problem, []ProblemStatistics, error) {
	params := make(map[string][]string)

	if len(tags) > 0 {
		params["tags"] = tags
	}
	if problemsetName != "" {
		params["problemsetName"] = []string{problemsetName}
	}

	var res struct {
		Problems          []Problem           `json:"problems"`
		ProblemStatistics []ProblemStatistics `json:"problemStatistics"`
	}

	err := c.makeApiCall("problemset.problems", params, &res)

	return res.Problems, res.ProblemStatistics, err
}

// Leave problemsetName empty for default problemset
// Returns all problems from problemset. Problems can be filtered by tags.
// GetProblemsetProblems is a wrapper around DefaultClient.GetProblemsetProblems
// Codeforces API docs: https://codeforces.com/apiHelp/methods#problemset.problems
func GetProblemsetProblems(tags []string, problemsetName string) ([]Problem, []ProblemStatistics, error) {
	return DefaultClient.GetProblemsetProblems(tags, problemsetName)
}

// Leave problemsetName empty for default problemset
// Returns recent submissions.
// Codeforces API docs: https://codeforces.com/apiHelp/methods#problemset.recentStatus
func (c *Client) GetProblemsetRecentStatus(count int, problemsetName string) ([]Submission, error) {
	params := make(map[string][]string)

	params["count"] = []string{strconv.FormatInt(int64(count), 10)}
	if problemsetName != "" {
		params["problemsetName"] = []string{problemsetName}
	}

	var res []Submission
	err := c.makeApiCall("problemset.recentStatus", params, &res)

	return res, err
}

// Leave problemsetName empty for default problemset
// Returns recent submissions.
// GetProblemsetRecentStatus is a wrapper around DefaultClient.GetProblemsetRecentStatus
// Codeforces API docs: https://codeforces.com/apiHelp/methods#problemset.recentStatus
func GetProblemsetRecentStatus(count int, problemsetName string) ([]Submission, error) {
	return DefaultClient.GetProblemsetRecentStatus(count, problemsetName)
}
