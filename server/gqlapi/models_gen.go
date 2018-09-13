// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlapi

// セッションに対していいね！する時の引数です。
type CreateLikeInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	SessionID        string  `json:"sessionId"`
}

type CreateLikePayload struct {
	ClientMutationID *string `json:"clientMutationId"`
	Like             Like    `json:"like"`
}

// セッションに対していいね！した情報です。
type Like struct {
	ID      string  `json:"id"`
	Session Session `json:"session"`
}

// お知らせです。
type News struct {
	ID        string  `json:"id"`
	NewsID    string  `json:"newsId"`
	Date      string  `json:"date"`
	Message   string  `json:"message"`
	MessageJa string  `json:"messageJa"`
	Link      *string `json:"link"`
}

type NewsConnection struct {
	PageInfo PageInfo   `json:"pageInfo"`
	Edges    []NewsEdge `json:"edges"`
	Nodes    []News     `json:"nodes"`
}

type NewsEdge struct {
	Cursor *string `json:"cursor"`
	Node   News    `json:"node"`
}

type Node interface{}

type PageInfo struct {
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
}

// セッション1つに関する情報。
type Session struct {
	ID        string    `json:"id"`
	SessionID int       `json:"sessionId"`
	Type      string    `json:"type"`
	Place     string    `json:"place"`
	Title     string    `json:"title"`
	TitleJa   string    `json:"titleJa"`
	StartTime string    `json:"startTime"`
	EndTime   string    `json:"endTime"`
	Outline   string    `json:"outline"`
	OutlineJa string    `json:"outlineJa"`
	Lang      string    `json:"lang"`
	Tags      []string  `json:"tags"`
	Liked     int       `json:"liked"`
	Speakers  []Speaker `json:"speakers"`
}

type SessionConnection struct {
	PageInfo PageInfo      `json:"pageInfo"`
	Edges    []SessionEdge `json:"edges"`
	Nodes    []Session     `json:"nodes"`
}

type SessionEdge struct {
	Cursor *string `json:"cursor"`
	Node   Session `json:"node"`
}

type SessionListInput struct {
	Tmp *string `json:"tmp"`
}

// スピーカー1人に関する情報。
type Speaker struct {
	ID         string    `json:"id"`
	SpeakerID  string    `json:"speakerId"`
	Name       string    `json:"name"`
	NameJa     string    `json:"nameJa"`
	Company    string    `json:"company"`
	Position   string    `json:"position"`
	PositionJa string    `json:"positionJa"`
	Profile    string    `json:"profile"`
	ProfileJa  string    `json:"profileJa"`
	IconURL    string    `json:"iconUrl"`
	TwitterID  string    `json:"twitterId"`
	GithubID   string    `json:"githubId"`
	Sessions   []Session `json:"sessions"`
}
