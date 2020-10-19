# Go-Ruler

> Go-Ruler is an rules engine for go and partially compatible with [K-Phoen/rulerz](https://github.com/K-Phoen/rulerz) which is written in PHP

## Introduction

Business rules can be written as text using a dedicated language, very close to SQL, in which case we refer to them as *rules* or they can be encapsulated in single classes and referred to as *specifications*.

Once a rule (or a specification) is written, it can be used to check if a single candidate satisfies it or directly to query a datasource.

## Quick usage

#### 1. Install

```shell
go get -u github.com/ibllex/go-ruler
```

#### 2. Write a rule

The rule hereafter describes a "*high ranked female player*" (basically, a female player having more than 9000 points).

```go
highRankFemalesRule := "gender = 'F' and points > 9000"
```

#### 3. Define a datasource

We have the following datasources:

```go
import "github.com/ibllex/go-ruler"

players := []ruler.T{
	ruler.T{"pseudo": "Joe", "gender": "M", "points": 40},
	ruler.T{"pseudo": "Moe", "gender": "M", "points": 20},
	ruler.T{"pseudo": "Alice", "gender": "F", "points": 60},
	ruler.T{"pseudo": "Birdie", "gender": "F", "points": 60},
}
```

#### 4. Create ruler

```go
import (
	"github.com/ibllex/go-ruler"
)

r := ruler.New(ruler.O{})
```

#### 5. Use a rule to query a datasource

retrieving the results is as simple as calling the `Filter` method:

```go
import (
	"github.com/ibllex/go-ruler"
)

highRankFemales, err := r.Filter(players, highRankFemalesRule, ruler.P{}, ruler.PP{})
```

Given a candidate, checking if it satisfies a rule boils down to calling the `Satisfies` method:

```go
import (
	"github.com/ibllex/go-ruler"
)

isHighRankFemale, err := r.Satisfies(players[0], highRankFemalesRule, ruler.P{}, ruler.PP{})
```

## License

This library is under the [MIT](https://github.com/ibllex/go-ruler/blob/master/LICENSE) license.