package skiplist

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"math/rand"
	"strconv"
	"testing"
)

func TestSkipList_Add(t *testing.T) {
	const N = 30
	sl := NewSkipList()
	for i := 0; i < N; i++ {
		k := rand.Intn(N)
		v := strconv.Itoa(k * 10)
		sl.Add(k, v)
	}

	fmt.Println(sl)
}

func TestSkipList_Get(t *testing.T) {
	const N = 30
	data := make(map[int]string)
	sl := NewSkipList()
	for i := 0; i < N; i++ {
		//k := rand.Intn(N)
		k := rand.Int()
		v := strconv.Itoa(k*10 + i)
		data[k] = v
		sl.Add(k, v)
	}

	fmt.Println(sl)

	for key, value := range data {
		get := sl.Get(key)
		assert.Equal(t, get, value, "fuck")
	}
}

func TestSkipList_Delete(t *testing.T) {
	const N = 20
	data := make(map[int]string)
	sl := NewSkipList()
	//for i := 0; i < N; i++ {
	//	k := rand.Intn(N)
	//	v := strconv.Itoa(k*10 + i)
	//	data[k] = v
	//	sl.Add(k, v)
	//}

	for i := 0; i < N; i++ {
		k := i
		v := strconv.Itoa(k * 10)
		sl.Add(k, v)
		data[k] = v
	}

	fmt.Println(sl)

	for key, value := range data {
		get := sl.Get(key)
		assert.Equal(t, get, value, "fuck")
	}

	//for key, value := range data {
	//	del := sl.Delete(key)
	//	assert.Equal(t, del, value, "fuck")
	//	data[key] = ""
	//}
	//
	//fmt.Println(sl)
	//
	//for key, value := range data {
	//	get := sl.Get(key)
	//	assert.Equal(t, get, value, "fuck")
	//}

	sl.Delete(17)
	//sl.Delete(6)
	//sl.Delete(10)
	fmt.Println(sl)

}

func TestSkipList(t *testing.T) {
	const N = 9000
	data := make(map[int]string)
	del := make(map[int]bool)
	sl := NewSkipList()
	for i := 0; i < N; i++ {
		num := rand.Int()
		k := strconv.Itoa(num)
		sl.Add(num, k)
		data[num] = k
	}

	i := 0
	for k, v := range data {
		if i%3 == 0 {
			v := v + "a"
			sl.Add(k, v)
			data[k] = v
		} else if i%3 == 1 {
			sl.Delete(k)
			del[k] = true
		}
		i++
	}

	//fmt.Println(sl)
	for k, v := range data {
		val := sl.Get(k)
		_, isDel := del[k]

		//fmt.Println(v, val)
		if isDel {
			if val != "" {
				t.Fatalf("got deleted key: %d", k)
			}
		} else {
			if v != val {
				t.Fatalf("get key:%d, expected: %s, got: %s", k, string(v), string(val))
			}
		}
	}
}
