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

// 自律工人的工作方式
func (worker Worker) Work(orderIn <-chan Order, productOut chan<- Product) {
	// 工人是独立工作的，当接到订单后就各自干各的，所以要新开一个 Goroutine
	go func() {
		for order := range orderIn {
			productOut <- order.Process()
		}
	}()
}

// 工厂
type Factory struct {
	orderIn    <-chan Order
	productOut chan<- Product
	workers    []Worker
}

// 拥有许多自律工人的先进工厂
func (factory Factory) Work() {
	for _, worker := range factory.workers {
		worker.Work(factory.orderIn, factory.productOut)
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
