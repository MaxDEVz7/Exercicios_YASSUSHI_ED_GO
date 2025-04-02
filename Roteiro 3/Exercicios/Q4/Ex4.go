package main

import (
	"fmt"
)

// Estrutura do nó (cada tarefa)
type Task struct {
	task string
	next *Task
}

// Estrutura da lista encadeada (lista de tarefas)
type TaskList struct {
	Head *Task // um elemento do tipo ponteiro de uma task, ou seja o elemento é do tipo task
}

// Método para adicionar uma nova tarefa à lista (Append)
func (t *TaskList) AddTask(name string) {
	newTask := &Task{task: name}
	if(t.Head == nil) {
		t.Head = newTask
		return
	}
	current := t.Head
	for current.next != nil {
		current = current.next
	}
	current.next = newTask
}

// Método para remover a primeira tarefa (RemoveFirst)
func (t *TaskList) CompleteTask() {
	if (t.Head!= nil){
		t.Head = t.Head.next
	}
}

// Método para exibir todas as tarefas pendentes
func (t *TaskList) ShowTasks() {
	current := t.Head
	for current != nil {
		fmt.Println(current.task)
		current = current.next
	}
}
func main() {
	lista := TaskList{}
	// Adicionando tarefas
	lista.AddTask("Estudar Go")
	lista.AddTask("Fazer compras")
	lista.AddTask("Academia")
	lista.ShowTasks()    // Exibe as tarefas pendentes
	lista.CompleteTask() // Remove a primeira tarefa concluída
	fmt.Println("-------------------------")
	lista.ShowTasks()    // Exibe as tarefas restantes
	lista.CompleteTask()
	fmt.Println("-------------------------")
	lista.ShowTasks()
}
