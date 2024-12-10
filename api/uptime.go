package handler

import (
	"net/http"

	"asynclab.club/AsyncFunction/pkg/lib/uptime"
	"asynclab.club/AsyncFunction/pkg/util"
)

func HandlerUptime(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		util.HttpResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	params, err := util.ReadParamsFromQuery[uptime.QueryParams](r.URL.Query())
	if err != nil {
		util.HttpResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := uptime.GetMetricsFromUptime(params.BaseUrl, params.ApiKey)
	if err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	statuses, err := uptime.Extract(data)
	if err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.HttpResponse(w, http.StatusOK, statuses)
}
