package database

type Workout struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Mgroup  string `json:"mgroup"`
	Ex_type string `json:"ex_type"`
}

type Friend struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Friends string `json:"friends"`
}

var default_leg = []Workout{
	{ID: "1", Name: "Leg Press", Mgroup: "leg", Ex_type: "leg"},
	{ID: "2", Name: "Leg Curl", Mgroup: "lot of legs", Ex_type: "leg"},
	{ID: "3", Name: "Squats", Mgroup: "very leg", Ex_type: "leg"},
}
