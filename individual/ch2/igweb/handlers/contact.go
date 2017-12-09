package handlers

import (
	"net/http"

	"github.com/EngineerKamesh/igb/igweb/common"
	"github.com/isomorphicgo/isokit"
)

func ContactHandler(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		env.TemplateSet.Render("contact_page", &isokit.RenderParams{Writer: w, Data: nil, PageTitle: "Contact"})

	})
}
