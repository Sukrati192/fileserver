# Server to upload and download files using IPFS

APIs:

1. API for uploading a file:

```
  POST /upload

  Response: 
  {
    fileId: "",
    size: "",
    timestamp: ""
  }
```

2. API for serving a file:

```
  GET /file/:fileId

  Response: 
  <file content>
```


## Features

* Handles large files 1GB+
* Uses IPFS for storage

## Technologies

* Golang

## Setup

### Prerequisites
- Install Golang version 1.20: https://go.dev/doc/install
- Install IPFS: https://docs.ipfs.tech/install/command-line/#install-kubo-linux

### Start Server
* Run ipfs server: `make run-ipfs`
* In a new terminal, run fileserver: `make run`

## API Documentation

API documentation is maintained by using Swagger, automated by https://github.com/swaggo/swag

To update the document, annotate new APIs or update existing ones. After that, run:

`make generate-swagger`

This will update the swagger artifacts in the docs/ folder of the repo.

You can see the documentation locally at http://localhost:8000/swagger/index.html after running the go server.

