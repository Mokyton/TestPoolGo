package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/olivere/elastic"
)

type SortJson struct {
	Sort []Sort `json:"sort"`
}
type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
type GeoDistance struct {
	Location       elastic.GeoPoint `json:"location"`
	Order          string           `json:"order"`
	Unit           string           `json:"unit"`
	Mode           string           `json:"mode"`
	DistanceType   string           `json:"distance_type"`
	IgnoreUnmapped bool             `json:"ignore_unmapped"`
}
type Sort struct {
	GeoDistance GeoDistance `json:"_geo_distance"`
}

type Place struct {
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	Address  string           `json:"address"`
	Phone    string           `json:"phone"`
	Location elastic.GeoPoint `json:"location"`
}

type Types struct {
	Places []Place `json:"_source"`
}

func New() *Types {
	return &Types{}
}

func (p *Types) GetPlaces(limit int, lat float64, lon float64) ([]Place, error) {
	var buf bytes.Buffer

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}

	//sort := SortJson{Sort: []Sort{{
	//	GeoDistance{
	//		Location:       Location{Lon: lon, Lat: lat},
	//		Order:          "asc",
	//		Unit:           "km",
	//		Mode:           "min",
	//		DistanceType:   "arc",
	//		IgnoreUnmapped: true,
	//	},
	//}}}

	sort := map[string]interface{}{
		"sort": map[string]interface{}{
			"_geo_distance": map[string]interface{}{
				"location": map[string]interface{}{
					"lat": lat,
					"lon": lon,
				},
				"order":           "asc",
				"unit":            "km",
				"mode":            "min",
				"distance_type":   "arc",
				"ignore_unmapped": true,
			},
		},
	}

	if err = json.NewEncoder(&buf).Encode(sort); err != nil {
		return nil, err
	}

	//req := esapi.SearchRequest{
	//	Index:        []string{"places"},
	//	DocumentType: []string{"place"},
	//	Size:         &limit,
	//	Body:         &buf,
	//}
	//
	//res, err := req.Do(context.Background(), es)
	//if err != nil {
	//	return nil, err
	//}

	res, _ := es.Search(
		es.Search.WithIndex("places"),
		es.Search.WithSize(limit),
		es.Search.WithBody(&buf),
	)

	if err = json.NewDecoder(res.Body).Decode(p); err != nil {
		return nil, err
	}
	res.Body.Close()
	fmt.Println(p)
	return p.Places, nil
}

func (p *Types) UnmarshalJSON(data []byte) error {
	p.Places = p.Places[:0]
	tmpl := struct {
		Hits struct {
			Hits []struct {
				Source struct {
					ID       int              `json:"id"`
					Name     string           `json:"name"`
					Address  string           `json:"address"`
					Phone    string           `json:"phone"`
					Location elastic.GeoPoint `json:"location"`
				} `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}{}

	if err := json.Unmarshal(data, &tmpl); err != nil {
		return err
	}
	for _, v := range tmpl.Hits.Hits {
		p.Places = append(p.Places, v.Source)
	}

	return nil
}
