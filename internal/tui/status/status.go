package status

type Status int

const (
	Ok Status = iota
	Start
	Success
	Err
)
