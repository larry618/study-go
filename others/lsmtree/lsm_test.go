package lsmtree

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func TestLSM(t *testing.T) {
	conf := LSMConf{}
	conf.BlockSize = 1024 * 16
	conf.Dir = "/Users/Test/Temp/data/lsm/"
	os.RemoveAll(conf.Dir)
	os.MkdirAll(conf.Dir, os.ModeDir|os.ModePerm)
	conf.MemSyncSize = 1024 * 1024
	lsm := NewLSM(&conf)
	const N = 10000
	data := make(map[int][]byte)
	del := make(map[int]bool)
	for i := 0; i < N; i++ {
		num := rand.Int()
		k := []byte(strconv.Itoa(num))
		lsm.Add(k, k)
		data[num] = k
	}
	lsm.SyncAll()

	fmt.Println("hehe levels ", lsm.Levels())

	lsm = NewLSM(&conf)
	i := 0
	for k, v := range data {
		key := []byte(strconv.Itoa(k))
		if i%3 == 0 {
			v := append([]byte("a"), v...)
			lsm.Add(key, v)
			data[k] = v
		} else if i%3 == 1 {
			lsm.Rem(key)
			del[k] = true
		}
		i++
	}
	lsm.SyncAll()

	lsm = NewLSM(&conf)
	for k, v := range data {
		val, ok := lsm.Get([]byte(strconv.Itoa(k)))
		_, isDel := del[k]
		if isDel {
			if ok {
				t.Fatalf("got deleted key: %d", k)
			}
		} else {
			if !ok || bytes.Compare(v, val) != 0 {
				t.Fatalf("get key:%d, expected: %s, got: %s", k, string(v), string(val))
			}
		}
	}
}
