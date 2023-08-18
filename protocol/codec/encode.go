package codec

type Encoder struct{}

func (e Encoder) Encode(msg any) ([]byte, error) {
	return nil, nil
}
