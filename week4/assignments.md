- **Exercise 1: The Bank Account (Pointer vs. Value Receivers)**
    - Define a struct `BankAccount` with a field `balance` (int).
    - Create a method `Deposit(amount int)` that increases the balance.
    - Create a method `Withdraw(amount int)` that decreases the balance (if funds are sufficient).
    - Create a method `GetBalance()` that returns the current balance.
    - **Critical Thinking:** Which methods need a pointer receiver `BankAccount` and which can use a value receiver `BankAccount`?
- **Exercise 2: The RPG Character (Composition)**
    - Create a basic struct `Stats` with fields `Level` (int) and `Health` (int).
    - Create a struct `Character` that **embeds** `Stats` and adds a `Name` (string) field.
    - In `main`, instantiate a `Character`.
    - Demonstrate "field promotion" by accessing `Level` directly on the `Character` variable (e.g., `player.Level`, not `player.Stats.Level`).
- **Exercise 3: The Shape Calculator (Polymorphism)**
    - Define an interface `Shape` with two methods: `Area() float64` and `Perimeter() float64`.
    - Create a struct `Rectangle` (width, height) and `Circle` (radius).
    - Implement the `Shape` interface for both.
    - Write a function `PrintShapeInfo(s Shape)` that accepts the interface and prints the results.
- **Exercise 4: The Stringer Interface (Standard Library)**
    
    ```go
    type Stringer interface {
        String() string
    }
    ```
    
    - Create a struct `Book` with `Title`, `Author`, and `Year`.
    - If you print this struct using `fmt.Println(myBook)`, it prints standard struct syntax `{Title Author Year}`.
    - Add a method `String() string` to `Book` that returns a formatted string like: *"Title by Author (Year)"*.
    - Run `fmt.Println` again to see your custom format.
- **Exercise 5: Dependency Injection (Refactoring)**
    
    ```go
    type PayPal struct {}
    func (p PayPal) Pay(amount float64) { fmt.Println("Paid via PayPal") }
    
    type Store struct {
        PaymentMethod PayPal // Tightly coupled!
    }
    ```
    
    - Create an interface `PaymentProcessor` with a `Pay(amount float64)` method.
    - Refactor `Store` to hold a `PaymentProcessor` instead of a concrete `PayPal` struct.
    - Inject `PayPal` (and a hypothetical `Stripe` struct) into the `Store` at runtime.