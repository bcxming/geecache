//并发控制
package geecache
import(
	"gee-cache/LRU"
	"sync"
)

type cache struct{
	mu sync.Mutex
	LRU *LRU.Cache
	cacheBytes int64
}

func (c* cache) add(key string ,value ByteView){
	c.mu.Lock()
	defer c.mu.Unlock()  // 在函数执行完毕后释放锁，以便其他 goroutine 访问

	if c.LRU==nil{
		c.LRU = LRU.New(c.cacheBytes,nil)
	}
	// 调用 LRU 缓存的 Add 方法，向缓存中添加键值对
	c.LRU.Add(key,value)
}

func (c *cache)get(key string)(value ByteView,ok bool){
	c.mu.Lock()
	defer c.mu.Unlock()  // 在函数执行完毕后释放锁，以便其他 goroutine 访问
	if c.LRU==nil{
		return
	}
	if v.ok:=c.LRU.Get(key);ok{
		return v.(ByteView),ok
	}
	return 
}