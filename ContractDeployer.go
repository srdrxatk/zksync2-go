package zksync2

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/zksync-sdk/zksync2-go/contracts/ContractDeployer"
	"math/big"
	"strings"
)

var contractDeployerABI *abi.ABI

func EncodeCreate2(bytecode, salt []byte) ([]byte, error) {
	if contractDeployerABI == nil {
		cda, err := abi.JSON(strings.NewReader(ContractDeployer.ContractDeployerMetaData.ABI))
		if err != nil {
			return nil, fmt.Errorf("failed to load ContractDeployer ABI: %w", err)
		}
		contractDeployerABI = &cda
	}
	// prepare
	if len(salt) == 0 {
		salt = make([]byte, 32)
	} else if len(salt) != 32 {
		return nil, errors.New("salt must be 32 bytes")
	}
	hash, err := HashBytecode(bytecode)
	if err != nil {
		return nil, fmt.Errorf("failed to get hash of bytecode: %w", err)
	}
	salt32 := [32]byte{}
	copy(salt32[:], salt)
	hash32 := [32]byte{}
	copy(hash32[:], hash)

	res, err := contractDeployerABI.Pack("create2", salt32, hash32, big.NewInt(0), []byte{})
	if err != nil {
		return nil, fmt.Errorf("failed to pack create2 function: %w", err)
	}
	return res, nil
}

func HashBytecode(bytecode []byte) ([]byte, error) {
	bytecodeHash := sha256.Sum256(bytecode)
	// get real length of bytecode, which is presented as 32-byte words
	length := big.NewInt(int64(len(bytecode) / 32))
	if length.BitLen() > 16 {
		return nil, errors.New("bytecode length must be less than 2^16 bytes")
	}
	length2b := make([]byte, 2)
	length2b = length.FillBytes(length2b) // 0-padded in 2 bytes
	// replace first 2 bytes of hash with bytecode length
	return append(length2b, bytecodeHash[2:]...), nil
}