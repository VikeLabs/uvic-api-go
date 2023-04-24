# UVIC-API

Started with VikeLabs' [StudySpaceFinder](https://github.com/VikeLabs/StudySpaceFinder), this project
provides an API for University of Victoria data, allowing future projects to access and utilize this
information. The currently used tech stack:

- Golang (v1.20)
  - [Chi](https://pkg.go.dev/github.com/go-chi/chi@v1.5.4): a small and simple router.
  - [GORM](https://gorm.io/): Object Relational Mapper for Golang.
- SQLite: for Study Space Finder data.
- Docker.

## Get started

- Assuming you have [installed Go](https://go.dev/doc/install).
- Clone the repo with:

```sh
git clone git@github.com:VikeLabs/uvic-api-go.git
```

### Basic Go

1. Install the dependencies:

```sh
go mod tidy
```

2. Install [air](https://github.com/cosmtrek/air), hot reloading module for Golang.

```sh
go install github.com/cosmtrek/air@latest
```

3. Run the server:

```sh
air
```

### Docker

```sh
docker compose up
```

## Contributing

This project is open source and contributions are welcome. If you encounter a bug or have
a feature request, please open an issue on the GitHub repository, and if you are interested
in becoming a maintainer, pm Scott (_@Closepanda#7203)_) or Hal (_haln#0584_) on Discord!

NOTE: In the interest of learning and sharing knowledge, the current tech stack is beginner-friendly,
frankly the Go language is beginner-friendly (have you used Rust before, _dammit_?!). You'd do well with
general knowledge of HTTP/HTTPS, JSON, SQL, and some basic knowledge of programming, really.

## License

This project is licensed under the [GNU GPLv3](https://choosealicense.com/licenses/gpl-3.0/)
License. See the [LICENSE](./LICENSE) file for details.
