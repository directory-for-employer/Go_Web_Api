package stat

import (
	"go/web-api/configs"
	"go/web-api/pkg/middleware"
	"go/web-api/pkg/req"
	"go/web-api/pkg/res"
	"net/http"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

type StatHandler struct {
	StatRepository *StatRepository
}

func NewStatHandler(router *http.ServeMux, deps *StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}
	router.Handle("GET /stat", middleware.IsAuthed(handler.GetStat(), deps.Config))
}

func (handler *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := req.DecodeTimeQuery(r, "from")
		if err != nil {
			http.Error(w, "Invalid Format Date", http.StatusBadRequest)
			return
		}
		to, err := req.DecodeTimeQuery(r, "to")
		if err != nil {
			http.Error(w, "Invalid Format Date", http.StatusBadRequest)
			return
		}
		by := req.DecodeStringQuery(r, "by")
		if *by != GroupByDay && *by != GroupByMonth {
			http.Error(w, "Param not day or month", http.StatusBadRequest)
			return
		}
		stats := handler.StatRepository.GetStat(*by, *from, *to)

		res.Json(w, stats, http.StatusOK)
	}
}
