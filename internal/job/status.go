package job

type Status int

const (
	Ready Status = iota
	Reserved
	Delayed
	Buried
)

func stringToStatus(value string) Status {
	switch value {
	case "reserved":
		return Reserved
	case "delayed":
		return Delayed
	case "buried":
		return Buried
	default:
		return Ready
	}
}

func (s Status) String() string {
	switch s {
	case Reserved:
		return "reserved"
	case Delayed:
		return "delayed"
	case Buried:
		return "buried"
	default:
		return "ready"
	}
}