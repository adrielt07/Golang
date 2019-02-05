# Unittest and Code Coverage

### Purpose
The purpose of this code is for me to explore how Unittesting and Code coverage is done in Go programming language.

Go was written with the importance of testing in mind; therefore, testing tools are built-in when you first install Go.

### To Run Test
```GOPATH=<your directory>/Testing_with_Go/ go test pack```

### Why Test???
#### Stability
1. Well Structured tests will increase the stability of an application. If you have a well-structured test, you can be more aggressive to apply new feature to your application. The test will tell you if a regression was introduced.

#### Performance
2. Testing the critical element of an application can ensure that you're aware of degration perfomance as soon as possible. Quickly Knowing that a degration issue occured will make sure the team is aware of the change that caused it while it's still fresh in mind.

#### Documentation
3. As long as testing is run on a daily basis and kept up-to-date. Testing can be used as a documentation as of how the application is built overtime.

#### Design tool
4. Test-driven development can be used to design an application. The idea behind this is for someone to describe how the application should behave before coding it. In other words, you'll create the test first for the new feature of your application before actually implementing the feature to development/production.

### Testing with Go
#### Naming Conventions
1. Test File - add '_test.go' suffix example: filename_test.go
Side Effects:
- Ignored during normal Go builds
- The 'test runner' will inspect the test files for testing. Only files with the correct test naming convention will be executed

#### Three Different Types of Test
1. Basic test - starts with Prefix "Test" example: Testhello
- used to verify the correctness of application
2. Benchmark test - prefix "Benchmark" example: BenchmarkHello
- used to mnake sure the application is performing under the specified limit
3. Example test - prefix "Example" example: ExampleHello
- integrates with Godoc tools