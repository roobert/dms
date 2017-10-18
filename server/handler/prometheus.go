package handler

import (
	"encoding/json"
	"net/http"
	"time"

	. "github.com/roobert/dms/server/client"
	. "github.com/roobert/dms/server/postdata"
	. "github.com/roobert/golang-error"
)

func Prometheus(w http.ResponseWriter, r *http.Request) {
	c := Client{Name: getSiteNamePrometheusJSON(r), TimeStamp: time.Now()}

	UpdateClient(c)
	UpdateResult(c)
}

func getSiteNamePrometheusJSON(r *http.Request) string {
	decoder := json.NewDecoder(r.Body)

	var pdp PostDataPrometheus

	err := decoder.Decode(&pdp)
	CheckErr(err)

	return pdp.Site()
}
