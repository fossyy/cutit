package indexHandler

import (
	"github.com/fossyy/cutit/app"
	"github.com/fossyy/cutit/db"
	"github.com/fossyy/cutit/middleware"
	"github.com/fossyy/cutit/types"
	"github.com/fossyy/cutit/utils"
	"github.com/fossyy/cutit/view/index"
	"net/http"
)

func GET(w http.ResponseWriter, r *http.Request) {
	session, _ := middleware.Store.Get(r, "session")
	userSession := middleware.GetUser(session)
	host := r.Host

	links, err := app.Server.Database.GetLinks(userSession.UserID.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	component := indexView.Main("main page", types.Message{
		Code:    3,
		Message: "",
	}, links, host)
	component.Render(r.Context(), w)
}

func POST(w http.ResponseWriter, r *http.Request) {
	session, _ := middleware.Store.Get(r, "session")
	userSession := middleware.GetUser(session)
	host := r.Host

	r.ParseForm()
	url := r.Form.Get("url")
	alias := utils.GenerateRandomString(10)
	link := db.Link{
		URL:     url,
		Alias:   alias,
		OwnerID: userSession.UserID,
	}
	err := app.Server.Database.CreateLink(&link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	links, err := app.Server.Database.GetLinks(userSession.UserID.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	component := indexView.Main("main page", types.Message{
		Code:    3,
		Message: "",
	}, links, host)
	component.Render(r.Context(), w)
}
