<div align="center">
<img src="assets/in-memory-key-value-store.svg" height="auto" width="400" />
<br />
<h1>In-Memory Key-Value Store</h1>
<p>
An in-memory key-value store HTTP API Service
</p>
</div>

## About

An in-memory key-value store HTTP API Service that supports:

- `/get/<key>`: Return the value of the key
- `/search`: Search for keys using the `prefix` and `suffix` filters
- `/set`: Post call which sets the key/value pair

#### Folder structure:

```shell
.
├── Dockerfile
├── LICENSE
├── README.md
├── assets
│   ├── 1.svg
│   ├── 2.png
│   ├── 3.webp
│   ├── 4.svg
│   ├── 5.svg
│   └── in-memory-key-value-store.svg
├── docker-compose.yaml
├── go.mod
├── go.sum
├── handlers
│   ├── handlers.go
│   └── handlers_test.go
├── main.go
├── main.tf
└── prometheus
    └── prometheus.yml
```

## Usage

Start the server using the command:

```shell
go run main.go
```

TO run the tests:

```shell
go test -v ./handlers/
```

This project is deployed on Google Cloud run at https://in-memory-key-value-store-44nkealrda-uc.a.run.app.

<img src="assets/1.svg" height="auto" width="100" /> <img src="assets/2.png" height="auto" width="100" /> <img src="assets/3.webp" height="auto" width="100" /> <img src="assets/4.svg" height="auto" width="100" /> <img src="assets/5.svg" height="auto" width="100" />

## License

[MIT](https://github.com/iamrajiv/in-memory-key-value-store/blob/main/LICENSE)
