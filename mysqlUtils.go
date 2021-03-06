// Package mysqlgoutils is a set a functions needed to extent the way skeema works

package mysqlgoutils

import (
    "github.com/skeema/tengo"
	"errors"
	"strings"
)

// SplitHostOptionalPortAndSchem takes an address string containing a hostname,
// ipv4 addr, or ipv6 addr; *optionally* followed by a colon and port number
// optionally followed by a pipe '|' and a schema name. It  splits the hostname
// portion from the port portion and the schema portion and returns them
// separately. If no port was present, 0 will be returned for that portion.
// If no schema was present, '' will be returned for that portion.
// If hostaddr contains an ipv6 address, the IP address portion must be
// wrapped in brackets on input, and the brackets will still be present on
// output.
func SplitHostOptionalPortAndSchema(hostaddr string) (string, int, string, error) {
	if len(hostaddr) == 0 {
		return "", 0, "", errors.New("Cannot parse blank host address")
	}

	// ipv6 without schema or ipv4 or hostname without schema
	if (len(strings.Split(hostaddr, "|")) == 1 ) {
        host, port, err := tengo.SplitHostOptionalPort(hostaddr)
        if err != nil {
            return "", 0, "", err
        }
		return host, port, "", nil
	}

    if strings.Count(hostaddr,"|") > 1 {
        return "", 0, "", errors.New("Too many schema separators")
    }

	var schema string
	schema = strings.Split(hostaddr, "|")[1]
    if len(schema) == 0 {
        return "", 0 ,"", errors.New("cannot parse schema")
    }

	host, port, err := tengo.SplitHostOptionalPort(strings.Split(hostaddr, "|")[0])
	if err != nil {
		return "", 0, "", err
	}

	return host, port, schema, nil 
}
