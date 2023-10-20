# Go Coffee Cost Calculator

## Description

The Go Coffee Cost Calculator is a command-line application that calculates the cost of a custom coffee based on its size, type, and flavor. It demonstrates the use of the decorator design pattern to add different features to the base coffee.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
    - [Getting Started](#getting-started)
- [How It Works](#how-it-works)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Ensure you have Go installed. If not, download and install it from [https://golang.org/doc/install](https://golang.org/doc/install).

2. Clone this repository to your local machine:

git clone <repository-url>

3. Navigate to the project directory:



4. Build the application:

## Usage

### Getting Started

1. Run the program with the following command:

    ./coffeeShop

    go run coffeeShop.go
    
2. Follow the on-screen prompts to input the size, type, and flavor of your custom coffee.

3. The program will calculate the total cost of your coffee and display it.

## How It Works

The Go Coffee Cost Calculator uses the decorator design pattern to add additional costs for coffee size, type, and flavor.

- `BaseCoffee` represents the base coffee.
- `SizeDecorator` adds cost based on the coffee's size.
- `TypeDecorator` adds cost based on the coffee's type.
- `FlavorDecorator` adds cost based on the coffee's flavor.

When you create a custom coffee, the program decorates the base coffee with additional features, and the total cost is calculated.

## Contributing

Contributions are welcome! If you'd like to contribute to the project, please follow these steps:

1. Fork the repository.

2. Create a new branch for your feature or bug fix.

3. Make your changes and commit them.

4. Push your changes to your fork.

5. Create a pull request to the main repository.

Please follow the [Contributor Covenant Code of Conduct](CONTRIBUTING.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.