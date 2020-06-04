package iotmaker_db_mongodb_config

import "strings"

type ComaList []string

func (el *ComaList) String() string {
	return `"` + strings.Join(*el, `","`) + `"`
}
