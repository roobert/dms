package main

type postDataPrometheus struct {
	Alerts []struct {
		Annotations struct {
			Site string `json:"site`
		} `json:"annotations"`
	} `json:"alerts"`
}

func (pdp postDataPrometheus) Site() string {
	return pdp.Alerts[0].Annotations.Site
}
