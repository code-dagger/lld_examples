package main

import "github.com/code-dagger/in-mem-sql-db/database"

func main() {
	mgr := database.NewManager()
	mgr.CreateTable("products", map[string]string{"name": "string", "sku_id": "string", "stock": "int"})
	mgr.CreateTable("orders", map[string]string{"name": "string", "product_id": "int", "quantity": "int"})

}
