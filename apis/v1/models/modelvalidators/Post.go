package modelvalidators

import (
	"gin-app/apis/v1/config/database"
	"gin-app/apis/v1/models"

	"github.com/go-playground/validator/v10"
)

func ValidatePostValuer(sl validator.StructLevel) {
	post := sl.Current().Interface().(models.Post)
	validateAssociatedUser(sl, post)
}

func validateAssociatedUser(sl validator.StructLevel, post models.Post) {
	u := models.User{}
	db := database.GetDB()
	_, err := u.GetUser(db, post.UserID)

	if err != nil {
		sl.ReportError(post.UserID, "user_id", "UserID", "UserIdDoesntExists", "Entered User ID doesn't exists.")
	}
}
