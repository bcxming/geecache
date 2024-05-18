//缓存值的抽象与封装
package geecache

//抽象了一个只读数据结构 `ByteView` 用来表示缓存值
type ByteView struct {
    b []byte // b 将会存储真实的缓存值
}

// Len 返回字节切片的长度，实现了 Value 接口
func (v ByteView) Len() int {
    return len(v.b) // 返回字节切片的长度
}

// ByteSlice 返回字节切片的副本，防止对原始数据的修改
func (v ByteView) ByteSlice() []byte {
    return cloneBytes(v.b) // 返回字节切片的副本，避免对原始数据的修改
}

// String 返回字节切片的字符串表示形式
func (v ByteView) String() string {
    return string(v.b) // 返回字节切片的字符串表示形式
}

// cloneBytes 克隆字节切片，返回一个新的字节切片，防止对原始数据的修改
func cloneBytes(b []byte) []byte {
    c := make([]byte, len(b)) // 创建一个与原始字节切片相同长度的新字节切片
    copy(c, b) // 将原始字节切片的内容复制到新字节切片中
    return c // 返回新字节切片
}
