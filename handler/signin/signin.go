package signinHandler

import (
	"github.com/fossyy/cutit/app"
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

	user, err := app.Server.Database.GetUser(email)
	if err != nil {
		component := signinView.Main("Sign in Page", types.Message{
			Code:    0,
			Message: "Database error : " + err.Error(),
		})
		component.Render(r.Context(), w)
		return
	}

	if email == user.Email && utils.CheckPasswordHash(password, user.Password) {
		session.Values["user"] = types.User{
			UserID:        user.UserID,
			Email:         email,
			Username:      user.Username,
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
