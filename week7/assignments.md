**Task:** Build a "Micro-Registry" without frameworks.

1. Create a server listening on port 9000.
2. Use `http.NewServeMux`.
3. Implement `POST /register`: Accepts a JSON body (you must decode it manually using `encoding/json`).
4. Implement `GET /user/{name}`: Returns the user details if registered.
5. **Constraint:** Do not use any external packages (like Gin, Fiber, or Chi).