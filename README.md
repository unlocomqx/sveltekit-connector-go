# SvelteKit Connector for Go

This is a proof of concept showing how it can be possible to connect a SvelteKit frontend to a Go backend using remote
functions.

## Usage

```shell
bun install
bun run go ;start go server
bun run dev
```

## Conventions

Prefix your go function with `Query`, `Form`, or `Command` to indicate its type. For example:- `QueryTodos`

- `FormCreateTodo`
- `CommandDeleteTodo`
- `QueryTodoByID`
- `CommandUpdateTodo`

## Contributing

This is a POC, feel free to contribute to make it better.

## Todos

- [ ] Handle function parameters
- [ ] Generate registry.go automatically or implement dynamic imports somehow
- [ ] Fix types for "virtual" js module
- [ ] Improve Go server structure
- [ ] Improve error handling
- [ ] Reload Vite server when a go file is changed
