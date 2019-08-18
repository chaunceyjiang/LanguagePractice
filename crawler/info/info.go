package info

type Hd struct {
	Url string
	Title []string
}
type Bd struct {
	Described string
	Rating string
	Num string
	Quote string
}
type Movie struct {
	Hd
	Bd
}

func NewMovie() *Movie {
	return &Movie{Hd{Title:[]string{}},Bd{}}
}