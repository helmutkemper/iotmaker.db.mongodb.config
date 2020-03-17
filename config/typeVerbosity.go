package iotmaker_db_mongodb_config

import (
	"strconv"
)

type Verbosity int

func (el Verbosity) String() string {

	return strconv.FormatInt(int64(el), 10)
}
