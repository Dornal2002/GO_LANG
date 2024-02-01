package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type response struct {
	Web_name []string `json:"websites"`
}

var mp map[string]string

func main() {
	mp = make(map[string]string)

	r := mux.NewRouter()

	r.HandleFunc("/websites", getdata).Methods(http.MethodGet)
	r.HandleFunc("/websites", createData).Methods(http.MethodPost)
	r.HandleFunc("/website", checkQuery).Methods(http.MethodGet)

	http.ListenAndServe("localhost:3000", r)
}

func getdata(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(mp)
	if err != nil {
		fmt.Println(err)
	}
}

func createData(w http.ResponseWriter, r *http.Request) {

	resp := response{}
	r.ParseForm()
	fmt.Fprint(w, r.Form)

	err := json.NewDecoder(r.Body).Decode(&resp)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprint(w, "Data posted Succesfully")
	for _, url := range resp.Web_name {
		go checkStatus(url)
	}

}

func checkStatus(url string) {
	for {
		res, err := http.Get(url)
		if err != nil {
			mp[url] = "DOWN"
		} else if res.StatusCode >= 200 && res.StatusCode < 300 {
			mp[url] = "UP"
		} else {
			mp[url] = "DOWN"
		}
		time.Sleep(1 * time.Minute)
	}
}

func checkQuery(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	m := r.Form
	val := m["websites"][0]

	res, ok := mp[val]
	if ok != true {
		fmt.Fprintf(w, "%v : DOWN", val)
		return
	}
	fmt.Fprintf(w, "%v is :%v", val, res)

}
