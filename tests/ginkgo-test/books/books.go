package books


const CategoryNovel = 0
const CategoryShortStory = 1


type Book struct {
	Title string
	Author string
	Pages int
	CategoryNovel string
}

func (b Book) Category() int {
	if b.Pages > 300 {
		return CategoryNovel
	}
	return CategoryShortStory
}
