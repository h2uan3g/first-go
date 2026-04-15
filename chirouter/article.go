package chirouter

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var articles = []*Article{
	{ID: "1", UserID: 100, Title: "Hi", Slug: "hi"},
	{ID: "2", UserID: 200, Title: "sup", Slug: "sup"},
	{ID: "3", UserID: 300, Title: "alo", Slug: "alo"},
	{ID: "4", UserID: 400, Title: "bonjour", Slug: "bonjour"},
	{ID: "5", UserID: 500, Title: "whats up", Slug: "whats-up"},
}

var users = []*User{
	{ID: 100, Name: "Peter"},
	{ID: 200, Name: "Julia"},
}

type articlesResource struct{}

func (rs articlesResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.With(Paginate).Get("/", ListArticles)
	r.Post("/", CreateArticle) // POST /articles
	// r.Get("/search", SearchArticles) // GET /articles/search

	r.Route("/{articleID}", func(r chi.Router) {
		r.Use(ArticleCtx)      // Load the *Article on the request context
		r.Get("/", GetArticle) // GET /articles/123
		// 	r.Put("/", UpdateArticle)    // PUT /articles/123
		// 	r.Delete("/", DeleteArticle) // DELETE /articles/123
	})

	// GET /articles/whats-up
	// r.With(ArticleCtx).Get("/{articleSlug:[a-z-]+}", GetArticle)
	//
	return r
}

func dbGetUser(id int64) (*User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("user not found.")
}

func NewUserPayloadResponse(user *User) *UserPayload {
	return &UserPayload{User: user}
}

func NewArticleResponse(article *Article) *ArticleResponse {
	resp := &ArticleResponse{Article: article}

	if resp.User == nil {
		if user, _ := dbGetUser(resp.UserID); user != nil {
			resp.User = NewUserPayloadResponse(user)
		}
	}

	return resp
}

func NewArticleListResponse(articles []*Article, pagination *Pagination, total int) *ArticleListResponse {
	resp := &ArticleListResponse{
		Data:  articles,
		Total: total,
		Page:  pagination.Page,
		Limit: pagination.Limit,
	}

	return resp
}

func ListArticles(w http.ResponseWriter, r *http.Request) {
	pagination := GetPagination(r)

	start := pagination.Offset
	end := start + pagination.Limit

	var paginatedArticles []*Article
	if start < len(articles) {
		if end > len(articles) {
			end = len(articles)
		}
		paginatedArticles = articles[start:end]
	}

	if err := render.Render(w, r, NewArticleListResponse(paginatedArticles, pagination, len(articles))); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func dbGetArticle(id string) (*Article, error) {
	for _, a := range articles {
		if a.ID == id {
			return a, nil
		}
	}
	return nil, errors.New("article not found.")
}

func dbGetArticleBySlug(slug string) (*Article, error) {
	for _, a := range articles {
		if a.Slug == slug {
			return a, nil
		}
	}
	return nil, errors.New("article not found.")
}

func ArticleCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var article *Article
		var err error

		if articleID := chi.URLParam(r, "articleID"); articleID != "" {
			article, err = dbGetArticle(articleID)
		} else if articleSlug := chi.URLParam(r, "articleSlug"); articleSlug != "" {
			article, err = dbGetArticleBySlug(articleSlug)
		} else {
			render.Render(w, r, ErrNotFound)
			return
		}
		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "article", article)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	article := r.Context().Value("article").(*Article)

	if err := render.Render(w, r, NewArticleResponse(article)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func dbNewArticle(article *Article) (string, error) {
	article.ID = fmt.Sprintf("%d", rand.Intn(100)+10)
	articles = append(articles, article)
	return article.ID, nil
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	data := &ArticleRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	article := data.Article
	dbNewArticle(article)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewArticleResponse(article))
}
