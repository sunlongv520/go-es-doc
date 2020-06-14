package Models

type SearchModel struct {
	BookName string `json:"book_name"`
	BookPress string `json:"book_press"`
	BookPrice1Start float32 `json:"book_price1_start"`
	BookPrice1End float32 `json:"book_price1_end"`
}

func NewSearchModel() *SearchModel  {
	return &SearchModel{}
}
