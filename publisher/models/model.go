package models

type NatsConnection struct {
	Ip, Port, Username, Password, Subject string
}

type Request struct {
	Command string      `json:"command"`
	Param   interface{} `json:"param"`
	Data    interface{} `json:"data"`
}
