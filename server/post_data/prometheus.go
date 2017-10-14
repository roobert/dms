package post_data

type PostDataPrometheus struct {
	Alerts []struct {
		Annotations struct {
			Site string `json:"site`
		} `json:"annotations"`
	} `json:"alerts"`
}

func (pdp PostDataPrometheus) Site() string {
	return pdp.Alerts[0].Annotations.Site
}
