package types

import "fmt"

type State int


const (
	STOPPED State = iota
	STARTING
	STARTED
	RUNNING
	EXITED
	FAILED
	UNKNOWN
)

// IsAlive returns true if the OS process representeted by this state exists
func (s State) IsAlive() bool {
	switch s {
	case STARTING, STARTED, RUNNING:
		return true
	case STOPPED, EXITED, FAILED:
		return false
	default:
		return false
	}
}

func (s State) String() string {
	switch s {
	case STOPPED:
		return "STOPPED"
	case STARTING:
		return "STARTING"
	case STARTED:
		return "STARTED"
	case RUNNING:
		return "RUNNING"
	case EXITED:
		return "EXITED"
	case FAILED:
		return "FAILED"
	case UNKNOWN:
		return "UNKNOWN"
	default:
		return "unknown"
	}
}

// MarshalText implements encoding.TextMarshaler interface
func (s State) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface
func (s *State) UnmarshalText(text []byte) error {
	switch string(text) {
	case "STOPPED":
		*s = STOPPED
	case "STARTING":
		*s = STARTING
	case "STARTED":
		*s = STARTED
	case "RUNNING":
		*s = RUNNING
	case "EXITED":
		*s = EXITED
	case "FAILED":
		*s = FAILED
	default:
		return fmt.Errorf("invalid state '%s'", text)
	}
	return nil
}
