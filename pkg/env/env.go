package env

import (
	"encoding/json"
	"expvar"
	"log"
	"os"
)

var env = expvar.NewMap("env")

// Get returns the value of the given environment variable.
func Get(name string, defaultValue string) string {

	value, ok := os.LookupEnv(name)
	if !ok {
		value = defaultValue
	}

	log.Println("Loaded env var: ", name)

	env.Set(name, jsonStringer(value))

	return value
}

type jsonStringer string

func (s jsonStringer) String() string {
	v, _ := json.Marshal(s)
	return string(v)
}
