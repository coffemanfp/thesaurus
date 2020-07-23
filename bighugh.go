package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BigHugh struct {
	APIKey string
}

type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

type words struct {
	Syn []string `json:"syn"`
}

func (b *BigHugh) Synonyms(term string) (syns []string, err error) {
	url := fmt.Sprintf(
		"http://words.bighugelabs.com/api/2/%s/%s/json",
		b.APIKey,
		term,
	)

	response, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf(
			"bighugh: Failed when looking for synonyms for \"%s\":\n%s",
			term,
			err,
		)
		return
	}
	defer response.Body.Close()
	var data synonyms

	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
		err = fmt.Errorf("bighugh: Failed when decoding response:\n%s", err)
		return
	}

	syns = append(syns, data.Noun.Syn...)
	syns = append(syns, data.Verb.Syn...)
	return
}
