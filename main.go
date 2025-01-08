package main

import (
	"fmt"
	"net/http"
)

var firstTask string = "Submit a report"
var secondTask = "Office meeting"
var thirdTask = "Go to Supermarket"
var fourthTask =  "Get a haircut"
var fifthTask = "Go on a trip"
	
var taskItems =[]string {firstTask,secondTask,thirdTask,fourthTask,fifthTask}

func main(){


	http.HandleFunc("/",helloUser)
	http.HandleFunc("/showTasks", showTasks)

	http.ListenAndServe(":8080",nil)
	// fmt.Println()
 
	// printTasks(taskItems)
	// taskItems= addTasks(taskItems,"Go for a run")
	
	// fmt.Println("updated list")
	// printTasks(taskItems)

}

func helloUser(writer http.ResponseWriter, request *http.Request){
	var greeting = Greet("Visal")
	fmt.Fprintln(writer, greeting)
}
func showTasks (writer http.ResponseWriter, request *http.Request){
	for _, task:= range taskItems{
		fmt.Fprintln(writer, task)
	}
	
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