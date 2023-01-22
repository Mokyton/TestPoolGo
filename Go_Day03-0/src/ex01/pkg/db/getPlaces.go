package db

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"time"
)

type Place struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type Types struct {
	ScrollID string  `json:"_scroll_id"`
	Places   []Place `json:"_source"`
}

func New() *Types {
	places := make([]Place, 0, 10)
	return &Types{Places: places}
}

func (p *Types) GetPlaces(limit int, offset int) ([]Place, int, error) {
	//p := Types{}
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, 0, err
	}
	req := esapi.SearchRequest{
		Index:        []string{"places"},
		DocumentType: []string{"place"},
		Size:         &limit,
		Scroll:       1 * time.Minute,
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return nil, 0, err
	}

	if err := json.NewDecoder(res.Body).Decode(&p); err != nil {
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

	if total.Total <= (offset-1)*10 {
		return nil, total.Total, errors.New("the page doesn't exists")
	}

	for i := 1; i < offset; i++ {
		res, err := es.Scroll(es.Scroll.WithScrollID(p.ScrollID), es.Scroll.WithScroll(time.Minute))
		if err = json.NewDecoder(res.Body).Decode(&p); err != nil {
			return nil, 0, err
		}
	}

	return p.Places, total.Total, nil
}

func (p *Types) UnmarshalJSON(data []byte) error {
	p.Places = p.Places[:0]
	tmpl := struct {
		ScrollID string `json:"_scroll_id"`
		Hits     struct {
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
	p.ScrollID = tmpl.ScrollID
	for _, v := range tmpl.Hits.Hits {
		p.Places = append(p.Places, v.Source)
	}

	return nil
}
