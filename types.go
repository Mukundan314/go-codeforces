package codeforces

// User represents a Codeforces user.
//
// Codeforces API docs: https://codeforces.com/apiHelp/objects#User
type User struct {
	Handle                  string `json:"handle"`
	Email                   string `json:"email,omitempty"`
	VkID                    string `json:"vkId,omitempty"`
	OpenID                  string `json:"openId,omitempty"`
	FirstName               string `json:"firstName,omitempty"`
	LastName                string `json:"lastName,omitempty"`
	Country                 string `json:"country,omitempty"`
	City                    string `json:"city,omitempty"`
	Organization            string `json:"organization,omitempty"`
	Contribution            int    `json:"contribution"`
	Rank                    string `json:"rank"`
	Rating                  int    `json:"rating"`
	MaxRank                 string `json:"maxRank"`
	MaxRating               int    `json:"maxRating"`
	LastOnlineTimeSeconds   int    `json:"lastOnlineTimeSeconds"`
	RegistrationTimeSeconds int    `json:"registrationTimeSeconds"`
	FriendOfCount           int    `json:"friendOfCount"`
	Avatar                  string `json:"avatar"`
	TitlePhoto              string `json:"titlePhoto"`
}

// BlogEntry represents a Codeforces blog entry. May be in either short or full
// version.
//
// Codeforces API docs: https://codeforces.com/apiHelp/objects#BlogEntry
type BlogEntry struct {
	ID                      int      `json:"id"`
	OriginalLocale          string   `json:"originalLocale"`
	CreationTimeSeconds     int      `json:"creationTimeSeconds"`
	AuthorHandle            string   `json:"authorHandle"`
	Title                   string   `json:"title"`
	Content                 string   `json:"content"`
	Locale                  string   `json:"locale"`
	ModificationTimeSeconds int      `json:"modificationTimeSeconds"`
	AllowViewHistory        bool     `json:"allowViewHistory"`
	Tags                    []string `json:"tags"`
	Rating                  int      `json:"rating"`
}

// Comment represents a comment.
//
// Codeforces API docs: https://codeforces.com/apiHelp/objects#Comment
type Comment struct {
	ID                  int    `json:"id"`
	CreationTimeSeconds int    `json:"creationTimeSeconds"`
	CommentatorHandle   string `json:"commentatorHandle"`
	Locale              string `json:"locale"`
	Text                string `json:"text"`
	ParentCommentID     int    `json:"parentCommentId,omitempty"`
	Rating              int    `json:"rating"`
}

// RecentAction represents a recent action.
//
// Codeforces API docs: https://codeforces.com/apiHelp/objects#RecentAction
type RecentAction struct {
	TimeSeconds int       `json:"timeSeconds"`
	BlogEntry   BlogEntry `json:"blogEntry"`
	Comment     Comment   `json:"comment"`
}

// RatingChange represents a participation of user in rated contest.
//
// Codeforces API docs: https://codeforces.com/apiHelp/objects#RatingChange
type RatingChange struct {
	ContestID               int    `json:"contestId"`
	ContestName             string `json:"contestName"`
	Handle                  string `json:"handle"`
	Rank                    int    `json:"rank"`
	RatingUpdateTimeSeconds int    `json:"ratingUpdateTimeSeconds"`
	OldRating               int    `json:"oldRating"`
	NewRating               int    `json:"newRating"`
}

// Contest represents a contest on Codeforces.
//
// Codeforces API docs: https://codeforces.com/apiHelp/objects#Contest
type Contest struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Type                string `json:"type"`
	Phase               string `json:"phase"`
	Frozen              bool   `json:"frozen"`
	DurationSeconds     int    `json:"durationSeconds"`
	StartTimeSeconds    int    `json:"startTimeSeconds,omitempty"`
	RelativeTimeSeconds int    `json:"relativeTimeSeconds,omitempty"`
	PreparedBy          string `json:"preparedBy,omitempty"`
	WebsiteURL          string `json:"websiteUrl,omitempty"`
	Description         string `json:"description,omitempty"`
	Difficulty          int    `json:"difficulty,omitempty"`
	Kind                string `json:"kind,omitempty"`
	IcpcRegion          string `json:"icpcRegion,omitempty"`
	Country             string `json:"country,omitempty"`
	City                string `json:"city,omitempty"`
	Season              string `json:"season,omitempty"`
}

