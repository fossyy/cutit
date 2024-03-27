package indexHandler

import (
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

	var links []db.Link
	db.DB.Table("links").Where("owner_id = ?", userSession.UserID).Find(&links)

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
	err := db.DB.Create(&link).Error
	if err != nil {
		var links []db.Link
		component := indexView.Main("main page", types.Message{
			Code:    0,
			Message: "Error : " + err.Error(),
		}, links, host)
		component.Render(r.Context(), w)
		return
	}

	var links []db.Link
	db.DB.Table("links").Where("owner_id = ?", userSession.UserID).Find(&links)

	component := indexView.Main("main page", types.Message{
		Code:    1,
		Message: "Short url berhasil dibuat dengan alias " + alias,
	}, links, host)
	component.Render(r.Context(), w)
}
