package signinHandler

import (
	"github.com/fossyy/cutit/db"
	"github.com/fossyy/cutit/middleware"
	"github.com/fossyy/cutit/types"
	"github.com/fossyy/cutit/utils"
	signinView "github.com/fossyy/cutit/view/signin"
	"net/http"
)

func GET(w http.ResponseWriter, r *http.Request) {
	component := signinView.Main("Sign in Page", types.Message{
		Code:    3,
		Message: "",
	})
	component.Render(r.Context(), w)
}

func POST(w http.ResponseWriter, r *http.Request) {
	session, _ := middleware.Store.Get(r, "session")
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	var userData db.User

	if err := db.DB.Table("users").Where("email = ?", email).First(&userData).Error; err != nil {
		component := signinView.Main("Sign in Page", types.Message{
			Code:    0,
			Message: "Database error : " + err.Error(),
		})
		component.Render(r.Context(), w)
	}
	if email == userData.Email && utils.CheckPasswordHash(password, userData.Password) {
		session.Values["user"] = types.User{
			UserID:        userData.UserID,
			Email:         email,
			Username:      userData.Username,
			Authenticated: true,
		}
		err = session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	component := signinView.Main("Sign in Page", types.Message{
		Code:    0,
		Message: "User atau password salah",
	})
	component.Render(r.Context(), w)
}
