# UVIC API

Started with VikeLabs'
[StudySpaceFinder](https://github.com/VikeLabs/StudySpaceFinder), this project
provides an API for University of Victoria data, allowing future
projects to access and utilize this information.

## Core dependencies

- Golang (v1.20)
  - [Chi](https://pkg.go.dev/github.com/go-chi/chi@v1.5.4): a small and simple
    router.
- Docker.

## Get started

- Start a dev server with docker!

  ```sh
  docker compose up
  ```

- If Docker is not an option, assume Go is
  [installed](https://go.dev/doc/install):

  ```sh
  go mod tidy # sync dependecies
  go install github.com/cosmtrek/air@latest # for hot reload
  air # start the server
  ```

- Server listens at [http://localhost:8080](http://localhost:8080).

## Contributing

This project is open source and contributions are welcome. If you encounter a
bug or have a feature request, please open an issue on the GitHub repository.

## License

This project is licensed under the **GNU GPLv3** License. See the
[LICENSE](./LICENSE) file for details.
