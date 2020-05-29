package main

import "fmt"

type item struct {
	productId int
	amount int
	price float64
}

type order struct {
	userId int
	items []item
}

func (o order) totalValue()float64{
	total := 0.0
	for _, item := range o.items{
		total += item.price * float64(item.amount)
	}
	return total
}


func main() {
	order := order{
		1,
		[]item{
			{1, 2, 12.10},
			{2, 2, 14.44},
			{11, 100, 3.12},
		},
	}

	fmt.Printf("Total Order %.2f", order.totalValue())
}
