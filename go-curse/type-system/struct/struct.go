package main

type product struct {
	name     string
	price    float64
	discount float64
}

func (p product) discountPrice() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	// create with named parameters
	p := product{
		name:     "Beer",
		price:    7.50,
		discount: .10,
	}
	println(p.discountPrice())

	// create with syntax sugar

	p = product{
		"Car",
		10.0,
		.1,
	}

	println(p.discountPrice())

}
