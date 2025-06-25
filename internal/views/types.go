package views

type UpdateViewMessage int

type FailedRequest struct {
	err error
}
type SuccessRequest struct {
	suggetions []string
}

type DummyJson struct {
	Products []Product `json:"products"`
}

type Product struct {
	Title string `json:"title"`
}
