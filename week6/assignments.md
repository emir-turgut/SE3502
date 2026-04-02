- **Exercise 1: The Concurrent Clock**
    - Create a function `tick(message string)` that prints the `message` to the console every 500 milliseconds (use `time.Sleep`).
    - In `main()`, launch `tick("Tick")` as a **goroutine**.
    - Immediately after, launch `tick("Tock")` as a **goroutine**.
    - Add a `fmt.Scanln()` or `time.Sleep` at the end of `main` to prevent the program from exiting immediately.

- **Exercise 2: The Ping-Pong Channel**
    - Create a channel of strings named `game`.
    - Create a function `pinger(c chan string)` that sends the word "ping" into the channel 5 times, then closes the channel.
    - Create a function `printer(c chan string)` that loops over the channel (using `range`) and prints whatever it receives.
    - In `main`, spawn `pinger` as a goroutine and call `printer` as a normal function call (so `main` waits for it to finish).

- **Exercise 3: The Safe Counter (Mutex)**
    - Create a global variable `count = 0`.
    - Write a function `increment(n int, wg *sync.WaitGroup)` that loops `n` times and increments `count` by 1 (`count++`).
    - In `main`, create a `sync.WaitGroup`.
    - Spawn two goroutines running `increment(1000, &wg)`.
    - Wait for them to finish and print the final value of `count`.
    - **The Bug:** Run this multiple times. You might see 2000, but often you will see 1950, 1500, etc. This is a race condition.
    - **The Fix:** Add a `sync.Mutex` to protect the increment operation. Pass the mutex (as a pointer!) or make it global, and lock/unlock around `count++`.

- **Exercise 4: The Fast and Slow Workers**
    - Create a buffered channel `jobs` with capacity 5.
    - Launch a "Fast Worker" goroutine that reads from `jobs` and sleeps only 100ms per job.
    - Launch a "Slow Worker" goroutine that reads from `jobs` and sleeps 500ms per job.
    - In `main`, send 10 integers (1 to 10) into the `jobs` channel.
    - Close the `jobs` channel after sending.
    - Use a `sync.WaitGroup` to ensure `main` does not exit until both workers are done processing.