package mock

import (
	"github.com/golang/mock/gomock"
	spider "github.com/qa-tools-family/go-practices-in-actions/tests/mock/spider/mock"
	"testing"
)

func TestGetGoVersion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSpider := spider.NewMockSpider(ctrl)
	// 生成了一个 mock 对象，它支持处理 GetBody 方法并且返回 go1.8.3
	mockSpider.EXPECT().GetBody().Return("go1.8.3")
	goVer := GetGoVersion(mockSpider)

	if goVer != "go1.8.3" {
		t.Errorf("Get wrong version %s", goVer)
	}
}

func TestGetName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSpider := spider.NewMockSpider(ctrl)
	// 生成了一个 mock 对象，它支持处理 GetBody 方法并且返回 go1.8.3
	mockSpider.EXPECT().GetName(gomock.Eq("123")).Return("go1.8.3")
	mockSpider.EXPECT().GetName(gomock.Eq("1234")).Return("go1.9.3")
	t.Logf(mockSpider.GetName("1234"))
	t.Logf(mockSpider.GetName("123"))
}