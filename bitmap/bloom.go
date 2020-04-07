package main

/**
 * Created by @CaomaoBoy on 2020/2/25.
 *  email:<115882934@qq.com>
 */

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/spaolacci/murmur3" //hash计算
	"github.com/willf/bitset"      //bitset
	"io"
	"math"
)
type BloomFilter struct {
	m uint    //总量
	k uint   //k
	b*bitset.BitSet //存储value
}

//判断
func max(x,y uint)uint{
	if x>y{
		return x
	}
	return y
}
//新建一个布隆过滤器
func  NewBloomFilter(m uint,k uint)*BloomFilter{
	return &BloomFilter{max(1,m),max(1,k),bitset.New(m)}
}
//根据数据规模新建布隆过滤器
func From(data []uint64,k uint)*BloomFilter{
	m:=uint(len(data)*64)
	return &BloomFilter{m,k,bitset.From(data)}
}
//基础哈希计算
func  baseHashes(data[]byte)[4]uint64{
	X:=[]byte{1}
	hasher:=murmur3.New128()//128 ，2^128
	hasher.Write(data)
	v1,v2:=hasher.Sum128() //返回两个整数
	hasher.Write(X)
	v3,v4:=hasher.Sum128()//返回两个整数
	return [4]uint64{v1,v2,v3,v4}
}
//求出本地哈希
func  location(h [4]uint64,i uint)uint64{
	ii:=uint64(i)
	//return h[ii%2]+ii*h[2+(((ii+(ii%2))%4)/2)]
	return h[ii%2] + ii*h[2+(((ii+(ii%2))%4)/2)]
}
//本地计算
func (f*BloomFilter) location(h [4]uint64,i uint)uint{
	return uint(location(h,i)%uint64(f.m))
}
//数据概率的预估
func EStamatewithParameters(n uint,p float64)(m uint,k uint){
	m=uint(math.Ceil(-1*float64(n)*math.Log(p))/math.Pow(math.Log(2),2))
	k=uint(math.Ceil(math.Log(2)*float64(m)/float64(n)))
	return m,k
}
//新建一个布隆过滤器，预估一下数据规模
func NewwithEstimates(n uint,p float64)*BloomFilter{
	m,k:=EStamatewithParameters(n,p)
	return NewBloomFilter(m,k)
}

func (f*BloomFilter)K()uint{
	return f.k //哈希
}
func (f*BloomFilter)Cap()uint{
	return f.m //总量
}

