package http_mock

import (
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFetchArticles(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// our database of articles
	articles := make([]map[string]interface{}, 0)

	// mock to list out the articles
	httpmock.RegisterResponder("GET", "http://api.missshi.com/articles",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, articles)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	// return an article related to the request with the help of regexp submatch (\d+)
	httpmock.RegisterResponder("GET", `=~^http://api\.missshi\.com/articles/id/(\d+)\z`,
		func(req *http.Request) (*http.Response, error) {
			// Get ID from request
			id := httpmock.MustGetSubmatchAsUint(req, 1) // 1=first regexp submatch
			return httpmock.NewJsonResponse(200, map[string]interface{}{
				"id":   id,
				"name": "My Great Article",
			})
		},
	)

	// mock to add a new article
	httpmock.RegisterResponder("POST", "http://api.missshi.com/articles",
		func(req *http.Request) (*http.Response, error) {
			article := make(map[string]interface{})
			if err := json.NewDecoder(req.Body).Decode(&article); err != nil {
				return httpmock.NewStringResponse(400, ""), nil
			}

			articles = append(articles, article)

			resp, err := httpmock.NewJsonResponse(200, article)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)
	response2 := FetchActicle("http://easyenv.baidu-int.com/health")
	t.Log(response2)
	response := FetchActicle("http://api.missshi.com/articles")
	assert.Equal(t, response, "[]")
	response1 := FetchActicle("http://api.missshi.com/articles/id/123123")
	assert.Equal(t, response1, "{\"id\":123123,\"name\":\"My Great Article\"}")

	// get count info
	httpmock.GetTotalCallCount()

	// get the amount of calls for the registered responder
	info := httpmock.GetCallCountInfo()
	t.Log(info)
}
