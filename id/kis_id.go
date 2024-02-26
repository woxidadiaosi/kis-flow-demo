package id

import (
	"github.com/google/uuid"
	"kis-flow-demo/common"
	"strings"
)

func KisID(prefix ...string) string {
	idStr := strings.Replace(uuid.New().String(), "-", "", -1)
	kisId := fromKisID(idStr, prefix...)
	return kisId
}

func fromKisID(idStr string, str ...string) string {
	var kisId string
	for _, pre := range str {
		kisId += pre
		kisId += common.KisIdJoinChar
	}
	kisId += idStr
	return kisId
}
