package util

import (
	"os"
	"time"
)

type AuditEvent struct {
	Event string
	Node  string
	Start time.Time
	End   time.Time
}

type Audit []AuditEvent

func ParseAudit(file os.File) {

}
