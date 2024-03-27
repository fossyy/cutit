package alias

import (
	"github.com/fossyy/cutit/db"
	errorView "github.com/fossyy/cutit/view/error"
	"github.com/gorilla/mux"
	"net/http"
)

func ALL(w http.ResponseWriter, r *http.Request) {
	alias := mux.Vars(r)
	var links db.Link
	if err := db.DB.Table("links").Where("alias = ?", alias["alias"]).First(&links).Error; err != nil {
		component := errorView.Main("Not Found")
		component.Render(r.Context(), w)
	}
	http.Redirect(w, r, links.URL, http.StatusSeeOther)
}
