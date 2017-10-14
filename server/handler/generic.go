package handler

//
//import (
//	"encoding/json"
//	"net/http"
//	"time"
//)
//
//func GenericHandler(w http.ResponseWriter, r *http.Request) {
//	c := Client{Name: getSiteNameGenericJSON(r), TimeStamp: time.Now()}
//
//	updateClient(c)
//	updateResult(c)
//}
//
//func getSiteNameGenericJSON(r *http.Request) string {
//	decoder := json.NewDecoder(r.Body)
//
//	var pdg postDataGeneric
//
//	err := decoder.Decode(&pdg)
//	checkErr(err)
//
//	return pdg.Site
//}
