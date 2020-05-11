package validator

import (
	"encoding/json"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
	"io/ioutil"
	"log"
	"os"
)

// ResumeValidator contains the data, validator and populates resume on validation
type ResumeValidator struct {
	data     []byte
	validate *validator.Validate
	resume   Resume
	filepath string
	source   string
}

// WithFile sets the filepath that will be validated
func (rv *ResumeValidator) WithFile(filepath string) *ResumeValidator {
	rv.filepath = filepath
	rv.source = "file"
	return rv
}

// WithData sets the bytes that will be validated
func (rv *ResumeValidator) WithData(data []byte) *ResumeValidator {
	rv.data = data
	rv.source = "raw"
	return rv
}

// Validate attempts to validate json resume either from file
// or data you provided using WithData/WithFile
func (rv *ResumeValidator) Validate() error {
	if rv.source == "file" {
		return rv.validateFromFile()
	}

	return rv.validateFromRaw()
}

func (rv *ResumeValidator) validateFromFile() error {
	log.Printf("reading file %s", rv.filepath)
	data, err := ioutil.ReadFile(rv.filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.Wrap(err, "file does not exist")
		}
		return errors.Wrap(err, "failed to read file")
	}

	rv.data = data
	return rv.validateFromRaw()
}

func (rv *ResumeValidator) validateFromRaw() error {
	if rv.data == nil {
		return errors.New("data/filepath not provided")
	}

	log.Printf("data is ready, unmarshaling to json")
	err := rv.unmarshalData()
	if err != nil {
		return errors.Wrap(err, "json error")
	}

	log.Printf("unmarsheled the data, validating now..")
	rv.validate = validator.New()
	err = rv.validate.Struct(rv.resume)
	return err
}

// IsValid return only a boolean indicating if the
// provided resume is valid
func (rv *ResumeValidator) IsValid() bool {
	err := rv.Validate()
	return err == nil
}

func (rv *ResumeValidator) unmarshalData() error {
	err := json.Unmarshal(rv.data, &rv.resume)
	if err != nil {
		return errors.Wrap(err, "could not unmarshal data")
	}

	return nil
}
