# Go Learning Syllabus

A focused syllabus based on the selected YouTube playlist, with a short list of supplementary topics.

**Goal:** Learn Go efficiently for practical development and Go-related SDE interviews without watching the entire playlist.

Playlist: https://www.youtube.com/watch?v=JoJ8Sw5Yb4c&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa

## 1. Playlist videos to watch

Watch these sections:

| Video numbers | Topics |
|---|---|
| 3–8 | Go fundamentals, types, packages, comma-ok |
| 11–17 | Memory, pointers, arrays, slices, maps, structs |
| 21–23 | Functions, methods, defer |
| 25–32 | HTTP, URLs, requests, JSON |
| 34–40 | REST APIs |
| 52–56 | Goroutines, WaitGroups, mutexes, race conditions, channels |

## 2. Videos to skip or skim

| Video numbers | Topics | Recommendation |
|---|---|---|
| 1–2 | Introduction | Skip |
| 9–10 | Time handling, cross-platform builds | Skim / learn when needed |
| 18–20 | Basic conditionals and loops | Skip or skim; concepts are familiar from C++ |
| 24 | File handling | Learn when needed |
| 33 | Go modules | Learn the basics |
| 41–51 | MongoDB implementation | Skip for now |
| 57 | Miscellaneous utilities | Learn when needed |

## 3. Supplementary topics

The playlist is a useful foundation, but it is not fully sufficient for interview preparation. Cover these topics separately:

- [ ] **Interfaces:** implicit implementation, interface composition, nil interfaces
- [ ] **Error handling:** error wrapping, `errors.Is`, `errors.As`
- [ ] **Context:** cancellation, deadlines, propagation
- [ ] **Channels:** buffered vs. unbuffered, nil channels, closed channels, `select`
- [ ] **Concurrency:** goroutine leaks, worker pools, synchronization
- [ ] **Generics:** basic syntax and use cases
- [ ] **Testing:** unit tests, table-driven tests, benchmarks
- [ ] **Memory:** escape analysis, garbage collection basics

## 4. Recommended resources

Use these as references rather than starting another long video course:

1. [A Tour of Go](https://go.dev/tour/)
2. [Effective Go](https://go.dev/doc/effective_go)
3. [Official Go documentation](https://go.dev/doc/)

## 5. Suggested study method

For each topic:

1. Watch the relevant lesson.
2. Write a small program yourself—avoid simply copying the instructor.
3. Test edge cases and understand the behaviour.
4. Practise a few interview-style questions.
5. Revisit the concept in your own code when possible.

## 6. Completion checklist

- [ ] Watch the selected playlist sections.
- [ ] Practise slices, maps, structs, pointers and methods.
- [ ] Understand interfaces and error handling.
- [ ] Practise goroutines, channels and synchronization.
- [ ] Understand context cancellation and goroutine leaks.
- [ ] Learn basic generics and Go testing.
- [ ] Review memory management and escape analysis.

**Bottom line:** Watch the selected playlist sections, then spend approximately 3–4 focused days covering the supplementary topics. The playlist alone is not enough for a strong interview-ready Go foundation, but this combination is a focused starting point.
