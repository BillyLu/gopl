package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func string2dollars(s string) (d dollars, err error) {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		err = fmt.Errorf("can not convert %s to dollars", s)
		return
	}
	d = dollars(f)
	return
}

type priceDB struct {
	sync.Mutex
	db map[string]dollars
}

func (p *priceDB) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range p.db {
		fmt.Fprintf(w, "%s: %s \n", item, price)
	}
}

func (p *priceDB) update(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	item := params.Get("item")
	price := params.Get("price")

	d, err := string2dollars(price)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "invalid price")
		return
	}

	// lock before modify, release after returns
	p.Lock()
	defer p.Unlock()

	p.db[item] = d
	time.Sleep(time.Duration(time.Second * 20))
	fmt.Fprintf(w, "%s: %s", item, p.db[item])
}

func main() {
	p := priceDB{
		db: map[string]dollars{"shoes": 50, "socks": 5},
	}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(p.list))
	mux.HandleFunc("/update", p.update)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
