# PLU

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ferhatelmas/plu)
[![Build Status](https://travis-ci.org/ferhatelmas/plu.svg?branch=master)](https://travis-ci.org/ferhatelmas/plu)

[Price look-up codes](http://en.wikipedia.org/wiki/Price_look-up_code) made easy

:watermelon: 4032
:banana: 4011
:grapes: 4023

[Download PLU codes](https://raw.github.com/ankane/plu/master/plu_codes.csv) - data cleaned up from the [Produce Marketing Association](http://www.plucodes.com)

## How To Use

Initialize (only first call reads csv file)

```go
import "github.com/ferhatelmas/plu"

codes, err := plu.New()
```

List known PLUs

```go
codes.All()
```

Get name from PLU

```go
codes.Name(4011) // Bananas
```

Check if valid

```go
codes.Valid(2000) // false
```

### 5-Digit PLUs

For PLUs with 5 digits, the first digit has a special meaning: 9 specifies organic, and 8 specifies genetically modified.

4011 - Bananas :banana:
94011 - Organic bananas :banana:
84011 - Genetically modified bananas

```go
codes.Organic(94011) // true
codes.GM(84011)      // true
```

### Retailer Assigned

```go
codes.RetailerAssigned(3170) // true
```

## Installation

By go tool

```go
go get github.com/ferhatelmas/plu
```

## TODO

- clean up data

## Resources

- [IFPS 2012 Users Guide](http://www.plucodes.com/docs/Users_Guide_July_2012_FINAL.pdf)
- [Ruby version for the initial inspiration](https://github.com/ankane/plu)

## Contributing

Everyone is encouraged to help improve this project. Here are a few ways you can help:

- [Report bugs](https://github.com/ferhatelmas/plu/issues)
- Fix bugs and [submit pull requests](https://github.com/ferhatelmas/plu/pulls)
- Write, clarify, or fix documentation
- Suggest or add new features
- Help clean up [data](https://github.com/ferhatelmas/plu/blob/master/plu_codes.csv)
