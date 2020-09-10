package service

type CowService interface {
	Say(moosage string) (string, error)
	Think(moosage string) (string, error)
}
