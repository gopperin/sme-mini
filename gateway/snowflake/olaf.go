package snowflake

import (
	"github.com/btnguyen2k/consu/olaf"
)

// Olaf Olaf
var Olaf OlafSnowflake

// OlafSnowflake OlafSnowflake
type OlafSnowflake struct {
	flake *olaf.Olaf
}

// Init Init
func (osf *OlafSnowflake) Init(nodeID int64, epoch int64) error {
	osf.flake = olaf.NewOlafWithEpoch(nodeID, epoch)
	return nil
}

// ID64 ID64
func (osf *OlafSnowflake) ID64() int64 {
	return int64(osf.flake.Id64())
}

// ID64Hex ID64Hex
func (osf *OlafSnowflake) ID64Hex() string {
	return osf.flake.Id64Hex()
}

// ID64Ascii ID64Ascii
func (osf *OlafSnowflake) ID64Ascii() string {
	return osf.flake.Id64Ascii()
}
