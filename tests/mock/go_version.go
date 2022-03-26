package mock

import (
	"github.com/qa-tools-family/go-practices-in-actions/tests/mock/spider"
)

func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}
