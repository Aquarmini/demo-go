package providers

type Logger struct {
	id   int64
	name string
}

func (this *Logger)Output() (r int64, err error) {
	this.id++
	r = this.id
	return
}
