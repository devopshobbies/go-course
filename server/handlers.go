package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (handler *server) calculateBMI(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/bmi" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		defer r.Body.Close()
		rawBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err) // TODO
		}

		body := &struct {
			Height float64 `json:"height"`
			Weight float64 `json:"weight"`
		}{}

		err = json.Unmarshal(rawBody, body)
		if err != nil {
			panic(err) // TODO
		}

		if body.Height < 40 || body.Height > 250 {
			fmt.Fprintf(w, "invalid height provided: %.2f", body.Height)
			return
		}
		body.Height *= 0.01 // height should be in meter

		if body.Weight < 5 || body.Weight > 300 {
			fmt.Fprintf(w, "invalid weight provided: %.2f", body.Weight)
			return
		}

		bmi := body.Weight / (body.Height * body.Height)
		fmt.Fprintf(w, "your BMI is = %.2f\n", bmi)

	default:
		fmt.Fprintf(w, "Sorry, only POST method is supported.")
	}
}

func (handler *server) incrementCounter(w http.ResponseWriter, r *http.Request) {
	handler.mutex.Lock()
	handler.counter++
	fmt.Fprintf(w, strconv.Itoa(handler.counter))
	handler.mutex.Unlock()
}
