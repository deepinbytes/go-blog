package services

import (
	"github.com/deepinbytes/go-blog/app"
	"github.com/deepinbytes/go-blog/models"
)

// articleDAO specifies the interface of the article DAO needed by ArticleService.
type articleDAO interface {
	// Get returns the article with the specified article ID.
	Get(rs app.RequestScope, id int) (*models.Article, error)
	// Count returns the number of articles.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of articles with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.Article, error)
	// Create saves a new article in the storage.
	Create(rs app.RequestScope, article *models.Article) error
	// Update updates the article with given ID in the storage.
	Update(rs app.RequestScope, id int, article *models.Article) error
	// Delete removes the article with given ID from the storage.
	Delete(rs app.RequestScope, id int) error
}

// ArticleService provides services related with article.
type ArticleService struct {
	dao articleDAO
}

// NewArticleService creates a new ArticleService with the given article DAO.
func NewArticleService(dao articleDAO) *ArticleService {
	return &ArticleService{dao}
}

// Get returns the article with the specified the article ID.
func (s *ArticleService) Get(rs app.RequestScope, id int) (*models.Article, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new article.
func (s *ArticleService) Create(rs app.RequestScope, model *models.Article) (*models.Article, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}

// Updates the article with the specified ID.
func (s *ArticleService) Update(rs app.RequestScope, id int, model *models.Article) (*models.Article, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the article with the specified ID.
func (s *ArticleService) Delete(rs app.RequestScope, id int) (*models.Article, error) {
	article, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return article, err
}

// Count returns the number of articles.
func (s *ArticleService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the articles with the specified offset and limit.
func (s *ArticleService) Query(rs app.RequestScope, offset, limit int) ([]models.Article, error) {
	return s.dao.Query(rs, offset, limit)
}
