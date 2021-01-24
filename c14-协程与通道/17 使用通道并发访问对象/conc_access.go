package main

import (
	"fmt"
	"strconv"
)

/**
为了保护对象被并发访问修改，我们可以使用协程在后台顺序
执行匿名函数来替代使用同步互斥锁。
*/

type Person struct {
	Name   string
	salary float64
	chF    chan func()
}

/**
1.初始化时会启动一个后台协程backend()。
2.
*/
func NewPersons(name string, salary float64) *Person {
	p := &Person{name, salary, make(chan func())}
	go p.backend()
	return p
}

/**
在一个无限循环中执行chF中放置的所有函数，有效的将它们序列化
从而提供了安全的并发访问。
*/
func (p *Person) backend() {
	for f := range p.chF {
		f()
	}
}

// Set salary
/**
更改和读取salary的方法会通过将一个
匿名函数写入chF通道中，然后让backend()按顺序执行以达到其目的。
*/
func (p *Person) SetSalary(sal float64) {
	p.chF <- func() {
		p.salary = sal
	}
}

// Retrieve salary
/**
序列化打印工资
*/
func (p *Person) Salary() float64 {
	fChan := make(chan float64)
	p.chF <- func() {
		fChan <- p.salary
	}
	return <-fChan
}

func (p *Person) String() string {
	return "Person - name is: " + p.Name + " - Salary is: " + strconv.FormatFloat(p.Salary(), 'f', 2, 64)
}

func main() {
	bs := NewPersons("Smith Bill", 2500.5)
	fmt.Println(bs)
	bs.SetSalary(4000.25)
	fmt.Println("Salary changed:")
	fmt.Println(bs)
}
