package j

import "os"

func init() {
	configDefault[Echo] = true
	configDefault[TimeFormat] = TimeMS
	configDefault[Caller] = CallerShort
	configDefault[PermDir] = os.FileMode(0755)
	configDefault[PermFile] = os.FileMode(0644)
}
