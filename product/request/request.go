package request

type CreateProductRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Rating      int    `json:"rating" validate:"required"`
	Image       string `json:"image" validate:"required"`
}

type UpdateProductRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
	Image       string `json:"image"`
}
