package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Agregar nueva tarea pasando el título")
	flag.StringVar(&cf.Edit, "edit", "", "Editar tarea pasando elíndice y el nuevo título. índice:título")
	flag.IntVar(&cf.Del, "del", -1, "Especificar el índice de la tarea a eliminar")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Especificar el índice de la tarea para cambiar el estado")
	flag.BoolVar(&cf.List, "list", false, "Listar todas las tareas")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: Debe especificar el índice y el nuevo título. índice:título")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: El índice debe ser un número")
			os.Exit(1)
		}
		todos.edit(index, parts[1])
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.Del != -1:
		todos.delete(cf.Del)
	default:
		fmt.Println("Error: Debe especificar una acción válida")
	}
}
