package main

import "fmt"

type bazar struct {
	prilavok string
	kiosk string
	lotok string
}

func (b bazar) vibor() {
	var a int
	fmt.Print("что хотите чекнуть? ")
	fmt.Scan(&a)
	if a== 1{
		fmt.Println(b.prilavok)
	}
	if a == 2{
		fmt.Println(b.kiosk)
	}
	if a == 3{
		fmt.Println(b.lotok)
	}
} 
type cost struct{
	k int
}
func (c *cost) cena(){
	var a string
	var b int
	fmt.Print("ну что, поторгуемся? ")
	fmt.Scan(&a)
	if a =="да"{
		fmt.Print("предложи цену ")
		fmt.Scan(&b)
		c.k = b
	}
}
func main() {
	bazar:= bazar{"помидоры","огурцы","детская присыпка"}
	bazar.vibor()
	cost:= cost{100}
	cost.cena()
	fmt.Println(cost.k)
}
