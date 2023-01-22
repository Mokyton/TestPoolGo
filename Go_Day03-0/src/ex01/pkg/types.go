package pkg

type Store interface {
	GetPlaces(limit int, offset int) ([]Place, int, error)
}

type Place struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
