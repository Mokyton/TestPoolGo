package db

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/olivere/elastic"
	"strconv"
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

type Geo struct {
	Lat float64
	Lon float64
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

func (p *Types) GetPlaces(limit int, lat string, lon string) ([]Place, error) {
	var buf bytes.Buffer
	if lat == "" || lon == "" {
		return nil, errors.New("Empty lat and lon")
	}

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}

	latFloat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return nil, err
	}

	lonFloat, err := strconv.ParseFloat(lon, 64)
	if err != nil {
		return nil, err
	}
	sort := SortJson{Sort: []Sort{{
		GeoDistance{
			Location:       elastic.GeoPoint(Location{Lon: lonFloat, Lat: latFloat}),
			Order:          "asc",
			Unit:           "km",
			Mode:           "min",
			DistanceType:   "arc",
			IgnoreUnmapped: true,
		},
	}}}

	if err = json.NewEncoder(&buf).Encode(sort); err != nil {
		return nil, err
	}

	req := esapi.SearchRequest{
		Index:        []string{"places"},
		DocumentType: []string{"place"},
		Size:         &limit,
		Body:         &buf,
	}

	res, err := req.Do(context.Background(), es)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	if err = json.NewDecoder(res.Body).Decode(p); err != nil {
		return nil, err
	}

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
