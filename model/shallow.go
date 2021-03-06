package model

type ShallowHealthChecks struct {
	HealthChecks []Shallow `json:"healthchecks"`
}

type Shallow struct {
	InfraName string `json:"name"`
	Msg       string `json:"message"`
	IsAlive   bool   `json:"alive"`
}

func NewShallow(iname string, msg string, isAlive bool) Shallow {
	return Shallow{
		InfraName: iname,
		IsAlive:   isAlive,
		Msg:       msg,
	}
}
