package http_mock

import (
	"fmt"
	"github.com/guonaihong/gout"
)

type Body struct {
	BodyType string `json:"type"`
	Content string `json:"content"`
}

type Message struct {
	Message Body `json:"message"`
}

func FetchActicle(url string) string {
	queryData := map[string]string{
		"access_token": "1234567890",
	}
	requestData := Message{Message: Body{
			BodyType: "TEXT",
			Content: "article1",
	}}
	var responseData string
	_ = gout.
		GET(url).
		SetQuery(queryData).
		SetJSON(&requestData).
		BindBody(&responseData).
		Do()
	fmt.Println(responseData)
	return responseData
}
