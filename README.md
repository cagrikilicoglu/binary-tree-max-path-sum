# Maximum Path Sum API

![GitHub followers](https://img.shields.io/github/followers/cagrikilicoglu?label=Follow)

Maximum path sum API calculates maximum path sum of a given binary tree. A path is a collection of connected nodes in a tree, where no node is connected to more than two other nodes; a path sum is the sum of the values of the nodes in a particular path.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- You have installed the latest version of Go.

## Installing Maximum Path Sum API

To clone Maximum Path Sum API to your repo follow these steps:

```
git clone https://github.com/cagrikilicoglu/binary-tree-max-path-sum
```

## Repository Overview

```bash
├── README.md
├── docs
│   └── tree.yaml
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   ├── node.go
│   │   ├── request.go
│   │   ├── result.go
│   │   └── tree.go
│   └── tree
│       ├── handler.go
│       ├── handler_test.go
│       ├── service.go
│       └── service_test.go
├── main.go
└── pkg
    ├── config
    │   ├── config.go
    │   └── local.yaml
    └── graceful
        └── shutdown.go

```

## Using Maximum Path Sum API

To use Maximum Path Sum API, follow these steps:

The app provides the following endpoint:

- `POST /api/v1/binary-tree/max-path-sum` : finds maximum path sum of binary tree supplied in request body. <br>Example request: `POST /api/v1/binary-tree/max-path-sum`<br>requests body: {
  "tree": {
  "nodes": [
  {"id": "1", "left": "2", "right": "3", "value": 1},
  {"id": "3", "left": "6", "right": "7", "value": 3},
  {"id": "7", "left": null, "right": null, "value": 7},
  {"id": "6", "left": null, "right": null, "value": 6},
  {"id": "2", "left": "4", "right": "5", "value": 2},
  {"id": "5", "left": null, "right": null, "value": 5},
  {"id": "4", "left": null, "right": null, "value": 4}
  ],
  "root": "1" }
  }

## Tool set

- Go
- Gin
- Viper
- Swagger

## Contributors

Thanks to the following people who have contributed to this project:

- [@cagrikilicoglu](https://github.com/cagrikilicoglu) 📖

## Links

[project repository](https://github.com/cagrikilicoglu/binary-tree-max-path-sum)

## License

This project uses the following license: [MIT](https://opensource.org/licenses/MIT).
