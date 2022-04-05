package main

import (
	"math/rand"
	"time"
)

// 产品
type Product struct{}

// 订单
type Order struct{}

func (order Order) Process() Product {
	// 假装处理订单
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return Product{}
}

// 工人
type Worker struct{}

func (worker Worker) Work(order Order) Product {
	return order.Process()
}

// 工厂
type Factory struct {
	orderIn    <-chan Order
	productOut chan<- Product
	workers    []Worker
}

func (factory Factory) Work() {
	for order := range factory.orderIn {
		factory.productOut <- factory.workers[0].Work(order)
	}
}

func main() {
	// 初始化订单输入与产品输出的 Channel
	orderIn := make(chan Order, 10)
	productOut := make(chan Product, 10)
	factory := Factory{
		orderIn,
		productOut,
		[]Worker{Worker{}, Worker{}, Worker{}},
	}
	factory.Work()
	for {
		select {}
	}
}
