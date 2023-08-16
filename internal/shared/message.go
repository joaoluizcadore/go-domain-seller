package shared

type Message struct {
	Name    string `form:"name" json:"name" binding:"required"`
	Email   string `form:"email" json:"email" binding:"required"`
	Message string `form:"message" json:"message" binding:"required"`
	Phone   string `form:"phone" json:"phone" binding:"required"`
	Domain  string `form:"domain" json:"domain"`
}
