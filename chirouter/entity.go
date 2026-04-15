package chirouter

import (
	"errors"
	"net/http"
	"strings"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type UserPayload struct {
	*User
	Role string `json:"role"`
}

type Article struct {
	ID     string `json:"id"`
	UserID int64  `json:"user_id"` // the author
	Title  string `json:"title"`
	Slug   string `json:"slug"`
}

type ArticleRequest struct {
	*Article

	User *UserPayload `json:"user,omitempty"`

	ProtectedID string `json:"id"`
}

func (a *ArticleRequest) Bind(r *http.Request) error {
	if a.Article == nil {
		return errors.New("missing required Article fields.")
	}
	a.ProtectedID = ""
	a.Article.Title = strings.ToLower(a.Article.Title)
	return nil
}

type ArticleResponse struct {
	*Article

	User *UserPayload `json:"user,omitempty"`

	Elapsed int64 `json:"elapsed"`
}

func (rd *ArticleResponse) Render(w http.ResponseWriter, r *http.Request) error {
	rd.Elapsed = 10
	return nil
}

type ArticleListResponse struct {
	Data  []*Article `json:"data"`
	Total int        `json:"total"`
	Page  int        `json:"page"`
	Limit int        `json:"limit"`
}

func (rlr *ArticleListResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
