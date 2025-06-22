package main

import (
	"fmt"
	"bufio"
	"os"
)

type Task struct {
	id int
	title string
	done bool
}

func options() int {
	fmt.Println("\n1. Criar nova tarefa")
	fmt.Println("2. Listar tarefas")
	fmt.Println("3. Marcar tarefa como feita")
	fmt.Println("4. Sair do programa")

	fmt.Print("Digite uma opção: ")

	var choice int
	fmt.Scan(&choice)

	return choice
}

func new_task(tasks []Task) []Task {
	fmt.Print("Insira o título da nova tarefa: ")
	buffer := bufio.NewReader(os.Stdin)
	title, _ := buffer.ReadString('\n')
	title = title[:len(title)-1]

	id := len(tasks) + 1

	task := Task{
		id: id,
		title: title,
		done: false,
	}

	tasks = append(tasks, task)

	fmt.Printf("Tarefa '%s' criada com o id: %d", title, id)

	return tasks
}

func view_tasks(tasks []Task) {
	if(len(tasks) == 0) {
		fmt.Println("Sem tarefas cadastradas!")
	} else {
		for _, task := range tasks {
			var status string

			if task.done {
				status = "✔️"
			} else {
				status = " "
			}

			fmt.Printf("%d. %s [%s]\n", task.id, task.title, status)
		}
	}
}

func make_task_done(tasks []Task) []Task {
	fmt.Print("Insira o Id da tarefa marcada como feita: ")
	var id int
	fmt.Scan(&id)

	if len(tasks) >= id {
		tasks[id - 1].done = true
	} else {
		fmt.Println("O Id indicado não existe!")
	}

	return tasks
}

func main() {
	var tasks []Task

	fmt.Print("\033[H\033[2J")

	fmt.Println("Bem-vindo a sua lista de tarefas!")

	for {
		choice := options()

		fmt.Print("\033[H\033[2J")

		switch choice {
		case 1:
			tasks = new_task(tasks)
		case 2:
			view_tasks(tasks)
		case 3:
			tasks = make_task_done(tasks)
		case 4:
			fmt.Println("Até mais!")
			return
		default:
			fmt.Println("Comando inválido!")
		}

		fmt.Print("\n")
	}
}
