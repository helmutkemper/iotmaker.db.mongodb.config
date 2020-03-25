package iotmaker_db_mongodb_config

type FreeState int

var FreeStates = [...]string{
	"",
	"runtime.",
	"reopen",
}

func (el FreeState) String() string {
	return FreeStates[el]
}

const (
	//Default. You can enable or disable free monitoring during runtime.
	//
	//To enable or disable free monitoring during runtime, see db.enableFreeMonitoring()
	//and db.disableFreeMonitoring().
	//
	//To enable or disable free monitoring during runtime when running with access control,
	//users must have required privileges. See db.enableFreeMonitoring() and
	//db.disableFreeMonitoring() for details.
	KFreeStateRuntime FreeState = iota + 1

	//Enables free monitoring at startup; i.e. registers for free monitoring.
	//When enabled at startup, you cannot disable free monitoring during runtime.
	KFreeStateOn

	//Disables free monitoring at startup, regardless of whether you have previously
	//registered for free monitoring.
	//When disabled at startup, you cannot enable free monitoring during runtime.
	KFreeStateOff
)
