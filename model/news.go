package model

// News .
type News struct {
	ID       string `xorm:"unique not null text" json:"id"`
	Title    string `xorm:"text" json:"title"`
	Abstract string `xorm:"text" json:"abstract"`
	Content  string `xorm:"text" json:"content"`
	Source   string `xorm:"text" json:"source"`
	Date     string `xorm:"text" json:"date"`
}
