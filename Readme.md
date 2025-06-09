# Go Bootcamp By Teeranuwat

Welcome to the **Go Bootcamp**! This guide covers essential Go concepts, practical assignments, and advanced topics to help you become proficient in Go development.

---

link from P'Bond
Check out [P'Bond's curated Go resources and sample projects on GitHub](https://github.com/codebangkok/golang?tab=readme-ov-file) — a treasure trove for Go learners and enthusiasts!

---

For those seeking to master their Go skills, check out the [Final Assignment: Jade Palace By P'หน่อง](https://github.com/AnuchitO/softcraft-valley/blob/main/Jade-Palace.md#learning-objectives) — a comprehensive challenge to solidify your learning.

## 📅 Day One

### Topics Covered

- **Go Install & Run**
- **Syntax & Types**
- **Basic Go Program Structure**
- **Printing to Console**
- **Functions in Go**
- **Packages**
- **Building a Basic REST API (Native Libraries)**

---

### 📝 Assignments

#### Assignment I: FizzBuzz Program

- Write a function that takes a number as input.
- Use a `for` loop and the function to print the FizzBuzz sequence up to that number.
- Print:
    - `"Fizz"` for multiples of 3
    - `"Buzz"` for multiples of 5
    - `"FizzBuzz"` for multiples of both

#### Assignment II: FizzBuzz with Slice Input

- Write a function that takes a slice of integers as input.
- For each integer, use FizzBuzz logic to print the appropriate value (`"Fizz"`, `"Buzz"`, `"FizzBuzz"`, or the number).
- Reuse the function from Assignment I.

#### Assignment III: Native Libraries REST API

- Build a simple REST API using Go's standard library (`net/http`).
- Implement at least two endpoints (e.g., `GET /ping`, `POST /reverse`).
- Use JSON for request and response bodies.
- Add basic error handling and return appropriate HTTP status codes.
- Test your endpoints using `curl` or Postman.

--- 

## 📅 Day Two

### Topics Covered
- **Introduction to Gin Framework**
- **Unit Testing** gomock ควร cover ใน จำนวน case client ใช้จริงๆ 
- **Hexagonal Architecture** Port Adapter
- **Interfaces & Implementations**

#### Assignment I: Gin REST API

- Create a simple REST API using the Gin framework.
- Implement at least two endpoints (e.g., `GET /hello`, `POST /echo`).
- Use JSON for request and response bodies.
- Add basic error handling and return appropriate HTTP status codes.
- Test your endpoints using `curl` or Postman.

#### Assignment II: Gin Handler Registration

- Refactor your Gin REST API by creating a `handler.go` file.
- Move all handler functions (e.g., for `/hello`, `/echo`) into `handler.go`.
- In your main application file, import and register these handlers with the Gin router.
- Ensure your project structure is modular and handlers are easy to maintain.
- Test your endpoints to confirm correct handler registration and functionality.

---

## 📅 Day Three

### Topics Covered

- **Recap Unit test**
- **Database**
- **query เบื้องต้น**
- **sqlc**
- ***sqlmock**
- **Logging with Slog**

ถ้าเวลาเหลือ
- **Redis Connections**
- **Templates พี่ยอด**
- **Coding Conventions**
- **sqlmock**


ถ้าทันก็เกริ่นๆ พูดไป

- **Kafka Connection**
- **Consumer & Producer Patterns**

---

Happy Coding! 🚀

---

## 🎁 Bonus Learning Resources

- [Official Go Documentation](https://go.dev/)
- [Go by Example](https://gobyexample.com/)

Explore these resources to deepen your understanding of Go and practice with real-world examples!


Note : Database HW Postgres
