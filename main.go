package main

func main() {

	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")

	storage.Load(&todos)

	todos.Print()

	storage.Save(todos)
}
