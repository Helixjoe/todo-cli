# todo-cli

This is a CLI Based Todo App built using Go

### Packages used:
1. `flag` - command flags
2. `aquasecurity/table` - todo table

### Steps followed during development

1. `go mod init todo` - initialized the todo module
2. `go get github.com/aquasecurity/table` - installs table package
3. `go build -o todo` - bulids the todo app into an executable

### Running the app

1. `go build -o todo` - builds the app
2. `alist todo="./todo"` - helps in running todo command directly
3. `todo -list` - list all todos
4. `todo -add "TODO ITEM"` - add todo item
5. `todo -del index` - delete todo item
6. `todo -edit index:"TODO EDIT"` - edit todo item
7. `todo -toggle index` - toggle todo item from incomplete to complete 
