package main

import (
	"fmt"

	"github.com/fengziren/designPatterns/singleton/dcl"
	"github.com/fengziren/designPatterns/singleton/hungry"
	"github.com/fengziren/designPatterns/singleton/lazy"
	"github.com/fengziren/designPatterns/singleton/lazysafety"
)

func main() {
	// 创建单例对象

	// 懒汉式-线程不安全
	obj1 := lazy.GetInstance()
	obj2 := lazy.GetInstance()
	//singleton not exported by package lazy
	// obj3 := new(lazy.singleton)
	// _ = obj3
	fmt.Printf("obj1.Value: %v\n", obj1.Value)
	fmt.Printf("obj2.Value: %v\n", obj2.Value)
	obj1.Value = 1024
	fmt.Printf("obj2.Value: %v\n", obj2.Value)
	fmt.Println("obj1 == obj2:", obj1 == obj2)
	fmt.Println("----------------------------------------")
	// 懒汉式-线程安全
	objs1 := lazysafety.GetInstance()
	objs2 := lazysafety.GetInstance()
	fmt.Printf("objs1.Value: %v\n", objs1.Value)
	fmt.Printf("objs2.Value: %v\n", objs2.Value)
	objs1.Value = 1024
	fmt.Printf("objs2.Value: %v\n", objs2.Value)
	fmt.Println("objs1 == objs2:", objs1 == objs2)
	fmt.Println("----------------------------------------")
	// 饿汉式
	objh1 := hungry.GetInstance()
	objh2 := hungry.GetInstance()
	fmt.Printf("objh1.Value: %v\n", objh1.Value)
	fmt.Printf("objh2.Value: %v\n", objh2.Value)
	objh1.Value = 1024
	fmt.Printf("objh2.Value: %v\n", objh2.Value)
	fmt.Println("objh1 == objh2:", objh1 == objh2)
	fmt.Println("----------------------------------------")
	// 双检锁/双重校验锁（DCL，即 double-checked locking）
	objd1 := dcl.GetInstance()
	objd2 := dcl.GetInstance()
	fmt.Printf("objd1.Value: %v\n", objd1.Value)
	fmt.Printf("objd2.Value: %v\n", objd2.Value)
	objd1.Value = 1024
	fmt.Printf("objd2.Value: %v\n", objd2.Value)
	fmt.Println("objd1 == objd2:", objd1 == objd2)
	fmt.Println("----------------------------------------")
}
