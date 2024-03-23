package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gin-gonic/gin"
)

var cache = memcache.New("localhost:11211")

type Profile struct {
	Name  string
	Email string
}

func getProfileFromCache(profileID string) (Profile, bool) {
	item, err := cache.Get(profileID)
	if err != nil {
		log.Println(err)
		return Profile{}, false
	}

	var p Profile

	err = json.Unmarshal(item.Value, &p)
	if err != nil {
		log.Println(err)
		return Profile{}, false
	}

	return p, true
}

func saveProfileToCache(key string, p Profile) error {
	bb, err := json.Marshal(p)
	if err != nil {
		return err
	}

	err = cache.Set(&memcache.Item{
		Key:        key,
		Value:      bb,
		Expiration: int32(time.Now().Add(time.Hour * 30).Unix()),
	})
	if err != nil {
		return err
	}

	return nil
}

func getProfile(c *gin.Context) {
	defer func(d time.Time) {
		log.Printf("Execution time: %fs", time.Since(d).Seconds())
	}(time.Now())

	id := c.Param("profileID")
	var p Profile

	p, ok := getProfileFromCache(id)
	if ok {
		c.JSON(200, p)
		return
	}

	log.Println("getting data from DB")
	time.Sleep(time.Millisecond * 600)

	p = Profile{
		Name:  fmt.Sprintf("Jhon_%s", id),
		Email: fmt.Sprintf("%s@example.com", id),
	}

	go saveProfileToCache(id, p)
	c.JSON(200, p)
}

func main() {
	r := gin.New()
	r.GET("/profiles/:profileID", getProfile)
	r.Run(":8080")
}
