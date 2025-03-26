package main


func main() {

	todos := Todos{}
	todos.add("Learn Go")
	todos.add("Create a Todo App")
	todos.add("Learn React")

	todos.Print()

	todos.Toggle(0)
	todos.Print()

	todos.Delete(1)
	todos.Print()
}