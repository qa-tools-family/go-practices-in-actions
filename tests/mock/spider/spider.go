package spider

type Spider interface {
	GetBody() string
	GetName(string) string
}
