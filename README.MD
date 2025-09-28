# ESEP Grade Calculator:

This library is a simple utility to calculate your grades.

## Grade Calculator Rules

- **Weights**
  - Assignments = 50%
  - Exams = 35%
  - Essays = 15%
- **Letter Grade Scale**
  - A ≥ 90  
  - B ≥ 80  
  - C ≥ 70  
  - D ≥ 60  
  - F < 60  

## Clone the project:
```
git clone https://github.com/YOUR_USERNAME/esep-grade-calculator.git
cd esep-grade-calculator
```

Runtime requirements:
Golang 1.24+ 
(run "brew install go" in case you don't have go in your local, check go version with cammand "go version")

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```

View coverage in browser:
```
go tool cover -html=coverage.out -o coverage.html
open coverage.html
```
