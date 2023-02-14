package address

type CreationOptions struct {
	UserID        string `json:"userID" binding:"required"`
	Country       string `json:"country" binding:"required"`
	City          string `json:"city" binding:"required"`
	StreetAddress string `json:"streetAddress" binding:"required"`
}
