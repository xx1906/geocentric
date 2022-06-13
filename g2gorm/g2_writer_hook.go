package g2gorm

type Hook interface {
	Fire(e Entry) (err error)
}
