package lib

import "flag"

type IdentityParserModule struct {
	isSet bool
}

func (pm IdentityParserModule) getName() string {
	return "identity"
}

func (pm IdentityParserModule) setParserKeyword() {
	flag.BoolVar(&pm.isSet, "token", false, "identity commands")
}
