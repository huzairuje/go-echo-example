package request

type CreateProductRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
	Image       string `json:"image"`
}

type UpdateProductRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
	Image       string `json:"image"`
}
