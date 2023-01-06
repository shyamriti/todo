package main

import "fmt"

func main() {

	var (
		todoList []string
		n        int
		u, t     string
	)
	fmt.Printf("Todo LIST\n")

	fmt.Println(" Todolist was empty. So please firstly you need to add some task.")
	fmt.Println("Enter number of tasks")
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%s \n", &t)
		todoList = append(todoList, t)
	}

	for i := range todoList {
		fmt.Printf("%v\t", todoList[i])
	}
	fmt.Println("\n\nFor more add type 'add' OR type 'remove' to remove")
	fmt.Scanf("%s", &u)
	if u == "add" {
		add(todoList)
	} else if u == "remove" {
		remove(todoList)
	}

}
func add(todo []string) {
	var (
		t string
		n int
	)

	fmt.Println("Enter number of tasks")
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%s \n", &t)
		todo = append(todo, t)
	}

	for i := range todo {
		fmt.Printf("%v\t", todo[i])
	}
}

func remove(todo []string) {
	var i int
	fmt.Println("Enter index to delete")
	fmt.Scanf("%d", &i)
	todo = append(todo[:i], todo[i+1:]...)
	fmt.Println(todo)
}
