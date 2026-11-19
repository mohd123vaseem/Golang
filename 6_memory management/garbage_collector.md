## Garbage Collector (GC) in Golang

Think of Go's **Garbage Collector (GC)** as an automatic cleanup crew that runs quietly in the background while your program is working.

In older languages like C or C++, when you reserve memory (like using `new()` or `make()`), **you** have to remember to manually free that memory when you're done with it. If you forget, your program eats up more and more RAM until it crashes (a "memory leak").

In Go, you don't have to do that manually. The Garbage Collector handles it for you.

---

### How It Works (The Simple Version)

1. **Tracking What You Use:** Every time you create variables, slices, or maps, Go places them in memory and keeps track of which parts of your code are still pointing to them.
2. **Finding Unused Data ("Garbage"):** Once a function finishes running or a variable is no longer referenced anywhere in your code, that data becomes unreachable—essentially "garbage."
3. **Sweeping It Away:** Periodically, the Garbage Collector scans your program's memory, marks everything that is still actively being used, and safely deletes everything else so that memory can be reused for new variables.

---

### Why Go's Garbage Collector is Special

* **No Big Freezes:** In some languages, the garbage collector has to completely pause your entire app ("Stop-the-World") while it cleans up, which can cause noticeable lag. Go's GC is **concurrent**, meaning it does almost all of its cleanup *at the same time* your code is running, keeping pauses tiny (usually less than a millisecond).
* **Zero Configuration Needed:** Out of the box, Go automatically decides when to run the cleanup based on how fast your program is using memory—you rarely ever need to touch or tweak it as a beginner.

---

### Simple Example

```go
package main

import "fmt"

func createOrder() {
    // 1. Go allocates memory for this slice
    items := make([]string, 3)
    items[0] = "Laptop"
    fmt.Println("Processing:", items[0])

    // 2. Function ends here. 'items' is never returned or used again.
}

func main() {
    createOrder()
    // 3. The GC notices nobody can reach 'items' anymore and automatically frees that memory!
}
```