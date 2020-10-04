package models

type UserSerializer struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type UserWithAssociationSerializer struct {
	UserSerializer

	Posts []Post `gorm:"ForeignKey:UserID" json:"posts"`
}
