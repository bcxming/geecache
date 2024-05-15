package singleflight
import "sync"

type call struct {
	wg sync.WaitGroup //用于等待异步操作的完成
	val interface{}  //用于存储异步操作的结果值
	err error  //用于存储异步操作的错误信息
}

type Group struct{
	mu sync.Mutex
	m map[string]*call //存储异步调用的结果
}

func (g *Group)Do(key string,fn func()(interface{},error))(interface{},error){
	//fn 这个函数没有参数，并返回两个值，一个是 interface{} 类型的结果值，另一个是 error 类型的错误信息
	g.mu.Lock()
	// 检查 g.m 是否为 nil，如果是，则初始化为一个新的 map
	if g.m==nil{
		g.m=make(map[string]*call)
	}
	//// 检查 g.m 中是否已存在指定 key 对应的异步调用结果
	if c,ok:=g.m[key];ok{
		g.mu.Unlock()
		c.wg.Wait()
		return c.val,c.err
	}
	c:=new(call)
	c.wg.Add(1) // 将异步操作的等待组设置为 l，表示异步操作执行完成后会调用 Done 方法
	g.m[key] = c
	g.mu.Unlock()

	// 执行异步操作并获取结果值和可能的错误信息
	c.val,c.err = fn()
	c.wg.Done()

	g.mu.Lock()
	// 异步操作执行完成后，从 g.m 中删除相应的条目，释放资源
	delete(g.m,key)
	g.mu.Unlock()
	
	return c.val,c.err
}