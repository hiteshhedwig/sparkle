package config

import (
	"encoding/json"
)

func (c *Config) ToMAP() (res map[string]interface{}) {
	a, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(a, &res)
	return
}

func (c *Config) LoadFromMap(m interface{}) error {
	data, err := json.Marshal(m)
	//fmt.Println(data)
	if err == nil {
		err = json.Unmarshal(data, c)
	}
	return err
}

func (c *Config) ToJSON() string {
	a, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(a)
}

func (c *Config) LoadFromJSON(jsonStr string) error {
	err := json.Unmarshal([]byte(jsonStr), c)
	return err
}

func (c *Config) ToByte() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Config) LoadFromByte(data []byte) error {
	return json.Unmarshal(data, c)
}
