package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson"
	// "bytes"
    // "text/tabwriter"
	"strings"
)

// Define the Item interface
type Item interface {
	GetPrice() float64
	GetDescription() string
}

type Product struct {
	Name     string
	Price    float64
	Category string // Add the Category field
}

type Burger struct {
	Price   float64
	Bun     bool
	Dressed bool
}

func (burger *Burger) GetPrice() float64 {
	return burger.Price
}

func (burger *Burger) GetDescription() string {
	description := fmt.Sprintf("Burger - Price: $%.2f, Bun: %v, Dressed: %v", burger.Price, burger.Bun, burger.Dressed)
	return description
}

type Drink struct {
	Price float64
	Type   string
}

func (drink *Drink) GetPrice() float64 {
	return drink.Price
}

func (drink *Drink) GetDescription() string {
	return fmt.Sprintf("Drink - Type: %s", drink.Type)
}

type Side struct {
	Price float64
	Type   string
}

func (side *Side) GetPrice() float64 {
	return side.Price
}

func (side *Side) GetDescription() string {
	return side.Type
}

type Combo struct {
	Items []Item
}

func (combo *Combo) GetPrice() float64 {
    total := 0.0
    for _, item := range combo.Items {
        total += item.GetPrice()
    }
    return total
}

func (combo *Combo) GetDescription() string {
    var description strings.Builder

    description.WriteString("You've selected:\n")
    description.WriteString("Combo:\n")

    totalPrice := 0.0
    for _, item := range combo.Items {
        itemDescription := item.GetDescription()
        itemPrice := item.GetPrice()

        itemLine := fmt.Sprintf("%-30s - Price: $%.2f\n", itemDescription, itemPrice)
        description.WriteString(itemLine)
        totalPrice += itemPrice
    }

    description.WriteString(fmt.Sprintf("Combo price: $%.2f\n", totalPrice))

    return description.String()
}




func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("mydb").Collection("products")

	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	var products []Product
	for cur.Next(context.Background()) {
		var product Product
		if err := cur.Decode(&product); err != nil {
			log.Printf("Error decoding product: %v", err)
			continue
		}
		products = append(products, product)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Products:")
	for _, product := range products {
		fmt.Printf("Name: %s, Price: $%.2f\n", product.Name, product.Price)
	}
	fmt.Println("WELCOME TO BURGHY")
	combo := &Combo{Items: []Item{}}

	for {
		menuUI(combo, products) 
		fmt.Println("Do you want to add another item? (y/n)")
		var response string
		_, err := fmt.Scanln(&response)
		if err != nil || (response != "y" && response != "Y") {
			break
		}
	}

	greetings(combo)
}

func menuUI(combo *Combo, products []Product) { // Add the products parameter
    fmt.Println("Choose an item")
    fmt.Println("1. Burger")
    fmt.Println("2. Drink")
    fmt.Println("3. Side")
    var choice int
    fmt.Println("Enter your choice (1/2/3)")
    _, err := fmt.Scanln(&choice)
    if err != nil {
        errorMessage()
    }

    var item Item

    switch choice {
    case 1:
        burger := createBurger()
        item = burger
    case 2:
        drink := createDrink(products, "Drink") // Pass the products slice and category
        item = drink
    case 3:
        side := createSide(products, "Side") // Pass the products slice and category
        item = side
    default:
        errorMessage()
    }

    combo.Items = append(combo.Items, item)
}

func createBurger() *Burger {
	fmt.Print("Do you want a bun? (y/n): ")
	var bunResponse string
	_, err := fmt.Scanln(&bunResponse)
	if err != nil {
		errorMessage()
	}

	fmt.Print("Do you want it dressed? (y/n): ")
	var dressedResponse string
	_, err = fmt.Scanln(&dressedResponse)
	if err != nil {
		errorMessage()
	}

	burger := &Burger{5.66, bunResponse == "y" || bunResponse == "Y", dressedResponse == "y" || dressedResponse == "Y"}
	return burger
}

func errorMessage() {
	fmt.Println("Invalid choice. Exiting.")
	os.Exit(1)
}

func greetings(combo *Combo) {
	fmt.Println("\n++++++++++++++++++")
	fmt.Println(combo.GetDescription())
	fmt.Printf("Combo price: $%.2f\n", combo.GetPrice())
	fmt.Println("++++++++++++++++++")
}

func createDrink(products []Product, category string) *Drink {
	fmt.Println("Choose a drink:")
	for i, product := range products {
		if product.Category == category {
			fmt.Printf("%d. %s\n", i+1, product.Name)
		}
	}
	var choice int
	fmt.Println("Enter your choice (1/2/3)")
	_, err := fmt.Scanln(&choice)
	if err != nil || choice < 1 || choice > len(products) {
		errorMessage()
	}

	drink := &Drink{products[choice-1].Price, products[choice-1].Name}
	return drink
}


func createSide(products []Product, category string) *Side {
	fmt.Println("Choose a side:")
	for i, product := range products {
		if product.Category == category {
			fmt.Printf("%d. %s\n", i+1, product.Name)
		}
	}
	var choice int
	fmt.Println("Enter your choice (1/2/3)")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		errorMessage()
	}

	var sideType string
	if choice >= 1 && choice <= len(products) && products[choice-1].Category == category {
		sideType = products[choice-1].Name
	} else {
		errorMessage()
	}

	side := &Side{products[choice-1].Price, sideType}
	return side
}
