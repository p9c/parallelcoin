package wire

import (
	"fmt"
	"io"
	
	chainhash "github.com/p9c/parallelcoin/pkg/chainhash"
)

// MsgGetHeaders implements the Message interface and represents a bitcoin
// getheaders message. It is used to request a list of block headers for blocks
// starting after the last known hash in the slice of block locator hashes. The
// list is returned via a headers message (MsgHeaders) and is limited by a
// specific hash to stop at or the maximum number of block headers per message,
// which is currently 2000. Set the HashStop field to the hash at which to stop
// and use AddBlockLocatorHash to podbuild up the list of block locator hashes. The
// algorithm for building the block locator hashes should be to add the hashes
// in reverse order until you reach the genesis block. In order to keep the list
// of locator hashes to a resonable number of entries, first add the most recent
// 10 block hashes, then double the step each loop iteration to exponentially
// decrease the number of hashes the further away from head and closer to the
// genesis block you get.
type MsgGetHeaders struct {
	ProtocolVersion    uint32
	BlockLocatorHashes []*chainhash.Hash
	HashStop           chainhash.Hash
}

// AddBlockLocatorHash adds a new block locator hash to the message.
func (msg *MsgGetHeaders) AddBlockLocatorHash(hash *chainhash.Hash) (e error) {
	if len(msg.BlockLocatorHashes)+1 > MaxBlockLocatorsPerMsg {
		str := fmt.Sprintf(
			"too many block locator hashes for message [max %v]",
			MaxBlockLocatorsPerMsg,
		)
		return messageError("MsgGetHeaders.AddBlockLocatorHash", str)
	}
	msg.BlockLocatorHashes = append(msg.BlockLocatorHashes, hash)
	return nil
}

// BtcDecode decodes r using the bitcoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgGetHeaders) BtcDecode(r io.Reader, pver uint32, enc MessageEncoding) (e error) {
	if e = readElement(r, &msg.ProtocolVersion); E.Chk(e) {
		return
	}
	// Read num block locator hashes and limit to max.
	var count uint64
	if count, e = ReadVarInt(r, pver); E.Chk(e) {
		return
	}
	if count > MaxBlockLocatorsPerMsg {
		str := fmt.Sprintf(
			"too many block locator hashes for message "+
				"[count %v, max %v]", count, MaxBlockLocatorsPerMsg,
		)
		return messageError("MsgGetHeaders.BtcDecode", str)
	}
	// Create a contiguous slice of hashes to deserialize into in order to reduce
	// the number of allocations.
	locatorHashes := make([]chainhash.Hash, count)
	msg.BlockLocatorHashes = make([]*chainhash.Hash, 0, count)
	for i := uint64(0); i < count; i++ {
		hash := &locatorHashes[i]
		if e = readElement(r, hash); E.Chk(e) {
			return
		}
		if e = msg.AddBlockLocatorHash(hash); E.Chk(e) {
		}
	}
	return readElement(r, &msg.HashStop)
}

// BtcEncode encodes the receiver to w using the bitcoin protocol encoding. This
// is part of the Message interface implementation.
func (msg *MsgGetHeaders) BtcEncode(w io.Writer, pver uint32, enc MessageEncoding) (e error) {
	// Limit to max block locator hashes per message.
	count := len(msg.BlockLocatorHashes)
	if count > MaxBlockLocatorsPerMsg {
		str := fmt.Sprintf(
			"too many block locator hashes for message "+
				"[count %v, max %v]", count, MaxBlockLocatorsPerMsg,
		)
		return messageError("MsgGetHeaders.BtcEncode", str)
	}
	if e = writeElement(w, msg.ProtocolVersion); E.Chk(e) {
		return
	}
	if e = WriteVarInt(w, pver, uint64(count)); E.Chk(e) {
		return
	}
	for _, hash := range msg.BlockLocatorHashes {
		if e = writeElement(w, hash);E.Chk(e){
			return
		}
	}
	return writeElement(w, &msg.HashStop)
}

// Command returns the protocol command string for the message.  This is part of the Message interface implementation.
func (msg *MsgGetHeaders) Command() string {
	return CmdGetHeaders
}

// MaxPayloadLength returns the maximum length the payload can be for the receiver. This is part of the Message
// interface implementation.
func (msg *MsgGetHeaders) MaxPayloadLength(pver uint32) uint32 {
	// Version 4 bytes + num block locator hashes (varInt) + max allowed block locators + hash stop.
	return 4 + MaxVarIntPayload + (MaxBlockLocatorsPerMsg *
		chainhash.HashSize) + chainhash.HashSize
}

// NewMsgGetHeaders returns a new bitcoin getheaders message that conforms to the Message interface. See MsgGetHeaders
// for details.
func NewMsgGetHeaders() *MsgGetHeaders {
	return &MsgGetHeaders{
		BlockLocatorHashes: make(
			[]*chainhash.Hash, 0,
			MaxBlockLocatorsPerMsg,
		),
	}
}
