package document

import (
	"crypto/sha256"
	"encoding/hex"
)

func computeID(doc *Document) string {
	sum := sha256.Sum256([]byte(doc.name + "\n" + doc.content))
	return hex.EncodeToString(sum[:])
}
