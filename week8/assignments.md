- **Exercise 1: The "Strict" Config Loader**
    
    **The JSON Input:**
    
    ```go
    {
      "server_port": 8080,
      "db_host": "localhost",
      "DB_PASSWORD": "super_secret_password",
      "feature_flags": {
        "beta_mode": true,
        "deprecated_v1": false
      },
      "optional_webhook": ""
    }
    ```
    
    - Define a set of Go structs (`Config` and `FeatureFlags`) to map this JSON data.
    - Use **Struct Tags** to:
        - Map `server_port` to a field named `Port`.
        - Map `DB_PASSWORD` to a field named `DatabasePass`.
        - Ensure `DatabasePass` is **never** outputted when the struct is marshaled back to JSON (use the ignore tag).
        - Ensure `optional_webhook` is omitted from the output if it is empty.
    - Write a `main` function that unmarshals the raw JSON string into your struct, prints the struct details, and then marshals it back to JSON to demonstrate that the password is gone.
- **Exercise 2: The Inventory Manager CLI**
    
    **The Interface to Implement:**
    
    ```go
    type InventoryStore interface {
        // AddProduct inserts a new item. Returns the new ID.
        AddProduct(name string, price int, stock int) (int, error)
        
        // UpdateStock changes the stock level for a product.
        // It should handle cases where the product does not exist.
        UpdateStock(id int, newStock int) error
        
        // GetLowStock retrieves all products where stock < threshold.
        GetLowStock(threshold int) ([]Product, error)
        
        // DeleteProduct removes a product by ID.
        DeleteProduct(id int) error
    }
    ```
    
    - Setup a SQLite database with a table `products` (columns: `id`, `name`, `price`, `stock`).
    - Implement the `InventoryStore` interface using the `database/sql` package.
    - **Critical Security Requirement:** You must use **Prepared Statements** (placeholders `?`) for all arguments. String concatenation is strictly forbidden.
    - Write a `main` function that simulates a workflow:
        - Add 3 items (e.g., "Laptop", "Mouse", "Keyboard").
        - Update the stock of the "Mouse".
        - Query for low stock items.
        - Delete the "Laptop".