package apis

import (
	"github.com/deepinbytes/go-blog/app"
	"github.com/deepinbytes/go-blog/models"
	"github.com/go-ozzo/ozzo-routing"
	"strconv"
)

type (
	// articleService specifies the interface for the article service needed by articleResource.
	articleService interface {
		Get(rs app.RequestScope, id int) (*models.Article, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Article, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.Article) (*models.Article, error)
		Update(rs app.RequestScope, id int, model *models.Article) (*models.Article, error)
		Delete(rs app.RequestScope, id int) (*models.Article, error)
	}

	// articleResource defines the handlers for the CRUD APIs.
	articleResource struct {
		service articleService
	}
)

// ServeArticle sets up the routing of article endpoints and the corresponding handlers.
func ServeArticleResource(rg *routing.RouteGroup, service articleService) {
	r := &articleResource{service}
	rg.Get("/articles/<id>", r.get)
	rg.Get("/articles", r.query)
	rg.Post("/articles", r.create)
	rg.Put("/articles/<id>", r.update)
	rg.Delete("/articles/<id>", r.delete)
}

func (r *articleResource) get(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Get(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}
	data := map[string]interface{}{"status": "201", "message": "ok", "data": response}
	return c.Write(data)
}

func (r *articleResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	count, err := r.service.Count(rs)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		return err
	}
	//paginatedList.Items = items
	data := map[string]interface{}{"status": "201", "message": "ok", "data": items}
	return c.Write(data)
	//return c.Write(paginatedList)
}

func (r *articleResource) create(c *routing.Context) error {
	var model models.Article
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	data := map[string]interface{}{"status": "201", "message": "ok", "data": response}
	if err != nil {
		return err
	}
	return c.Write(data)
}

func (r *articleResource) update(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	rs := app.GetRequestScope(c)

	model, err := r.service.Get(rs, id)
	if err != nil {
		return err
	}

	if err := c.Read(model); err != nil {
		return err
	}

	response, err := r.service.Update(rs, id, model)
	if err != nil {
		return err
	}
	data := map[string]interface{}{"status": "201", "message": "ok", "data": response}
	return c.Write(data)
}

func (r *articleResource) delete(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Delete(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}
	data := map[string]interface{}{"status": "201", "message": "ok", "data": response}
	return c.Write(data)
}
