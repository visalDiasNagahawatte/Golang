package main

import "fmt"



func main(){

	var firstTask = "Submit a report"
	var secondTask = "Office meeting"
	var thirdTask = "Go to Supermarket"
	var fourthTask =  "Get a haircut"
	var fifthTask = "Go on a trip"
	
	var taskItems =[]string {firstTask,secondTask,thirdTask,fourthTask,fifthTask}

    fmt.Println("Welcome to To-Do List App")
	fmt.Println()
 
	printTasks(taskItems)
	taskItems= addTasks(taskItems,"Go for a run")
	
	fmt.Println("updated list")
	printTasks(taskItems)

}

func printTasks(taskItems []string){
	fmt.Println("My To-do List")
	fmt.Println()

	for index, task := range taskItems{
		fmt.Printf("%d. %s\n",index+1,task)

	} 
}

func addTasks(taskItems []string, newTask string) []string{
     var updatedTask = append(taskItems,newTask)
	 return updatedTask
}