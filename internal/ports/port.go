package ports

type Port struct {
	ID          string    `json:"-"`
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Code        string    `json:"code"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Unlocs      []string  `json:"unlocs"`
}
