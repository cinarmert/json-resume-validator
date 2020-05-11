package validator

import (
	"testing"
)

const jsonStream = `
	{"basics": {
    "name": "John Doe",
    "label": "Programmer",
    "email": "john@gmail.com",
    "phone": "(912) 555-4321",
    "website": "http://johndoe.com",
    "summary": "A summary of John Doe...",
    "profiles": [{
      "network": "Twitter",
      "username": "john"
    }]
  }}
`

func TestResumeValidator_ValidationFromFile_NoData(t *testing.T) {
	rv := ResumeValidator{}
	err := rv.Validate()

	if err == nil {
		t.Fatalf("expected error but got nil")
	}
}

func TestResumeValidator_ValidationFromFile_Successful(t *testing.T) {
	rv := ResumeValidator{}
	err := rv.WithFile("samples/schema.json").Validate()

	if err != nil {
		t.Fatalf("unexpected err: %+v\n", err)
	}
}

func TestResumeValidator_ValidationFromFile_Unsuccessful(t *testing.T) {
	rv := ResumeValidator{}
	err := rv.WithFile("samples/schema1.json").Validate()

	if err == nil {
		t.Fatalf("expected error but got nil")
	}
}

func TestResumeValidator_ValidationFromRawData_Successful(t *testing.T) {
	rv := ResumeValidator{}
	err := rv.WithData([]byte(jsonStream)).Validate()

	if err != nil {
		t.Fatalf("unexpected err: %+v\n", err)
	}
}

func TestResumeValidator_ValidationFromRawData_Unsuccessful(t *testing.T) {
	rv := ResumeValidator{}
	err := rv.WithData([]byte(jsonStream + "}")).Validate()

	if err == nil {
		t.Fatalf("expected error but got nil")
	}
}

func TestResumeValidator_IsValid_Unsuccessful(t *testing.T) {
	rv := ResumeValidator{}
	got := rv.WithData([]byte(jsonStream + "}")).IsValid()

	if got == true {
		t.Fatalf("expected false, but got true")
	}
}

func TestResumeValidator_IsValid_Successful(t *testing.T) {
	rv := ResumeValidator{}
	got := rv.WithData([]byte(jsonStream)).IsValid()

	if got == false {
		t.Fatalf("expected true, but got false")
	}
}

func TestResumeValidator_unmarshalData(t *testing.T) {
	type fields struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Successful Marshaling",
			fields: fields{
				data: []byte(jsonStream),
			},
			wantErr: false,
		},
		{
			name: "Unsuccessful Marshaling",
			fields: fields{
				data: []byte(jsonStream + "}"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rv := &ResumeValidator{
				data: tt.fields.data,
			}
			if err := rv.unmarshalData(); (err != nil) != tt.wantErr {
				t.Errorf("unmarshalData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
