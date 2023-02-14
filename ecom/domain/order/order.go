package order

type CreationOptions struct {
	ProductID string `json:"productID" binding:"required"`
	UserID    string `json:"userID" binding:"required"`
	AddressID string `json:"addressID" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}
