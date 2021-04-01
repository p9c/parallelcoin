package ecc

import (
	"compress/zlib"
	"encoding/base64"
	"encoding/binary"
	"io"
	"io/ioutil"
	"strings"
)

// loadS256BytePoints decompresses and deserializes the pre-computed byte points
// used to accelerate scalar base multiplication for the secp256k1 curve.
//
// This approach is used since it allows the compile to use significantly less
// ram and be performed much faster than it is with hard-coding the final
// in-memory data structure.
//
// At the same time, it is quite fast to generate the in-memory data structure
// at init time with this approach versus computing the table.
func loadS256BytePoints() (e error) {
	// There will be no byte points to load when generating them.
	bp := secp256k1BytePoints
	// if len(bp) == 0 {
	// 	return nil
	// }
	// Decompress the pre-computed table used to accelerate scalar base
	// multiplication.
	decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(bp))
	var r io.ReadCloser
	if r, e = zlib.NewReader(decoder); E.Chk(e) {
		return
	}
	var serialized []byte
	if serialized, e = ioutil.ReadAll(r); E.Chk(e) {
		return
	}
	// Deserialize the precomputed byte points and set the curve to them.
	offset := 0
	var bytePoints [32][256][3]fieldVal
	for byteNum := 0; byteNum < 32; byteNum++ {
		// All points in this window.
		for i := 0; i < 256; i++ {
			px := &bytePoints[byteNum][i][0]
			py := &bytePoints[byteNum][i][1]
			pz := &bytePoints[byteNum][i][2]
			for i := 0; i < 10; i++ {
				px.n[i] = binary.LittleEndian.Uint32(serialized[offset:])
				offset += 4
			}
			for i := 0; i < 10; i++ {
				py.n[i] = binary.LittleEndian.Uint32(serialized[offset:])
				offset += 4
			}
			for i := 0; i < 10; i++ {
				pz.n[i] = binary.LittleEndian.Uint32(serialized[offset:])
				offset += 4
			}
		}
	}
	secp256k1.bytePoints = &bytePoints
	return nil
}
