package apis

import (
	"net/http"
	"testing"

	"github.com/deepinbytes/go-blog/daos"
	"github.com/deepinbytes/go-blog/services"
	"github.com/deepinbytes/go-blog/testdata"
)

func TestArticle(t *testing.T) {
	testdata.ResetDB()
	router := newRouter()
	ServeArticleResource(&router.RouteGroup, services.NewArticleService(daos.NewArticleDAO()))

	notFoundError := `{"status":404, "error_code":"NOT_FOUND", "message":"NOT_FOUND"}`
	FieldsRequiredError := `{"status":400,"error_code":"INVALID_DATA","message":"INVALID_DATA","details":[{"field":"author","error":"cannot be blank"},{"field":"content","error":"cannot be blank"},{"field":"title","error":"cannot be blank"}]}`
	TitleRequiredError := `{"status":400,"error_code":"INVALID_DATA","message":"INVALID_DATA","details":[{"field":"title","error":"cannot be blank"}]}`
	runAPITests(t, router, []apiTestCase{
		{"t1 - get an article", "GET", "/articles/1", "", http.StatusOK, `{"data":{"id":1,"title":"AC/DC","content":"test","author":"test"},"message":"ok","status":"201"}`},
		{"t2 - get a nonexisting article", "GET", "/articles/99999", "", http.StatusNotFound, notFoundError},
		{"t3 - create an article", "POST", "/articles", `{"Title":"Test","Content":"TestContent","Author":"TestAuthor"}`, http.StatusOK, `{"data":{"id":12,"title":"Test","content":"TestContent","author":"TestAuthor"},"message":"ok","status":"201"}`},
		{"t4 - create an article with validation error", "POST", "/articles", `{"Title":""}`, http.StatusBadRequest, FieldsRequiredError},
		{"t5 - update an article", "PUT", "/articles/2", `{"Title":"TestUpdated"}`, http.StatusOK, `{"data":{"id":2, "title":"TestUpdated", "content":"test", "author":"test"}, "message":"ok", "status":"201"}`},
		{"t6 - update an article with validation error", "PUT", "/articles/2", `{"Title":""}`, http.StatusBadRequest, TitleRequiredError},
		{"t7 - update a nonexisting article", "PUT", "/articles/99999", "{}", http.StatusNotFound, notFoundError},
		{"t8 - delete an article", "DELETE", "/articles/2", ``, http.StatusOK, `{"data":{"title":"TestUpdated", "content":"test", "author":"test", "id":2}, "message":"ok", "status":"201"}`},
		{"t9 - delete a nonexisting article", "DELETE", "/articles/99999", "", http.StatusNotFound, notFoundError},
		{"t10 - get a list of article", "GET", "/articles", "", http.StatusOK, `{"data":[{"id":1,"title":"AC/DC","content":"test","author":"test"},{"id":3,"title":"Aerosmith","content":"test","author":"test"},{"id":4,"title":"Alanis Morissette","content":"test","author":"test"},{"id":5,"title":"Alice In Chains","content":"test","author":"test"},{"id":6,"title":"Apocalyptica","content":"test","author":"test"},{"id":7,"title":"Audioslave","content":"test","author":"test"},{"id":8,"title":"BackBeat","content":"test","author":"test"},{"id":9,"title":"Billy Cobham","content":"test","author":"test"},{"id":10,"title":"Black Label Society","content":"test","author":"test"},{"id":11,"title":"Black Sabbath","content":"test","author":"test"},{"id":12,"title":"Test","content":"TestContent","author":"TestAuthor"}],"message":"ok","status":"201"}`},
	})
}
