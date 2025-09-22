package process

import (
	"fmt"
)

func SendMessage(conn string, message string) string {

	return "Message Sent" + fmt.Sprintf(" %s->%s", message, conn)
}
