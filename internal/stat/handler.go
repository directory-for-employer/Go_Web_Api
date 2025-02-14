package stat

import (
	"go/web-api/pkg/req"
	"go/web-api/pkg/res"
	"net/http"
	"time"
)

type StatHandlerDeps struct {
}

type StatHandler struct {
}

func NewStatHandler(router *http.ServeMux, deps *StatHandlerDeps) {
	handler := &StatHandler{}
	router.HandleFunc("GET /stat", handler.Statistic())
}

type GetParamData struct {
	From string
	To   time.Time
	By   string
}

func (handler *StatHandler) Statistic() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from := req.DecodeStringQuery(r, "from")
		to, err := req.DecodeTimeQuery(r, "to")
		if err != nil {
			res.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		by := req.DecodeStringQuery(r, "by")
		res.Json(w, GetParamData{By: *by, To: *to, From: *from}, http.StatusOK)
	}
}
