package withfunctions

func Simple() {}

func WithParams(a int, b string) {}

func WithReturn(a int) (string, error) {
	return "", nil
}

func (r *Receiver) Method() {}
