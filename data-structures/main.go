package main

import "fmt"

// ARRAYS
/*func main() {
	var arr [5]int // [0,0,0,0,0]
	arr[0] = 10    // [10,0,0,0,0]
	arr[0] = 20    // [20,0,0,0,0]

	fmt.Println(arr)    // Output: [10, 20,0,0,0]
	fmt.Println(arr[1]) // Output: 20

	arr2 := [3]string{"Go", "Rust", "Python"}
	for i, v := range arr2 {
		fmt.Printf("Index %d, %s\n", i, v)
	}
}*/

// SLICES
/*func main() {
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5)
	fmt.Println(slice) // Output: [1,2,3,4,5]

	arr := [5]int{10, 20, 30, 40, 50}
	slice2 := arr[1:4]
	fmt.Println(slice2) // Output: [20,30,40]
}*/

// MAPS

/*type Person struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]int)
	m["apple"] = 2
	m["banana"] = 5
	fmt.Println(m["apple"]) // Output: 2

	p := Person{Name: "Alice", Age: 36}
	fmt.Println(p)
}*/

// CHANNELS
/*func main() {
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	fmt.Println(<-ch)
}*/

// POINTERS
/*func main() {
	x := 10
	p := &x
	fmt.Println(p)  // Output: Direccion de memoria de x(ej: 0x14000010100)
	fmt.Println(*p) // Output: 10
}*/
/*func main() {
	x := 20
	p := &x
	*p = 50
	fmt.Println(x) // Output: 50
}*/

// Sin pointer - no modifica el original
/*func incrementarPorValor(x int) {
	x = x + 1
	fmt.Println("Dentro de la function:", x) // Output: 11
}

func incrementarPorPointer(x *int) {
	*x = *x + 1
	fmt.Println("Dentro de la function:", *x) // Output: 22
}

func main() {
	num := 10
	incrementarPorValor(num)
	fmt.Println("Despues de incrementarPorValor:", num)

	incrementarPorPointer(&num)
	fmt.Println("Desde de incrementarPorPointer:", num)
}*/

/*type LargeStruc struct {
	Data [1000]int
	Name string
	ID   int
}

// Sin pointer - copia toda la estructura (ineficiente)
func processByValue(ls LargeStruc) {
	fmt.Println("Procesando:", ls.Name)
	// Aqui se copiaron 1000 enteros + string + int
}

// Con pointer - solo se pasa la direccion (eficiente)
func processByPointer(ls *LargeStruc) {
	fmt.Println("Procesando:", ls.Name)
	// Solo se paso una direccion de memoria (8 bytes en sistema)
}

func main() {
	large := LargeStruc{
		Name: "Mi estructura grande",
		ID:   2,
	}

	processByValue(large)    // Copia toda la estructura
	processByPointer(&large) // Solo pasa la direccion
}*/

type CuentaDeBanco struct {
	Dueño   string
	Balance float64
}

// Metodo con receptor por valor - NO modifica el original
func (ba CuentaDeBanco) DepositoPorValor(amount float64) {
	ba.Balance += amount
	fmt.Printf("Saldo despues del deposito (metodo por valor): %.2f\n", ba.Balance)
}

// Metodo con receptor por pointer - SI modifica el original
func (ba *CuentaDeBanco) Deposito(cantidad float64) {
	ba.Balance += cantidad
	fmt.Printf("Saldo despues del deposito: %.2f\n", ba.Balance)
}

func (ba *CuentaDeBanco) Retirar(cantidad float64) {
	if ba.Balance >= cantidad {
		ba.Balance -= cantidad
		fmt.Printf("Retiro exitoso. Saldo actual: %.2f\n", ba.Balance)
	} else {
		fmt.Println("Fondos insuficientes")
	}
}

func main() {
	account := CuentaDeBanco{
		Dueño:   "John Doe",
		Balance: 1000.0,
	}

	fmt.Printf("Saldo inicial: %.2f\n", account.Balance) // 1000

	// Metodo por valor - no modifica el original
	account.DepositoPorValor(100.0)
	fmt.Printf("Saldo real despues de DepositByValue: %.2f\n", account.Balance) // Sigue siendo 1000

	// Metodo por pointer - modifican el original
	account.Deposito(200.0)
	fmt.Printf("Saldo real despues de DepositByValue: %.2f\n", account.Balance) // Sigue siendo 1200

	account.Retirar(150.0)
	fmt.Printf("Saldo final: %.2f\n", account.Balance) // Ahora es 1050
}
