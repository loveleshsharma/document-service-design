package document

type AccessType int

const (
	Read     AccessType = 0
	Edit     AccessType = 1
	NoAccess AccessType = 2
)
