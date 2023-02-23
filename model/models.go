package model

type MessageStruct struct {
	PostType    string
	SelfQQ      string
	Sender      string
	SenderQQ    string
	Message     string
	MessageType string
	GroupID     string
	GroupName   string
}

type PreConfig struct {
	Ports
	Account
}

type Ports struct {
	Bridge string
	Bot    string
}
type Account struct {
	ID       string
	Password string
}
