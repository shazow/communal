package store

import "time"

type NewsID int
type News struct {
	NewsID      NewsID
	OwnerUserID UserID
}

type UserID int
type User struct {
	UserID UserID

	DisplayName string
}

type FeedID int
type Feed struct {
	FeedID FeedID

	Link string // Link is the canonical URL
}

type ItemID int
type Item struct {
	ItemID ItemID
	FeedID FeedID

	Title string
	Body  string

	TimeFetched   *time.Time
	TimePublished *time.Time

	Link         string // Link is the canonical URL of the post
	RelatedLinks []string
}

type Store interface {
	AddUser(*User) error
	GetUserByID(userID UserID) (*User, error)

	AddItem(*Item) error
}
