//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price float64
}

func main() {
	var shoppingList [4]Product
	milk := Product{"Milk", 2.99}
	bread := Product{"Bread", 1.99}
	eggs := Product{"Eggs", 3.99}

	shoppingList[0] = milk
	shoppingList[1] = bread
	shoppingList[2] = eggs

	printList(shoppingList)
	shoppingList = addItem(shoppingList, "Butter", 2.99)
	printList(shoppingList)
}

func addItem(shoppingList [4]Product, name string, price float64) [4]Product {
	product := Product{name, price}
	for i := 0; i < len(shoppingList); i++ {
		if shoppingList[i].name == "" {
			shoppingList[i] = product
			break
		}
	}
	return shoppingList
}

func printList(shoppingList [4]Product) {
	var total float64
	var totalItems int
	var lastItem Product

	for i := 0; i < len(shoppingList); i++ {
		if shoppingList[i].name == "" {
			lastItem = shoppingList[i-1]
			break
		}
		total += shoppingList[i].price
		totalItems++
	}
	if lastItem.name == "" {
		lastItem = shoppingList[0]
	}

	fmt.Printf("Last item: %s\n", lastItem.name)
	fmt.Printf("Total number of items: %d\n", totalItems)
	fmt.Printf("Total cost of items: $%.2f\n\n", total)
}
