package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

var cache = memcache.New("localhost:11211")

type Profile struct {
	Name  string
	Email string
}

func logexecutionTime(d time.Time) {
	fmt.Printf("Execution time: %fs\n", time.Since(d).Seconds())
}

func saveProfileToCache(id string, p Profile) error {
	bb, err := json.Marshal(p)
	if err != nil {
		return err
	}

	return cache.Set(&memcache.Item{
		Key:        id,
		Value:      bb,
		Expiration: 5,
	})
}

func getProfileFromCache(key string) Profile {
	var p Profile

	item, err := cache.Get(key)
	if err != nil {
		fmt.Println(err)
		return Profile{}
	}

	err = json.Unmarshal(item.Value, &p)
	if err != nil {
		fmt.Println(err)
		return Profile{}
	}

	return p
}

func handleGetProfile(w http.ResponseWriter, req *http.Request) {
	defer logexecutionTime(time.Now())

	id := req.PathValue("id")

	fmt.Println("getting data from cache")
	p := getProfileFromCache(id)
	if p.Email != "" {
		fmt.Fprintf(w, "%+v", p)
		return
	}

	fmt.Println("getting data from DB")
	time.Sleep(time.Millisecond * 600)

	p = Profile{
		Name:  fmt.Sprintf("Jhon_%s", id),
		Email: fmt.Sprintf("%s@example.com", id),
	}

	saveProfileToCache(id, p)
	fmt.Fprintf(w, "%+v", p)
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/users/{id}", handleGetProfile)
	http.ListenAndServe(":8080", r)
}
