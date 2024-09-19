# Todo App CLI GO

## Ejecutar el proyecto

```bash
go run ./ -[flag]
```

### -list
Listar todas las tareas

```bash
go run ./ -list
```

### -add <título> string
Agregar nueva tarea pasando el título

```bash
go run ./ -add
```

### -edit <índice:título> int:string
Editar tarea pasando el índice y el nuevo título.

```bash
go run ./ -edit
```

### -toggle <índice> int
Especificar el índice de la tarea para cambiar el estado (default -1)

```bash
go run ./ -toggle
```

### -del <índice> int
Especificar el índice de la tarea a eliminar (default -1)

```bash
go run ./ -del
```