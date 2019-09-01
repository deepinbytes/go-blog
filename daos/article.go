package daos

import (
	"github.com/deepinbytes/go-blog/app"
	"github.com/deepinbytes/go-blog/models"
)

// ArticleDAO persists article data in database
type ArticleDAO struct{}

// NewArticleDAO creates a new ArticleDAO
func NewArticleDAO() *ArticleDAO {
	return &ArticleDAO{}
}

// Get reads the article with the specified ID from the database.
func (dao *ArticleDAO) Get(rs app.RequestScope, id int) (*models.Article, error) {
	var article models.Article
	err := rs.Tx().Select().Model(id, &article)
	return &article, err
}

// Create saves a new article record in the database.
// The Article.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *ArticleDAO) Create(rs app.RequestScope, article *models.Article) error {
	article.Id = 0
	return rs.Tx().Model(article).Insert()
}

// Update saves the changes to an article n the database.
func (dao *ArticleDAO) Update(rs app.RequestScope, id int, article *models.Article) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	article.Id = id
	return rs.Tx().Model(article).Exclude("Id").Update()
}

// Delete deletes an article with the specified ID from the database.
func (dao *ArticleDAO) Delete(rs app.RequestScope, id int) error {
	article, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(article).Delete()
}

// Count returns the number of the article records in the database.
func (dao *ArticleDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("article").Row(&count)
	return count, err
}

// Query retrieves the article records with the specified offset and limit from the database.
func (dao *ArticleDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Article, error) {
	articles := []models.Article{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&articles)
	return articles, err
}
