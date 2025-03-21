package main

import "time"

func main() {
	firstName := "Ali"
	names := []string{"Ali", "Reza", "Peyman", "Pejman", "Hamed1", "Hamed`2", "Hamed3", "Hamed4"}

	printFirstNameFunc := func() {
		println(firstName)
	}

	firstName = "Reza"

	printFirstNameFunc()

	//---------------------------------------

	for i, name := range names {
		go func() {
			name = "*" + name
			println(i, name)
		}()
	}

	time.Sleep(time.Second * 1)

	for i, item := range names {
		go func(index int, name string) {
			name = "*" + name
			println(index, name)
		}(i, item)
	}

	time.Sleep(time.Second * 1)

}
