package model

// Todo -
type Todo struct {
	ID   int    `json:"id" gorm:"primary_key,auto_increment"`
	Text string `json:"text"`
	Done bool   `json:"done"`

	// -----------
	BaseModel
}
