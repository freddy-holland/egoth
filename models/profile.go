package models

import (
	"fmt"
	"time"
)

type Profile struct {
	Name     string
	Nickname string
	Picture  string
	Expiry   time.Time
	Issued   time.Time
}

func (p Profile) IsExpired() bool {
	if time.Now().After(p.Expiry) {
		return true
	} else {
		return false
	}
}

func ProfileFromMap(data map[string]interface{}) (Profile, error) {
	var profile Profile
	var ok bool

	if profile.Name, ok = data["name"].(string); !ok {
		return Profile{}, fmt.Errorf("Missing name parameter")
	}

	if profile.Nickname, ok = data["nickname"].(string); !ok {
		return Profile{}, fmt.Errorf("Missing nickname parameter")
	}

	if profile.Picture, ok = data["picture"].(string); !ok {
		return Profile{}, fmt.Errorf("Missing picture parameter")
	}

	if expFloat, ok := data["exp"].(float64); ok {
		profile.Expiry = time.Unix(int64(expFloat), 0)
	} else {
		return Profile{}, fmt.Errorf("Missing exp parameter")
	}

	if issuedFloat, ok := data["iat"].(float64); ok {
		profile.Issued = time.Unix(int64(issuedFloat), 0)
	} else {
		return Profile{}, fmt.Errorf("Missing iat parameter")
	}

	return profile, nil
}
