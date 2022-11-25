package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func display_time(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		date_1 := time.Now()
		date := date_1.Hour()
		str_1 := strconv.Itoa(date)

		date = date_1.Minute()
		str_2 := strconv.Itoa(date)
		fmt.Fprintf(w, "%s h %s", str_1, str_2)
	default:
		fmt.Fprintf(w, "Sorry, only GET methods are supported.")
	}
}

func display_D1000(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		rand.Seed(time.Now().UnixNano())

		min := 1
		max := 1000
		nb := rand.Intn(max-min) + min

		fmt.Fprintf(w, "%04d", nb)
	default:
		fmt.Fprintf(w, "Sorry, only GET methods are supported.")
	}

}

func display_dices(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:

		id := req.URL.Query().Get("type")

		dice := []int{4, 6, 8, 10, 12, 20, 100}

		rand.Seed(time.Now().UnixNano())

		if id == "d6" {
			min := 1
			max := 6
			nb := rand.Intn(max-min) + min
			fmt.Fprintf(w, "%d ", nb)
		} else {

			for i := 0; i < 15; i++ {
				min := 1
				nb := rand.Intn(dice[rand.Intn(len(dice))]-min) + min

				fmt.Fprintf(w, "%d ", nb)
			}
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET methods are supported.")
	}
}

func randomize_words(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("sentence")
	switch req.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "hello")
	case http.MethodPost:
		str := " hi sjafij"
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length:", string(len(str)))

	default:
		fmt.Fprintf(w, "Sorry, only GET methods are supported.")
	}
}

func semi_capitalize(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("sentence")
	switch req.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "hello")
	case http.MethodPost:
		str := " hi sjafij"
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	default:
		fmt.Fprintf(w, "Sorry, only GET methods are supported.")
	}
}
func main() {
	http.HandleFunc("/", display_time)
	http.HandleFunc("/dice", display_D1000)
	http.HandleFunc("/dices", display_dices)
	http.HandleFunc("/randomize-words", randomize_words)
	http.HandleFunc("/semi-capitalize-sentence", semi_capitalize)

	http.ListenAndServe(":4567", nil)
}
