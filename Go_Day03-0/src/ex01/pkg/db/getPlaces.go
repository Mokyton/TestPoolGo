package db

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type Place struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type types struct {
	Places []Place `json:"_source"`
}

func GetPlaces(limit int, offset int) ([]Place, int, error) {
	data := types{}
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, 0, err
	}
	req := esapi.SearchRequest{
		Index:        []string{"places"},
		DocumentType: []string{"place"},
		Size:         &limit,
		From:         &offset,
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return nil, 0, err
	}

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, 0, err
	}

	res, err = es.Count()
	if err != nil {
		return nil, 0, err
	}

	total := struct {
		Total int `json:"count"`
	}{}
	if err := json.NewDecoder(res.Body).Decode(&total); err != nil {
		return nil, 0, err
	}

	return data.Places, total.Total, nil
}

func (p *types) UnmarshalJSON(data []byte) error {
	tmpl := struct {
		Hits struct {
			Hits []struct {
				Source struct {
					Name    string `json:"name"`
					Address string `json:"address"`
					Phone   string `json:"phone"`
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
