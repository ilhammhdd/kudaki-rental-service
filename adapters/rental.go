package adapters

type EventDrivenAdapter interface {
	Adapt(msg []byte) bool
	Log(partition int32, offset int64, key string)
}

type SubmitRental struct {
	usecase func()
}

func (sr *SubmitRental) Adapt(msg []byte) bool {
	stat := false

	if stat {
		sr.usecase()
	}

	return stat
}

func (sr *SubmitRental) Log(partition int32, offset int64, key string) {

}
