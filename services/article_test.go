package services

import (
	"errors"
	"testing"

	"github.com/deepinbytes/go-blog/app"
	"github.com/deepinbytes/go-blog/models"
	"github.com/stretchr/testify/assert"
)

func TestNewArticleService(t *testing.T) {
	dao := newMockArticleDAO()
	s := NewArticleService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestArticleService_Get(t *testing.T) {
	s := NewArticleService(newMockArticleDAO())
	article, err := s.Get(nil, 1)
	if assert.Nil(t, err) && assert.NotNil(t, article) {
		assert.Equal(t, "aaa", article.Title)
	}

	article, err = s.Get(nil, 100)
	assert.NotNil(t, err)
}

func TestArticleService_Create(t *testing.T) {
	s := NewArticleService(newMockArticleDAO())
	article, err := s.Create(nil, &models.Article{
		Title:   "ddd",
		Content: "test",
		Author:  "test",
	})
	if assert.Nil(t, err) && assert.NotNil(t, article) {
		assert.Equal(t, 4, article.Id)
		assert.Equal(t, "ddd", article.Title)
		assert.Equal(t, "test", article.Author)
		assert.Equal(t, "test", article.Content)
	}

	// dao error
	_, err = s.Create(nil, &models.Article{
		Id:      100,
		Title:   "ddd",
		Content: "test",
		Author:  "test",
	})
	assert.NotNil(t, err)

	// validation error
	_, err = s.Create(nil, &models.Article{
		Title: "",
	})
	assert.NotNil(t, err)
}

func TestArticleService_Update(t *testing.T) {
	s := NewArticleService(newMockArticleDAO())
	article, err := s.Update(nil, 2, &models.Article{
		Title:   "ddd",
		Content: "test",
		Author:  "test",
	})
	if assert.Nil(t, err) && assert.NotNil(t, article) {
		assert.Equal(t, 2, article.Id)
		assert.Equal(t, "ddd", article.Title)
		assert.Equal(t, "test", article.Author)
		assert.Equal(t, "test", article.Content)
	}

	// dao error
	_, err = s.Update(nil, 100, &models.Article{
		Title:   "ddd",
		Content: "test",
		Author:  "test",
	})
	assert.NotNil(t, err)

	// validation error
	_, err = s.Update(nil, 2, &models.Article{
		Title: "",
	})
	assert.NotNil(t, err)
}

func TestArticleService_Delete(t *testing.T) {
	s := NewArticleService(newMockArticleDAO())
	article, err := s.Delete(nil, 2)
	if assert.Nil(t, err) && assert.NotNil(t, article) {
		assert.Equal(t, 2, article.Id)
		assert.Equal(t, "bbb", article.Title)
	}

	_, err = s.Delete(nil, 2)
	assert.NotNil(t, err)
}

func TestArticleService_Query(t *testing.T) {
	s := NewArticleService(newMockArticleDAO())
	result, err := s.Query(nil, 1, 2)
	if assert.Nil(t, err) {
		assert.Equal(t, 2, len(result))
	}
}

func newMockArticleDAO() articleDAO {
	return &mockArticleDAO{
		records: []models.Article{
			{Id: 1, Title: "aaa"},
			{Id: 2, Title: "bbb"},
			{Id: 3, Title: "ccc"},
		},
	}
}

type mockArticleDAO struct {
	records []models.Article
}

func (m *mockArticleDAO) Get(rs app.RequestScope, id int) (*models.Article, error) {
	for _, record := range m.records {
		if record.Id == id {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *mockArticleDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Article, error) {
	return m.records[offset : offset+limit], nil
}

func (m *mockArticleDAO) Count(rs app.RequestScope) (int, error) {
	return len(m.records), nil
}

func (m *mockArticleDAO) Create(rs app.RequestScope, article *models.Article) error {
	if article.Id != 0 {
		return errors.New("Id cannot be set")
	}
	article.Id = len(m.records) + 1
	m.records = append(m.records, *article)
	return nil
}

func (m *mockArticleDAO) Update(rs app.RequestScope, id int, article *models.Article) error {
	article.Id = id
	for i, record := range m.records {
		if record.Id == id {
			m.records[i] = *article
			return nil
		}
	}
	return errors.New("not found")
}

func (m *mockArticleDAO) Delete(rs app.RequestScope, id int) error {
	for i, record := range m.records {
		if record.Id == id {
			m.records = append(m.records[:i], m.records[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
