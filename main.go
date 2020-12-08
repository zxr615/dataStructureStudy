package main

import (
	"fmt"
	"study/single"
	"study/sqlist"
)

func main() {
	//sqlistTest()
	singleTest()
}

func singleTest() {
	list := single.New()
	//list.Add("t1")
	list.Append("A")
	list.Append("B")
	list.Append("C")
	list.Append("D")
	//list.Insert(3, "C")
	list.Remove(2)

	list.Show()
}

func sqlistTest() {
	l := sqlist.New(5)

	_ = (*l).Insert(1, 11)
	_ = l.Insert(2, 22)
	_ = l.Insert(3, 33)
	_ = l.Insert(4, 44)
	_ = l.Insert(5, 55)

	v, e := l.Get(7)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println("get:", v)
	}

	if err := l.Del(5); err == nil {
		fmt.Println("del succ")
	} else {
		fmt.Println("del err:", err.Error())
	}

	if err := l.Add(98); err == nil {
		fmt.Println("add succ")
	} else {
		fmt.Println("del err:", err.Error())
	}

	if err := l.Add(99); err == nil {
		fmt.Println("add succ")
	} else {
		fmt.Println("del err:", err.Error())
	}
}
