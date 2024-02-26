package common

const (
	KisIdTypeFlow      = "flow"
	KisIdTypeFunction  = "function"
	KisIdTypeConnector = "conn"
	KisIdTypeGlobal    = "global"
	KisIdJoinChar      = "-"
)

type KisType string

const (
	KisFunction   KisType = "func"
	KisFlow       KisType = "flow"
	KisConnection KisType = "conn"
)

type FMode string

const (
	V FMode = "verify"
	S FMode = "Save"
	L FMode = "Load"
	C FMode = "Calculate"
	E FMode = "Expand"
)

type KisOnOff int

const (
	OFF KisOnOff = iota
	ON
)

type KisConnType string

const (
	Mysql KisConnType = "mysql"
	Redis KisConnType = "redis"
	Kafka KisConnType = "kafka"
)
