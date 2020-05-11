[![Go Report Card](https://goreportcard.com/badge/github.com/cinarmert/json-resume-validator)](https://goreportcard.com/report/github.com/cinarmert/json-resume-validator)
[![Build Status](https://travis-ci.com/cinarmert/json-resume-validator.svg?branch=master)](https://travis-ci.com/cinarmert/json-resume-validator)
[![codecov](https://codecov.io/gh/cinarmert/json-resume-validator/branch/master/graph/badge.svg)](https://codecov.io/gh/cinarmert/json-resume-validator)

# JSON Resume Validator

Go library to validate your JSON Resume according to [defined schema](https://jsonresume.org/schema/)

> For more information on JSON resumes, please visit [official JSON Resume](https://jsonresume.org)


## Installation

```
$ go get github.com/cinarmert/json-resume-validator
```

## Sample Application

- [A sample Go Web Server](https://github.com/cinarmert/json-resume-validator-client)

## Sample Usage

```go

// Initialize with a file
rv := new(resumeValidator.ResumeValidator).WithFile("path/to/file")

// Initialize with data
rv := new(resumeValidator.ResumeValidator).WithData([]bytes("{}"))

// Validate and process errors
err := rv.Validate()

// Check if it is valid 
rv.IsValid()

```

More samples can be found in tests.
