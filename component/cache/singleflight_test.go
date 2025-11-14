package cache

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"sync"
	"testing"
	"time"
)

func TestSingle(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)

	//模拟10个并发
	d := &userDemo{Id: 1}
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			result, err := SingleGet(d.GetKey(), d.GetFromCache, d.GetFromDB)
			fmt.Println(result, err)
		}()
	}
	wg.Wait()

}

type userDemo struct {
	redisClient redis.UniversalClient
	Id          int
}

func (d *userDemo) GetFromCache() (interface{}, error) {
	result, err := d.redisClient.Get(d.GetKey()).Result()
	if errors.Is(err, redis.Nil) {
		return nil, NotExist
	} else if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *userDemo) GetFromDB() (interface{}, error) {
	val := fmt.Sprintf("user:%d", d.Id)
	log.Printf("get %s from database", d.GetKey())
	d.redisClient.Set(d.GetKey(), val, time.Minute)
	return val, nil
}
func (d *userDemo) GetKey() string {
	return fmt.Sprintf("userDemo:%d", d.Id)
}
