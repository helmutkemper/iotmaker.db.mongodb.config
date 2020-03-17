package iotmaker_db_mongodb_config

import (
	"reflect"
	"strconv"
)

type Verbosity int

func (el Verbosity) String() string {

	strconv.FormatInt(int64(el), 10)
}
