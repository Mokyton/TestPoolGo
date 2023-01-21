package db

type Place struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Location struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"location"`
}

type types struct {
	Places []Place `json:"_source"`
}

func GetPlaces(limit int, offset int) ([]Place, int, error) {

}
