package product

type CreationOptions struct {
	Name        string `json:"name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Description string `json:"description" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
}

type UpdateOptions struct {
	Price       int    `json:"price" binding:"required"`
	Description string `json:"description" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
}
