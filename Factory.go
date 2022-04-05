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
type Worker struct {
	orderIn    chan Order
	productOut chan Product
	stop       <-chan struct{}
}

func NewWorker(stop <-chan struct{}) Worker {
	orderIn := make(chan Order)
	productOut := make(chan Product)
	go func() {
		for {
			select {
			case <-stop:
				return
			case order := <-orderIn:
				productOut <- order.Process()
			}
		}
	}()
	return Worker{orderIn, productOut, stop}
}

// 工厂
type Factory struct {
	orderIn    <-chan Order
	productOut chan<- Product
	stop       chan struct{}
	workers    []Worker
}

func (factory Factory) Work() {
	for _, worker := range factory.workers {
		// 将订单分配给工人
		go func(orderIn chan<- Order) {
			for order := range factory.orderIn {
				orderIn <- order
			}
		}(worker.orderIn)

		// 收集工人生产的产品
		go func(productOut <-chan Product) {
			for product := range productOut {
				factory.productOut <- product
			}
		}(worker.productOut)
	}
}

func (factory Factory) Stop() {
	close(factory.stop)
}

func main() {
	// 初始化订单输入与产品输出的 Channel
	orderIn := make(chan Order, 10)
	productOut := make(chan Product, 10)
	stop := make(chan struct{})
	factory := Factory{
		orderIn,
		productOut,
		stop,
		[]Worker{NewWorker(stop), NewWorker(stop), NewWorker(stop)},
	}
	factory.Work()
	for {
		select {}
	}
}
