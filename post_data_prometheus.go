package main

type postDataPrometheus struct {
	Alerts []struct {
		Annotations struct {
			Site string `json:"site`
		} `json:"annotations"`
	} `json:"alerts"`
}
