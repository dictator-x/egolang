package model

import (
	// "fmt"
	"encoding/json"
)

type Profile struct {
	Name string
	// Gender     string
	Age int
	// Height     int
	// Weight     int
	// Income     string
	// Marriage   string
	// Education  string
	// Occupation string
	// Hokou      string
	// Xinzuo     string
	// House      string
	// Car        string
}

// func (p *Profile) String() string {
// 	return fmt.Sprintf("name: %s; age: %d", p.Name, p.Age)
// }

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(s, &profile)
	return profile, err
}
