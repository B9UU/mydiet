package search

type DummyJson struct {
	Products []Product `json:"products"`
}

type Product struct {
	Title string `json:"title"`
}