func (f*BloomFilter)Add(data []byte)*BloomFilter{
	h:=baseHashes(data)//计算哈希
	for i:=uint(0);i<f.k;i++{
		f.b.Set(f.location(h,i)) //设置数据
	}
	return f
}
func (f*BloomFilter) Merge(g*BloomFilter)error{
	if f.m!=g.m{
		return fmt.Errorf("大小不一样")
	}
	if f.k!=g.k{
		return fmt.Errorf("大key不一样")
	}
	f.b.InPlaceUnion(g.b)//归并bitset解决空间
	return nil
}
//拷贝新建一个布隆过滤器
func (f*BloomFilter)Copy()*BloomFilter{
	fc:=NewBloomFilter(f.m,f.k)
	fc.Merge(f)//归并
	return fc
}
//增加字符串
func (f*BloomFilter)AddString(data string)*BloomFilter{
	return f.Add([]byte(data))
}
//判断数据存在与否
func (f*BloomFilter)Test(data[]byte)bool{
	h:=baseHashes(data)
	for i:=uint(0);i<f.k;i++{
		if !f.b.Test(f.location(h,i)){
			return false
		}
	}
	return true
}
//判断字符串是否存在
func (f*BloomFilter)TestString(data string)bool{
	return f.Test([]byte(data))
}
//测试整数是否存在
func (f*BloomFilter)TestLocations(locs []uint64)bool{
	for i:=0;i<len(locs);i++{
		if !f.b.Test(uint(locs[i]%uint64(f.m))){
			return false
		}
	}
	return true
}
//测试是否存在并添加。存在就更新，不存在插入
func (f*BloomFilter)TestAndAdd(data[]byte)bool{
	isin:=true
	h:=baseHashes(data)
	for i:=uint(0);i<f.k;i++{
		l:=f.location(h,i)
		if !f.b.Test(f.location(h,i)){
			isin=false
		}
		f.b.Set(l)
	}
	return isin
}
//测试字符串是否存在，
func (f*BloomFilter)TestAndAddString(data string)bool{
	return f.TestAndAdd([]byte(data))
}
//清空布隆过滤恶气
func (f*BloomFilter)Clear()*BloomFilter{
	f.b.ClearAll()
	return f
}
//测试正确率
func (f*BloomFilter)EstimateFalsePositiveRate(n uint,rounds uint32)(fpRate float64){

	f.Clear()
	n1:=make([]byte,4)//开辟字节数组
	for i:=uint32(0);i<uint32(n);i++{
		binary.BigEndian.PutUint32(n1,i)
		f.Add(n1)//测试追加
	}
	fp:=0
	for i:=uint32(0);i<rounds;i++{
		binary.BigEndian.PutUint32(n1,i+uint32(n)+1)
		if f.Test(n1) {
			fp++
		}

	}
	fpRate=float64(fp)/float64(rounds)//速率
	f.Clear()
	return


}
//布隆过滤器的
type BloomFilterJson struct{
	M uint  `json:"m"`
	K uint  `json:"k"`
	B* bitset.BitSet  `json:"b"`
}
//映射
func (f*BloomFilter)MarshaJson()([]byte,error){
	return json.Marshal(BloomFilterJson{f.m,f.k,f.b})
}
//字节转化对象
func (f*BloomFilter)UnMarshaJson(data []byte)(error){
	var  j   BloomFilterJson
	err:=json.Unmarshal(data,&j)
	if err!=nil{
		return err
	}
	f.m=j.M
	f.k=j.K
	f.b=j.B
	return nil
}
func  (f*BloomFilter)Writeto(stream io.Writer)(int64,error){
	err:=binary.Write(stream,binary.BigEndian,uint64(f.m))
	if err!=nil{
		return 0,err
	}
	err=binary.Write(stream,binary.BigEndian,uint64(f.k))
	if err!=nil{
		return 0,err
	}
	numbytes,err:=f.b.WriteTo(stream)
	return numbytes+int64(2*binary.Size(uint64(0))),err


}
func  (f*BloomFilter)Readfrom(stream io.Reader)(int64,error){
	var m,k uint64
	err:=binary.Read(stream,binary.BigEndian,&m)
	if err!=nil{
		return 0,err
	}
	err=binary.Read(stream,binary.BigEndian,&k)
	if err!=nil{
		return 0,err
	}
	b:=&bitset.BitSet{}
	numbyte,err:=b.ReadFrom(stream)
	if err!=nil{
		return 0,err
	}
	f.m=uint(m)
	f.k=uint(k)
	f.b=b
	return numbyte+int64(2*binary.Size(uint64(0))),err

}
func  (f*BloomFilter)GobEncode()([]byte,error){
	var buf bytes.Buffer
	_,err:=f.Writeto(&buf)
	if err!=nil{
		return nil,err
	}
	return buf.Bytes(),nil

}
func  (f*BloomFilter)GobDecode(data []byte)(error){
	buf:=bytes.NewBuffer(data)
	_,err:=f.Readfrom(buf)
	return err
}
//判断布隆过滤恶气是否相等
func  (f*BloomFilter)Equal(g*BloomFilter)bool{
	return f.m==g.m&&f.k==g.k&&f.b.Equal(g.b)
}
//本地计算哈希
func locations(data []byte, k uint)[]uint64{
	locs:=make([]uint64,k)
	h:=baseHashes(data)
	for i:=uint(0);i<k;i++{
		locs[i]=location(h,i)
	}

	return locs
}



func main(){
	f:=NewwithEstimates(1000000000,0.0001)
	n1:=[]byte("yincheng")
	n2:=[]byte("QQ77025077")
	n3:=[]byte("mobile18510341407")
	n4:=[]byte("afsdasddsdsds")
	n5:=[]byte("afsdasddsdsdsx")
	f.Add(n1)
	f.Add(n2)
	f.Add(n3)
	data,_ := f.MarshaJson()
	var g BloomFilter
	err := g.UnMarshaJson(data)
	if err != nil{
		fmt.Println("解码失败!")
	}
	fmt.Println(f.Test(n1))
	fmt.Println(f.Test(n3))
	fmt.Println(f.Test(n4))
	fmt.Println(f.Test(n5))
	fmt.Println(f.Test(n3))
	f.Clear()
	//fmt.Println(f.TestLocations([]uint64{1,2,3,4,5,6}))
	fmt.Println(f.EstimateFalsePositiveRate(100000000,100000))
}