package singleflight
import(
	"testing"
)

func TestDo(t *testing.T){
	var g Group
	v,err:=g.Do("key",func()(interface{},error){
		// 在本例中，这个函数直接返回结果值 "Bar" 和一个 nil 错误
		return "Bar",nil
	})
	if v!="Bar"||err!=nil{
		t.Errorf("Do v = %v,error = %v",v,err)
	}
}