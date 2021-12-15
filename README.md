# Test with Go
Workbook of the course Test With Go

**Name of branches equal to lessons. Parts of a lesson equal the commit.**

###—> 1. What is testing, and why does it matter 0'38''
- 1.1 What is a test?
- 1.2 Why do tests matter
- 1.3 Writing great tests

###—> 2. Tests are just Go code 0'22''
- 2.1 Testing with a main package
- 2.2 Testing with Go testing package
- 2.3 What happens when we run go test

###—> 3. Naming conventions 0'17''
- 3.1 File naming conventions
- 3.2 Function naming conventions
- 3.3 Variable naming conventions

###—> 4. Failing tests 0'44''
- 4.1 Ways to signal test failure
- 4.2 When to use Error vs Fatal
- 4.3 Writing useful failure messages

###—> 5. Examples as test cases 0'37''
- 5.1 A basic example as a test case
- 5.2 Viewing examples in the docs
- 5.3 Unordered example output
- 5.4 Complex examples
- 5.5 Examples in the standard library

###—> 6. Testing multiple cases 0'50''
- 6.1 Table driven tests
- 6.2 Generating table driven test code
- 6.3 Subtests
- 6.4 Shared setup and teardown
- 6.5 TestMain

###—> 7. Parallel tests 0'32''
- 7.1 Running tests in parallel
- 7.2 Parallel subtests
- 7.3 Setup and teardown with parallel subtests
- 7.4 Gotchas with closures and parallel tests

###—> 8. Testing race conditions 0'35''
- 8.1 What is a race condition
- 8.2 The race detection flag
- 8.3 Testing explicitly for race conditions

###—> 9. Comparing objects for equality 0'37''
- 9.1 Simple comparisons
- 9.2 Reflect's DeepEqual function
- 9.3 Golden files (brief overview)
- 9.4 Helper comparison functions

###—> 10. Testing utilities 0'39''
- 10.1 Building things with helper functions
- 10.2 Generating test data
- 10.3 Gos quick testing package
- 10.4 Public testing utilities

###—> 11. Controlling which tests are run 0'22''
- 11.1 Running specific tests
- 11.2 Running tests for subpackages
- 11.3 Skipping tests
- 11.4 Custom flags
- 11.5 Build tags

###—> 12. Additional testing flags 0'57''
- 12.1 Benchmarks
- 12.2 Verbose testing
- 12.3 Code coverage
- 12.4 The timeout flag
- 12.5 Parallel testing flags

###—> 13. External and internal testing 0'40''
- 13.1 Differences between external and internal
- 13.2 How to write internal and external tests
- 13.3 When to use external tests
- 13.4 Exporting unexported vars, funcs, and types
- 13.5 When to use internal tests

###—> 14. Types of tests 0'32''
- 14.1 Overview of test types
- 14.2 Unit tests
- 14.3 Integration tests
- 14.4 End-to-end tests
- 14.5 Which test type should I use

###—> 15. State 0'16''
- 15.1 What is global state
- 15.2 Testing with global state (if you must)

###—> 16. Dependency injection (DI) 1'10''
- 16.1 What is dependency injection
- 16.2 DI enables implementation agnostic code
- 16.3 DI makes testing easier
- 16.4 DI and useful zero values
- 16.5 Removing global state with DI
- 16.6 Package level functions
- 16.7 Summary of DI

###—> 17. Mocks, stubs, and fakes 0'50''
- 17.1 What is mocking
- 17.2 Types of mock objects
- 17.3 Why do we mock
- 17.4 Third party packages
- 17.5 Faking APIs

###—> 18. Interface test suites 0'29''
- 18.1 What are interface test suites
- 18.2 Interface test suite setup and teardown
- 18.3 Interface test suites in the wild

###—> 19. Testing with HTTP 0'35''
- 19.1 httptest.ResponseRecorder
- 19.2 httptest.Server
- 19.3 Build HTTP helpers

###—> 20. Golden Files 0'14''
- 20.1 What are golden files
- 20.2 Updating golden files

###—> 21. Testing subprocesses 0'42''
- 21.1 What is a subprocess
- 21.2 Running the subprocess in tests
- 21.3 Mocking simple subprocesses
- 21.4 Mocking complex subprocesses

###—> 22. Testing with time 0'38''
- 22.1 Why are dates and times problematic
- 22.2 Inject your time and sleep functions
- 22.3 Testing timeouts

###—> 23. Bonus material 0'05''
- 23.1 Colorizing your terminal output
- 23.2 Coverage info function
