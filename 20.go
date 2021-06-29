package main

func main() {}func main() {
	slice := []string{"a", "a"}
  
	func(slice []string) {
	   slice = append(slice, "a") // slice уже указывает на другую область памяти
	   slice[0] = "b"
	   slice[1] = "b"
	   fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
  }
  