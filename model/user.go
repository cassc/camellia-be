package model

// User Feedback
type UserFeedback struct {
	ID       string `xorm:"unique not null text" json:"id"`
	Username string `xorm:"text" json:"username"`
	Message  string `xorm:"text" json:"message"`
	Phone    string `xorm:"text" json:"phone"`
	Email    string `xorm:"text" json:"email"`
}
