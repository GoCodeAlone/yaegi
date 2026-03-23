//go:build go1.26

package stdlib

//go:generate ../internal/cmd/extract/extract crypto/hkdf crypto/pbkdf2 crypto/sha3
//go:generate ../internal/cmd/extract/extract iter structs testing/synctest unique weak
