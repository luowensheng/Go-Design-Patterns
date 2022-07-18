package facade

import "fmt"


func MakeAmericano(size float32){
	fmt.Println("making americano")

	americano := &CoffeeMachine{}

	beansAmount := (size/8.0)*5.0
	americano.startCoffee(beansAmount, 2)
    americano.grindBeans()
	americano.useHotWater(size)
	americano.endCoffee()

	fmt.Println("americano is ready")

}

func MakeLatte(size float32){
	fmt.Println("making Latte")

	latte := &CoffeeMachine{}

	beansAmount := (size/8.0)*5.0
	latte.startCoffee(beansAmount, 4)
    latte.grindBeans()
	latte.useHotWater(size)

	milk := (size/8.0)*2.0
	latte.useMilk(milk)
	latte.doFoam(true)
    latte.endCoffee()
	fmt.Println("Latte is ready")

}