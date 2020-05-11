package validator

import (
	"strings"
	"time"
)

const timeLayout = "2006-01-02"

type jsonTime struct {
	time.Time
}

func (ts *jsonTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")

	t, err := time.Parse(timeLayout, s)
	if err != nil {
		return err
	}

	ts.Time = t
	return nil
}

type interval struct {
	StartDate jsonTime `json:"startDate" validate:"ltfield=EndDate"`
	EndDate   jsonTime `json:"endDate" validate:"gtfield=StartDate"`
}

// Resume holds the information according to schema
// described in https://jsonresume.org/schema/
type Resume struct {
	Basics       basics        `json:"basics"`
	Work         []work        `json:"work"`
	Volunteer    []volunteer   `json:"volunteer"`
	Education    []education   `json:"education"`
	Awards       []award       `json:"awards"`
	Publications []publication `json:"publications"`
	Skills       []skill       `json:"skills"`
	Languages    []language    `json:"languages"`
	Interests    []interest    `json:"interests"`
	References   []reference   `json:"references"`
}

type location struct {
	Address     string `json:"address"`
	PostalCode  string `json:"postalCode"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
}

type profile struct {
	Network  string `json:"network"`
	Username string `json:"username"`
	URL      string `json:"url" validate:"url"`
}

type basics struct {
	Name     string    `json:"name" validate:"required"`
	Label    string    `json:"label" validate:"required"`
	Picture  string    `json:"picture"`
	Email    string    `json:"email" validate:"required,email"`
	Phone    string    `json:"phone"`
	Website  string    `json:"website" validate:"url"`
	Summary  string    `json:"summary"`
	Location location  `json:"location"`
	Profiles []profile `json:"profiles"`
}

type work struct {
	interval
	Company    string   `json:"company"`
	Position   string   `json:"position"`
	Website    string   `json:"website" validate:"url"`
	Summary    string   `json:"summary"`
	Highlights []string `json:"highlights"`
}

type volunteer struct {
	interval
	Organization string   `json:"organization"`
	Position     string   `json:"position"`
	Website      string   `json:"website" validate:"url"`
	Summary      string   `json:"summary"`
	Highlights   []string `json:"highlights"`
}

type education struct {
	interval
	Institution string   `json:"institution"`
	Area        string   `json:"area"`
	StudyType   string   `json:"studyType"`
	Gpa         string   `json:"gpa"`
	Courses     []string `json:"courses"`
}

type award struct {
	Title   string   `json:"title"`
	Date    jsonTime `json:"date" validate:"lte"`
	Awarder string   `json:"awarder"`
	Summary string   `json:"summary"`
}

type publication struct {
	Name        string   `json:"name"`
	Publisher   string   `json:"publisher"`
	ReleaseDate jsonTime `json:"releaseDate" validate:"lte"`
	Website     string   `json:"website" validate:"url"`
	Summary     string   `json:"summary"`
}

type skill struct {
	Name     string   `json:"name"`
	Level    string   `json:"level"`
	Keywords []string `json:"keywords"`
}

type language struct {
	Language string `json:"language"`
	Fluency  string `json:"fluency"`
}

type interest struct {
	Name     string   `json:"name"`
	Keywords []string `json:"keywords"`
}

type reference struct {
	Name      string `json:"name"`
	Reference string `json:"reference"`
}
