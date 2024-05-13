package LRU
import(
	"reflect"
	"testing"
	"./LRU"
)

type String string

func (d String) len() int{
	return len(d)
}

func TestGet(t *testing.T){
	LRU:=New(int64(0),nil)
	LRU.Add("Key1",String("1234"))
	if v,ok:=LRU.Get("Key1");!ok||string(v.(String))!="1234"{
		t.Fatalf("缓存未命中")
	}
	if _,ok:=LRU.Get("Key2");ok{
		t.Fatalf("缓存丢失")
	}
}

func TestRemoveOld(t *testing.T){
	k1, k2, k3 := "key1", "key2", "k3"
	v1, v2, v3 := "value1", "value2", "v3"
	cap:=len(k1+k2+v1+v2)
	LRU:=New(int64(cap),nil)
	LRU.Add(k1,String(v1))
	LRU.Add(k2,String(v2))
	LRU.Add(k3,String(v3))
	if _,ok:=LRU.Get("key1");ok||LRU.Len()!=2{
		t.Fatalf("移除失败")
	}
}


func TestOnEvicted(t *testing.T) {
	keys := make([]string, 0)
	callback := func(key string, value Value) {
		keys = append(keys, key)
	}
	LRU := New(int64(10), callback)
	LRU.Add("key1", String("123456"))
	LRU.Add("k2", String("k2"))
	LRU.Add("k3", String("k3"))
	LRU.Add("k4", String("k4"))

	expect := []string{"key1", "k2"}

	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("回调失败")
	}
}

func TestAdd(t *testing.T) {
	LRU := New(int64(0), nil)
	LRU.Add("key", String("1"))
	LRU.Add("key", String("111"))

	if LRU.nbytes != int64(len("key")+len("111")) {
		t.Fatal("添加失败")
	}
}
