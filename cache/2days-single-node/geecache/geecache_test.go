package geecache

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

// 测试单个方法
// go test -v -test.run TestGetGroup

var db = map[string]string{
	"Tom":  "630",
	"jack": "589",
	"sam":  "657",
}

func TestGetter(t *testing.T) {
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})
	expect := []byte("key")

	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Fatal("callback failed")
	}
}

func TestGetGroup(t *testing.T) {

	groupNmae := "scores"

	NewGroup(groupNmae, 2<<10, GetterFunc(
		func(key string) (bytes []byte, err error) { return }))

	if group := GetGroup(groupNmae); group == nil || group.name != groupNmae {
	}
	t.Fatalf("group %s not exist", groupNmae)

	if group := GetGroup(groupNmae + "11111"); group != nil {
		t.Fatalf("expecte nikk ,but %s got", group.name)
	}
}

func TestGet(t *testing.T) {
	loadCounts := make(map[string]int, len(db))

	gee := NewGroup("scores", 2<<10, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[Slow DB]，search key ", key)
			if v, ok := db[key]; ok {
				if _, ok := loadCounts[key]; !ok {
					loadCounts[key] = 0
				}

				loadCounts[key]++
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s,not exist", key)
		}))

	for k, v := range db {
		if view, err := gee.Get(k); err != nil || view.String() != v {
			t.Fatal("failed to get valued of Tom")
		}

		if _, err := gee.Get(k); err != nil || loadCounts[k] > 1 {
			t.Fatalf("cache %s miss", k)
		}
	}

	if view, err := gee.Get("unknown"); err == nil {
		t.Fatalf("the value of unknow should be empty,but %s got", view)
	}
}
