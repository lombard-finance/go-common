package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	uuidLength = 16
)

func Uint64FromUUID(id uuid.UUID) uint64 {
	return binary.BigEndian.Uint64(id[:8])
}

// BackupUUID recovers a half of UUID. Reverse operation of `Uint64FromUUID`
func BackupUUID(val uint64) string {
	// UUID is 16 bytes array
	b := make([]byte, uuidLength)
	binary.BigEndian.PutUint64(b[:uuidLength/2], val)
	// `uuid.FromBytes` returns error when invalid length, it's impossible
	id, _ := uuid.FromBytes(b)
	str := id.String()
	// recovered only first half of uuid
	return str[:len(str)/2] + "..."
}

func BinaryFromStringUUID(u string) []byte {
	bytes, _ := uuid.MustParse(u).MarshalBinary()
	return bytes
}

func TryParseUUID(u string) uuid.UUID {
	res, _ := uuid.Parse(u)
	return res
}

func DataToUuid(data []byte) (uuid.UUID, error) {
	hash := sha256.Sum256(data)
	hash[6] = (hash[6] & 0x0f) | 0x40 // version is 4
	hash[8] = (hash[8] & 0x3f) | 0x80 // variant is 10
	return uuid.FromBytes(hash[:16])  // uuid.FromBytes takes exactly 16 bytes
}

func UUIDFromEventData(transactionHash []byte, eventIndex uint32, chainId []byte) uuid.UUID {
	data := make([]byte, 96)
	copy(data[:32], transactionHash)
	copy(data[32:64], common.Hex2BytesFixed(hex.EncodeToString(new(big.Int).SetUint64(uint64(eventIndex)).Bytes()), 32))
	// NOTE: transaction hash + event index is not unique in case of different chains, if add chain id it's make set of values unique between chains
	copy(data[64:96], chainId)

	hash := md5.Sum(data)
	hash[6] = (hash[6] & 0x0f) | 0x40 // version is 4
	hash[8] = (hash[8] & 0x3f) | 0x80 // variant is 10

	id, err := DataToUuid(data)
	if err != nil {
		logrus.WithError(err).Panicf("failed to create UUIDv4 from for event (%x:%d chainid=%x)", transactionHash, eventIndex, chainId)
	}
	return id
}

// UUIDFromBitcoinAddress Compute UUID by hashing bitcoin address and converting first 16 bytes to a UUID
func UUIDFromBitcoinAddress(scriptAddress []byte) uuid.UUID {
	id, err := DataToUuid(scriptAddress)
	if err != nil {
		logrus.WithError(err).Panicf("failed to create UUIDv4 from for script address %x", scriptAddress)
	}
	return id
}
