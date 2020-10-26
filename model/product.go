package model

// Product
type Product struct {
	ID       string `xorm:"unique not null text" json:"id"`
	Title    string `xorm:"text" json:"title"`
	Subtitle string `xorm:"text" json:"subtitle"`
	Wcover   string `xorm:"text" json:"wcover"`
	Hcover   string `xorm:"text" json:"hcover"`
}
