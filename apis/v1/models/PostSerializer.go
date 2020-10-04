package models

type PostSerializer struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
}

type PostWithAssociationSerializer struct {
	PostSerializer

	User User `gorm:"ForeignKey:UserID,references:ID" json:"user"`
}
