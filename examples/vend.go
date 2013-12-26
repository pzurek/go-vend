package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/pzurek/go-vend/tools"
	"github.com/pzurek/go-vend/vend"
)

// to run:
// go run vend.go <store name> <username> <password>
func main() {
	args := os.Args
	storeName := args[1]
	username := args[2]
	password := args[3]

	client := vend.NewClient(storeName, username, password, nil)

	products, err := client.Products.List()
	if err != nil {
		log.Fatalln(err)
	}

	taxes, err := client.Taxes.List()
	if err != nil {
		log.Fatalln(err)
	}

	outlets, err := client.Outlets.List()
	if err != nil {
		log.Fatalln(err)
	}

	registers, err := client.Registers.List()
	if err != nil {
		log.Fatalln(err)
	}

	users, err := client.Users.List()
	if err != nil {
		log.Fatalln(err)
	}

	suppliers, err := client.Suppliers.List()
	if err != nil {
		log.Fatalln(err)
	}

	customers, err := client.Customers.List()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("You have downloaded:")
	fmt.Printf("%v products\n", len(products))
	fmt.Printf("%v taxes\n", len(taxes))
	fmt.Printf("%v outlets\n", len(outlets))
	fmt.Printf("%v registers\n", len(registers))
	fmt.Printf("%v users\n", len(users))
	fmt.Printf("%v suppliers\n", len(suppliers))
	fmt.Printf("%v customers\n", len(customers))

	productMap := buildProductMap(products)
	fmt.Printf("%v unique products\n", len(productMap))

	// buildStockLevelsCsv(products, outlets)
}

func buildProductMap(products []vend.Product) map[string]vend.Product {
	pMap := make(map[string]vend.Product)

	for _, product := range products {
		pMap[*product.Id] = product
	}

	return pMap
}

func buildRegisterMap(registers []vend.Register) map[string]vend.Register {
	rMap := make(map[string]vend.Register)

	for _, register := range registers {
		rMap[*register.Id] = register
	}

	return rMap
}

func buildOutletMap(outlets []vend.Outlet) map[string]vend.Outlet {
	oMap := make(map[string]vend.Outlet)

	for _, outlet := range outlets {
		oMap[*outlet.Id] = outlet
	}

	return oMap
}

func buildInvenotryMap(inventoryItems []vend.InventoryItem) map[string]vend.InventoryItem {
	iMap := make(map[string]vend.InventoryItem)

	for _, item := range inventoryItems {
		iMap[*item.OutletId] = item
	}

	return iMap
}

func buildStockLevelsCsv(products []vend.Product, outlets []vend.Outlet) {
	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	monthStr := strconv.Itoa(month)
	dayStr := strconv.Itoa(day)
	hourStr := strconv.Itoa(hour)
	minuteStr := strconv.Itoa(minute)
	secondStr := strconv.Itoa(second)

	monthStr = tools.PadLeft(monthStr, "0", 2)
	dayStr = tools.PadLeft(dayStr, "0", 2)
	hourStr = tools.PadLeft(hourStr, "0", 2)
	minuteStr = tools.PadLeft(minuteStr, "0", 2)
	secondStr = tools.PadLeft(secondStr, "0", 2)

	fname := fmt.Sprintf("vend_stock_levels_%v_%v_%v_%v_%v_%v.csv", year, monthStr, dayStr, hourStr, minuteStr, secondStr)
	f, err := os.Create(fmt.Sprintf("./%s", fname))
	if err != nil {
		log.Fatal("Error: %s", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)

	/*
	   Product,Sku,Outlet,Count,Value,Avg. Item Value,Reorder point
	*/

	var headerLine []string
	headerLine = append(headerLine, "Product")
	headerLine = append(headerLine, "Sku")
	headerLine = append(headerLine, "Outlet")
	headerLine = append(headerLine, "Count")
	headerLine = append(headerLine, "Value")
	headerLine = append(headerLine, "Avg. Item Value")
	headerLine = append(headerLine, "Reorder point")

	w.Write(headerLine)

	for _, product := range products {

		// Skip the product if it has no invenotry or is a composite
		if product.Inventory == nil || product.Composites != nil {
			continue
		}

		name := *product.Name
		sku := *product.Sku
		cost := parseFloat(*product.SupplyPrice)
		costStr := formatFloat(cost)
		inventoryMap := buildInvenotryMap(*product.Inventory)

		for _, outlet := range outlets {

			oName := *outlet.Name
			count := 0.0
			rPoint := 0.0

			if item, ok := inventoryMap[*outlet.Id]; ok {
				count = parseFloat(*item.Count)
				rPoint = parseFloat(*item.ReorderPoint)
			}

			value := cost * count

			countStr := formatFloat(count)
			valueStr := formatFloat(value)
			rPointStr := formatFloat(rPoint)

			var record []string
			record = append(record, name)      // Product
			record = append(record, sku)       // Sku
			record = append(record, oName)     // Outlet
			record = append(record, countStr)  // Count
			record = append(record, valueStr)  // Value
			record = append(record, costStr)   // Avg. Item Value
			record = append(record, rPointStr) // Reorder point

			w.Write(record)
		}
	}

	w.Flush()
}

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', 5, 32)
}

func parseFloat(s string) float64 {

	if s == "" {
		return 0.00
	} else {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			log.Fatal("Error: %s", err)
		}
		return f
	}
}
