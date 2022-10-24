package service

type event interface {
	SendMsg()
	SaveData()
}
