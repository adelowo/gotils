package hasher

type Hasher interface {
	Hash(p string) (string, error)
	Verify(hashed, plain string) bool
}
