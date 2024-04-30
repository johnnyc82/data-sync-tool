package synctool

import (
	"errors"
)

// custom errors to throw for check failures
var ErrSyncDirectionMissing = errors.New("sync direction not specified, use one of: --pull --push")
var ErrEnvironmentMissing = errors.New("environment not specified, use one of: --live --staging --local")
var ErrMoreThanOneEnvironment = errors.New("more than one environment selected, select only one of: --live --staging --local")

// consolidate command flags into single struct
// for ease of management
type SyncFlags struct {
	ProjectName string
	FromPath    string
	ToPath      string
	Push        bool
	Pull        bool
	Live        bool
	Staging     bool
	Local       bool
	Data        bool
	Assets      bool
	Media       bool
	Masterplan  bool
	All         bool
}

// Method for flags to identify any problems and to
// set any flags based on logical interdependencies
func (sf *SyncFlags) Check() error {

	// check sync direction specified
	if !sf.Push && !sf.Pull {
		return ErrSyncDirectionMissing
	}

	// check for target environment (live, staging, local)
	if !sf.Live && !sf.Staging && !sf.Local {
		return ErrEnvironmentMissing
	}
	// also raise error if more than one is set
	envFlags := []bool{sf.Live, sf.Staging, sf.Local}
	boolval := 0
	for _, v := range envFlags {
		if v {
			boolval++
		}
		if boolval > 1 {
			return ErrMoreThanOneEnvironment
		}
	}

	// handle --all flag; if set then data, assets, masterplan & media
	// should all be setto true
	if sf.All {
		sf.Data, sf.Assets, sf.Media, sf.Masterplan = true, true, true, true
	}

	return nil
}
