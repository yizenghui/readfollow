// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

func init() {}

type Person struct {
	Name string
	Age  int
	Date time.Time
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d %s", p.Name, p.Age, p.Date)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age > a[j].Age }

func Test_NewBookSort(t *testing.T) {

	// toBeCharge := "2015-01-01 00:00:00"              //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	// timeLayout := "2006-01-02 15:04:05"              //转化所需模板
	// theTime, _ := time.Parse(timeLayout, toBeCharge) //使用模板在对应时区转化为time.time类型

	t1, _ := time.Parse("2006-01-02 15:04:05", "2015-01-01 00:00:00")
	t2, _ := time.Parse("2006-01-02 15:04:05", "2015-01-02 00:00:00")
	t3, _ := time.Parse("2006-01-02 15:04:05", "2015-01-03 00:00:00")
	t4, _ := time.Parse("2006-01-02 15:04:05", "2015-01-04 00:00:00")

	// fmt.Println(theTime)
	pm := Person{"Michael", 17, t1}

	people := []Person{
		{"Bob", 31, t2},
		{"John", 42, t3},
		pm,
		{"Jenny", 26, t4},
	}

	fmt.Println(pm)
	fmt.Println(people)
	sort.Sort(sort.Reverse(ByAge(people))) // 反序
	sort.Sort(ByAge(people))
	// sort.Reverse(ByAge(people))
	// delete(people, 1) 这个只能删除键
	fmt.Println(people[0:3])
}

func Test_MakeSlice(t *testing.T) {

	var ss []string
	fmt.Printf("[ local print ]\t:\t length:%v\taddr:%p\tisnil:%v\n", len(ss), ss, ss == nil)
	print("func print", ss)
	//切片尾部追加元素append elemnt
	for i := 0; i < 12; i++ {
		ss = append(ss, fmt.Sprintf("s%d", i))
	}
	fmt.Printf("[ local print ]\t:\tlength:%v\taddr:%p\tisnil:%v\n", len(ss), ss, ss == nil)
	print("after append", ss)
	//删除切片元素remove element at index
	index := 5
	ss = append(ss[:index], ss[index+1:]...)
	print("after delete", ss)
	//在切片中间插入元素insert element at index;
	//注意：保存后部剩余元素，必须新建一个临时切片
	rear := append([]string{}, ss[index:]...)
	ss = append(ss[0:index], "inserted")
	ss = append(ss, rear...)
	print("after insert", ss)

	max := 10
	count := len(ss)
	if count > max {
		start := count - max
		ss = append([]string{}, ss[start:]...)

		print("max limit", ss[start:])
	}
	print("max limit", ss)

}
func print(msg string, ss []string) {
	fmt.Printf("[ %20s ]\t:\tlength:%v\taddr:%p\tisnil:%v\tcontent:%v", msg, len(ss), ss, ss == nil, ss)
	fmt.Println()
}
