Take a look at the program [matching.go](src/matching.go). Explain what happens and why it happens if you make the following changes. Try first to reason about it, and then test your hypothesis by changing and running the program.

- What happens if you remove the `go-command` from the `Seek` call in the `main` function?

  - **Answer**: If you remove the `go` keyword from the Seek call in the main function, the program will run sequentially instead of concurrently.

- What happens if you switch the declaration `wg := new(sync.WaitGroup)` to `var wg sync.WaitGroup` and the parameter `wg *sync.WaitGroup` to `wg sync.WaitGroup`?

  - **Answer**: If you make the above mentioned change you'll be passing a copy of the sync.WaitGroup to the Seek function instead of a pointer. This would result in each goroutine working with its own local copy of the wait group which would render the main functions wg.Wait() useless as it would return immediately without waiting for the goroutines to complete. This would likely result in incorrect behavior.

- What happens if you remove the buffer on the channel match?

  - **Answer**: If you remove the buffer on the match channel it will become an unbuffered channel. Which means that the send and receive operations on the channel will be synchronous, meaning that the sender and receiver must be ready at the same time to perform the operation. This will likely lead to a deadlock in the program because there will be no free space in the channel for unmatched sends.

- What happens if you remove the default-case from the case-statement in the `main` function?
  - **Answer**: If you remove the default-case from the case-statement the select statement will block until there is a pending send operation in the match channel. If there is no unmatched send remaining in the channel the program will halt indefinitely while waiting for something that will never occur.

Hint: Think about the order of the instructions and what happens with arrays of different lengths.
