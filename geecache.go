// 负责与外部交互，控制缓存存储和获取的主流程
package geecache
import(
	"fmt"
	"log"
	"sync"
)

type Group strcut{
	name string
	getter getter //缓存未命中时获取源数据的回调
	mainCache cache //一开始实现的并发缓存
}

//定义接口 Getter 和 回调函数 `Get(key string)([]byte, error)`，参数是 key，返回值是 []byte
type Getter interface{
	Get(key string)([]byte,error)
}

type GetterFunc func(key string)([]byte,error)

//定义函数类型 GetterFunc，并实现 Getter 接口的 `Get` 方法
func (f GetterFunc)Get(key string)([]byte,error){
	return f(key)
}

var(
	mu sync.RWMutex
	groups = make(map[stirng]*Group)
)

func NewGroup(name string,cacheBytes int64,getter Getter)*Group{
	if getter ==nil{
		panic("nil Getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := & Group{
		name:name,
		getter:getter,
		mainCache:cache{cacheBytes:cacheBytes}
	}
	groups[name] = g
	return g
}

func GetGroup(name string)*Group{
	mu.RLock()
	g:=groups[name]
	mu.RUnlock()
	return g
}

func (g* Group)Get(Key string)(ByteView ,error){
	if key==""{
		return ByteView{},fmt.Errorf("key is required")
	}
	if v,ok:=g.mainCache.get(key);ok{
		log.Println("[GeeCache] hit")
		return v,nil
	}
	return g.load(key)
}

func (g *Group)load(key string)(value ByteView,err error){
	return g.getlocaly(key)
}

func (g *Group)getlocaly(key string)(ByteView ,error){
	bytes,err := g.getter.Get(key)
	if err!=nil{
		return ByteView{},err
	}
	value := ByteView{b:cloneBytes(bytes)}
	g.populateCache(key,value)
	return value,nil
}

func (g *Group)populateCache(key stirng,value ByteView){
	g.mainCache.add(key,value)
}
