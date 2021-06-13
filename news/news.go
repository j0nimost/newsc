package news

type Newser interface {
	GetNews() map[int] string
}