// Party represents a party, participating in a contest.
//
// Codeforces API docs: https://codeforces.com/apiHelp/objects#Party
type Party struct {
	ContestID        int      `json:"contestId,omitempty"`
	Members          []Member `json:"members"`
	ParticipantType  string   `json:"participantType"`
	TeamID           int      `json:"teamId,omitempty"`
	TeamName         string   `json:"teamName,omitempty"`
	Ghost            bool     `json:"ghost"`
	Room             int      `json:"room,omitempty"`
	StartTimeSeconds int      `json:"startTimeSeconds,omitempty"`
}

// Member represents a member of a party.
//
// Codeforces API docs: https://codeforces.com/apiHelp/objects#Member
type Member struct {
	Handle string `json:"handle"`
}

// Problem represents a problem.
//
// Codeforces API docs: https://codeforces.com/apiHelp/objects#Problem
type Problem struct {
	ContestID      int      `json:"contestId,omitempty"`
	ProblemsetName string   `json:"problemsetName,omitempty"`
	Index          string   `json:"index"`
	Name           string   `json:"name"`
	Type           string   `json:"type"`
	Points         float64  `json:"points,omitempty"`
	Rating         int      `json:"rating,omitempty"`
	Tags           []string `json:"tags"`
}

// ProblemStatistics represents a statistic data about a problem.
//
// Codeforces API docs: https://codeforces.com/apiHelp/objects#ProblemStatistics
type ProblemStatistics struct {
	ContestID   int    `json:"contestId,omitempty"`
	Index       string `json:"index"`
	SolvedCount int    `json:"solvedCount"`
}

// Submission represents a submission.
//
// Codeforces API docs: https://codeforces.com/apiHelp/objects#Submission
type Submission struct {
	ID                  int     `json:"id"`
	ContestID           int     `json:"contestId,omitempty"`
	CreationTimeSeconds int     `json:"creationTimeSeconds"`
	RelativeTimeSeconds int     `json:"relativeTimeSeconds"`
	Problem             Problem `json:"problem"`
	Author              Party   `json:"author"`
	ProgrammingLanguage string  `json:"programmingLanguage"`
	Verdict             string  `json:"verdict,omitempty"`
	Testset             string  `json:"testset"`
	PassedTestCount     int     `json:"passedTestCount"`
	TimeConsumedMillis  int     `json:"timeConsumedMillis"`
	MemoryConsumedBytes int     `json:"memoryConsumedBytes"`
}

// Hack represents a hack, made during Codeforces Round.
//
// Codeforces API docs: https://codeforces.com/apiHelp/objects#Hack
type Hack struct {
	ID                  int     `json:"id"`
	CreationTimeSeconds int     `json:"creationTimeSeconds"`
	Hacker              Party   `json:"hacker"`
	Defender            Party   `json:"defender"`
	Verdict             string  `json:"verdict,omitempty"`
	Problem             Problem `json:"problem"`
	Test                string  `json:"test,omitempty"`
	JudgeProtocol       struct {
		Manual   string `json:"manual"`
		Protocol string `json:"protocol"`
		Verdict  string `json:"verdict"`
	} `json:"judgeProtocol,omitempty"`
}

// RanklistRow represents a ranklist row.
//
// Codeforces API docs: https://codeforces.com/apiHelp/objects#RanklistRow
type RanklistRow struct {
	Party                     Party           `json:"party"`
	Rank                      int             `json:"rank"`
	Points                    float64         `json:"points"`
	Penalty                   int             `json:"penalty"`
	SuccessfulHackCount       int             `json:"successfulHackCount"`
	UnsuccessfulHackCount     int             `json:"unsuccessfulHackCount"`
	ProblemResults            []ProblemResult `json:"problemResults"`
	LastSubmissionTimeSeconds int             `json:"lastSubmissionTimeSeconds,omitempty"`
}

// ProblemResult represents a submissions results of a party for a problem.
//
// Codeforces API docs: https://codeforces.com/apiHelp/objects#ProblemResult
type ProblemResult struct {
	Points                    float64 `json:"points"`
	Penalty                   int     `json:"penalty"`
	RejectedAttemptCount      int     `json:"rejectedAttemptCount"`
	Type                      string  `json:"type"`
	BestSubmissionTimeSeconds int     `json:"bestSubmissionTimeSeconds"`
}
