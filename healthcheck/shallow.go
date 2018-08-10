package healthcheck

type Shallow struct {
	Msg     string `json:"message"`
	IsAlive bool   `json:"isAlive"`
}
