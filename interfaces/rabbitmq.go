package interfaces

type Queue struct {
	Name      string
	Durable   bool
	Delete    bool
	Exclusive bool
	NoWait    bool
}

type SensorInfo struct {
	Disp        string
	Temperatura float32
	Presion     float32
	Nivel       float32
	Corriente   float32
}

type PLCinfo struct {
	Timestamp int64     `json:"timestamp"`
	Values    []PLCdata `json:"values"`
}

type PLCdata struct {
	ID string `json:"id"`
	V  int    `json:"v"`
	Q  bool   `json:"q"`
	T  int64  `json:"t"`
}

type PLC struct {
	ID           string
	PresionC     int
	PresionCh    int
	TemperaturaC int
}
