// 1. PACKAGE
package main

// 2. IMPORT
import (
	"fmt"
	//"learning-go/mathutil"
)

//"math"

/*func main() {
	fmt.Println(mathutil.Add(2, 5))
	fmt.Println(mathutil.Subtract(1, 2))
	fmt.Println(math.Sqrt(144))
}*/

// 3. VAR
/*func main() {
	var a, b = 5, 3
	fmt.Println(a, b)

	y := 5
	fmt.Println(y)

	nombre, age := "Juan", 32

	fmt.Println(nombre) // "Juan"
	fmt.Println(age)   // 32

	fmt.Println(mathutil.Age)
}

// 4. CONST
/*const Pi = 3.14

const (
	A = 1
	B = 4
)

func main() {
	fmt.Println(Pi)
	fmt.Println(A + B)
}*/

// 5. FUNC
/*func sayHello() {
	fmt.Println("Hellooooo!")
}

// 6. IF

func main() {
	sayHello()

	sumar := add(4, 4)
	fmt.Println(sumar)

	divide, err := divide(144, 12)
	if err != nil {
		fmt.Println("There was an error.")
	}
	fmt.Println(divide)
}

func add(x, y int) int {
	return x + y
}

func divide(x, y int) (int, error) {
	if x <= 0 {
		return 0, errors.New("cannot be zero")
	}
	return x / y, nil
}*/

// 7. ELSE
/*func main() {
	x := -1
	if x > 0 {
		fmt.Println("is positive")
	} else if x == 0 {
		fmt.Println(("is zero"))
	} else {
		fmt.Println(("is negative"))
	}
}*/

// 8. FOR
/*func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	num := []int{1, 2, 3}
	for _, num := range num {
		fmt.Println(num)
	}
}*/

// 9. RETURN
/*func sayGoodbye() string {
	return "Goodbye"
}

func divide(a, b int) int {
	if b == 0 {
		return 0
	}

	return a / b
}

func main() {
	goodbye := sayGoodbye()
	fmt.Println(goodbye)

	divided := divide(1, 2)
	fmt.Println(divided)
}*/

// 10. SWITCH - CASE - DEFAULT
/*
func main() {
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Other day")
	}

	letter := "a"
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("Vocal")
	default:
		fmt.Println("Consonante")
	}

	number := 0
	switch number {
	case 0, 1:
		fmt.Println("Number is either 0 or 1")
	default:
		fmt.Println("Number is something else")
	}

	x := 3
	switch x {
	case 1:
		fmt.Println("One")
	default:
		fmt.Println("Not One")
	}

	fruit := "pear"
	switch fruit {
	case "apple":
		fmt.Println("apple")
	default:
		fmt.Println("It is not a fruit")
	}
}
*/

// 13. BREAK
/*func main() {
	for i := 0; i < 5; i++ {
		if i == 0 {
			break
		}
		fmt.Println(i) // output: 0, 1, 2
	}

	x := 2
	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Second")
	case 3:
		fmt.Println("Three")
	}
	fmt.Println("Switch complete!")
}*/

// 14. CONTINUE
/*func main() {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i) // Output: 1, 3
	}

	for i := 1; i <= 3; i++ {
		for j := 1; i <= 3; i++ {
			if j == 2 {
				continue
			}
			fmt.Println(i, j)
		}
	}
}*/

// 15. RANGE
/*func main() {
	nums := []int{1, 2, 3, 4}
	for i, num := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}

	m := map[string]int{"a": 1, "b": 2}
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}*/

// 16. FALLTHROUGH
/*func main() {
	x := 1
	switch x {
	case 1:
		fmt.Println("Case 1")
		fallthrough
	case 2:
		fmt.Println("case 2")
	}

	y := 2
	switch y {
	case 1:
		fmt.Println("Start Case 1")
		fallthrough
	case 2:
		fmt.Println("Ending case 1 and Case 2")
	case 3:
		fmt.Println("Case 3")
	}
}*/

// 17. DEFER
/*func main() {
	defer fmt.Println("Esto esta impreso ultimo")
	fmt.Println("Esto esta impreso primero")

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Loop finished")
}*/

// 18. TYPE
/*type MyInt int
type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
	return float64(c * 9 / 5 * 32)
}

func main() {
	var m = 5
	fmt.Println(m)

	temp := Celsius(25)
	fmt.Println(temp.ToFahrenheit())
}*/

// 18. STRUCT

/*type Person struct {
	Name string
	Age  int
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	p := Person{
		Name: "Emiliano",
		Age:  36,
	}
	fmt.Println(p)

	rect := Rectangle{
		Width:  5,
		Height: 5,
	}

	fmt.Println(rect.Area()) // Output: 25

}*/

// 19. GO

/*func Saluda() {
	fmt.Println("Hola desde un goroutine")
}

func printNumbers(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go Saluda()
	time.Sleep(time.Second)

	fmt.Println("main function completed")

	go printNumbers(5)
	go printNumbers(5)
	time.Sleep(time.Second) // Allow time for both goroutine
}*/

// 20. CHAN
/*func main() {
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	val := <-ch
	fmt.Println(val) // Output: 42

	ch2 := make(chan string, 2)
	ch2 <- "Hello"
	ch2 <- "World"
	fmt.Println(<-ch2) // Output: Hello
	fmt.Println(<-ch2) // Output: World
}*/

// 20. SELECT
/*func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	ch := make(chan string)

	go func() { ch1 <- "Channel 1" }()
	go func() { ch2 <- "Channel 2" }()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Done"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout!")
	}
}*/

// 21. MAP
/*func main() {
	m := map[string]int{"manzana": 2, "banana": 5}
	fmt.Println(m["manzana"]) // Output: 2

	for key, value := range m {
		fmt.Printf("key es: %s; value es: %d", key, value)
	}
}*/

// 22. INTERFACE

/*type Hablador interface {
	Hablar() string
}

type Perro struct{}

func (p Perro) Hablar() string {
	return "Wuuf!"
}

func imprimirCualquierCosa(v interface{}) {
	fmt.Println(v)
}

func main() {
	var p Hablador = Perro{}
	fmt.Println(p.Hablar())

	imprimirCualquierCosa(42)
	imprimirCualquierCosa("Hola")
	imprimirCualquierCosa(423)
}*/

// 23. GOTO
func main() {
	fmt.Println("Antes de goto")
	goto End
	fmt.Println("Esto no se ejecutara")
End:
	fmt.Println("Despues de goto")

	i := 0
Loop:
	if i < 3 {
		fmt.Println(i)
		i++
		goto Loop
	}

}
