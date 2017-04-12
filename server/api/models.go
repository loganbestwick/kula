package api

type Workout struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Series      string `json:"series"`
	Category    string `json:"category"`
	SubCategory string `json:"subcategory"`
	Duration    int    `json:"duration"`
}
