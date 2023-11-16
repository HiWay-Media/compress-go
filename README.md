# compress-go
[![Go build](https://github.com/HiWay-Media/compress-go/actions/workflows/go-build.yml/badge.svg)](https://github.com/HiWay-Media/compress-go/actions/workflows/go-build.yml)
[![Go test workflow](https://github.com/HiWay-Media/compress-go/actions/workflows/go-test.yml/badge.svg)](https://github.com/HiWay-Media/compress-go/actions/workflows/go-test.yml)

The Compress Go lang library provides access to the Compress API for encoding videos, restreamers

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)


## Features

- Easy-to-use API

## Installation

To install Compress-Go, you can use `go get`:

```shell
go get github.com/HiWay-Media/compress-go
```

## Usage
Here's a quick example of how to use Compress-Go

```go

func main(){
    //
    apiKey := os.Getenv("API_KEY")
    customerName := os.Getenv("CUSTOMER_NAME")
    //
    c, err := compress.NewCompress(customerName, apiKey, true)
    if err != nil {
        return nil, err
    }
}

``````

## Contributing
Contributions are welcome! If you would like to contribute to this project, please follow these steps:

## Fork the repository.
- Create a feature or bugfix branch.
- Make your changes.
- Create a pull request with a detailed description of your changes.
- Please refer to our Contributing Guidelines for more information.

## License
This project is licensed under the MIT License. See the LICENSE file for details.