package types

import (
	"io"
)

//------------------------------
type HttpResponse struct {
	Json  string
	Error error
}

type HttpMsg struct {
	Response chan HttpResponse
}

func (h *HttpMsg) SendResponse(json string) {
	if h.Response != nil {
		h.Response <- HttpResponse{Json: json}
	}
}

func (h *HttpMsg) SendError(err error) {
	if h.Response != nil {
		h.Response <- HttpResponse{Error: err}
	}
}

type ShowInfo struct {
	HttpMsg
}

type ShowProcess struct {
	HttpMsg
	Group string
	Name  string
}

type ListProcesses struct {
	HttpMsg
	FailedOnly bool
}

type ReadinessProbe struct {
	HttpMsg
}

type DeleteProcess struct {
	HttpMsg
	Group string
	Name  string
}

type CreateProcess struct {
	HttpMsg
	Group string
	Name  string
	Json  io.ReadCloser
}

type UpdateProcess struct {
	HttpMsg
	Group string
	Name  string
	Json  io.ReadCloser
}

type ProcessStateUpdate struct {
	HttpMsg
	Group string
	Name  string
	State string
}

type StartProcess struct {
	HttpMsg
	Group string
	Name  string
}

type StopProcess struct {
	HttpMsg
	Group string
	Name  string
}

type RestartProcess struct {
	HttpMsg
	Group string
	Name  string
}

type InitdStateUpdate struct {
	HttpMsg
	Name  string
	State string
}

type StartGroup struct {
	HttpMsg
	Name string
}

type StopGroup struct {
	HttpMsg
	Name string
}

type RestartGroup struct {
	HttpMsg
	Name string
}

//------------------------------
type FileCreated struct {
	Path    string
	Process string
}
