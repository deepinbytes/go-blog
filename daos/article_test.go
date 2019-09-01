package daos

import (
	"testing"

	"github.com/deepinbytes/go-blog/app"
	"github.com/deepinbytes/go-blog/models"
	"github.com/deepinbytes/go-blog/testdata"
	"github.com/stretchr/testify/assert"
)

func TestArticleDAO(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewArticleDAO()

	{
		// Get
		testDBCall(db, func(rs app.RequestScope) {
			article, err := dao.Get(rs, 2)
			assert.Nil(t, err)
			if assert.NotNil(t, article) {
				assert.Equal(t, 2, article.Id)
			}
		})
	}

	{
		// Create
		testDBCall(db, func(rs app.RequestScope) {
			article := &models.Article{
				Id:    1000,
				Title: "tester",
			}
			err := dao.Create(rs, article)
			assert.Nil(t, err)
			assert.NotEqual(t, 1000, article.Id)
			assert.NotZero(t, article.Id)
		})
	}

	{
		// Update
		testDBCall(db, func(rs app.RequestScope) {
			article := &models.Article{
				Id:    2,
				Title: "tester",
			}
			err := dao.Update(rs, article.Id, article)
			assert.Nil(t, err)
		})
	}

	{
		// Update with error
		testDBCall(db, func(rs app.RequestScope) {
			article := &models.Article{
				Id:    2,
				Title: "tester",
			}
			err := dao.Update(rs, 99999, article)
			assert.NotNil(t, err)
		})
	}

	{
		// Delete
		testDBCall(db, func(rs app.RequestScope) {
			err := dao.Delete(rs, 2)
			assert.Nil(t, err)
		})
	}

	{
		// Delete with error
		testDBCall(db, func(rs app.RequestScope) {
			err := dao.Delete(rs, 99999)
			assert.NotNil(t, err)
		})
	}

	{
		// Query
		testDBCall(db, func(rs app.RequestScope) {
			articles, err := dao.Query(rs, 1, 3)
			assert.Nil(t, err)
			assert.Equal(t, 3, len(articles))
		})
	}

	{
		// Count
		testDBCall(db, func(rs app.RequestScope) {
			count, err := dao.Count(rs)
			assert.Nil(t, err)
			assert.NotZero(t, count)
		})
	}
}
