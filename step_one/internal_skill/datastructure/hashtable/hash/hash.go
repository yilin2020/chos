package hash

func BkdrHash(key []byte) uint {
	seed := uint(131)
	hash := uint(0)
	for _, v := range key {
		if v == '\n' || v == '0' {
			continue
		}
		hash = hash*seed + uint(v)
	}
	return hash & uint(0x7FFFFFFF)
}
