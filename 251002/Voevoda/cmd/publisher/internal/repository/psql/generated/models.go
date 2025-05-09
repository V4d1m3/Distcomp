// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package generated

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type TblLabel struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type TblNote struct {
	ID      int64  `json:"id"`
	Tweetid *int64 `json:"tweetid"`
	Content string `json:"content"`
}

type TblTweet struct {
	ID       int64            `json:"id"`
	UserID   *int64           `json:"user_id"`
	Title    string           `json:"title"`
	Content  string           `json:"content"`
	Created  pgtype.Timestamp `json:"created"`
	Modified pgtype.Timestamp `json:"modified"`
}

type TblTweetLabel struct {
	ID      int64  `json:"id"`
	Tweetid *int64 `json:"tweetid"`
	Labelid *int64 `json:"labelid"`
}

type TblUser struct {
	ID        int64  `json:"id"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
