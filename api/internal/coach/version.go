package coach

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
)

// packVersionPrefix starts every pack version; packVersionHexChars is how much of the
// hash it keeps.
const (
	packVersionPrefix   = "pack-v1."
	packVersionHexChars = 16
)

// PackVersion is "pack-v1." + a hash of RulesVersion, the spec rule file's bytes and the
// catalogue's bytes (§14): a change to any of them moves it. Each part is length-prefixed
// so bytes moving between parts move the version too.
func PackVersion(specRules, catalogue []byte) string {
	h := sha256.New()
	for _, part := range [][]byte{[]byte(RulesVersion), specRules, catalogue} {
		var n [8]byte
		binary.BigEndian.PutUint64(n[:], uint64(len(part)))
		h.Write(n[:])
		h.Write(part)
	}
	return packVersionPrefix + hex.EncodeToString(h.Sum(nil))[:packVersionHexChars]
}
