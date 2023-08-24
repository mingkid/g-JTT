package msg

type M0102 struct {
	Head
	M0102Body
}

type M0102Body struct {
	Token string
}
