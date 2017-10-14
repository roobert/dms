package handler

import (
	"encoding/json"
	"net/http"
	"time"

	. "github.com/roobert/dms/error"
	. "github.com/roobert/dms/server/client"
	. "github.com/roobert/dms/server/db"
	. "github.com/roobert/dms/server/post_data"
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
