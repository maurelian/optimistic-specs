package l2

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum-optimism/optimistic-specs/opnode/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func ParseL1InfoDepositTxData(data []byte) (nr uint64, time uint64, baseFee *big.Int, blockHash common.Hash, err error) {
	if len(data) != 4+8+8+32+32 {
		err = fmt.Errorf("data is unexpected length: %d", len(data))
		return
	}
	offset := 4
	nr = binary.BigEndian.Uint64(data[offset : offset+8])
	offset += 8
	time = binary.BigEndian.Uint64(data[offset : offset+8])
	offset += 8
	baseFee = new(big.Int).SetBytes(data[offset : offset+32])
	offset += 32
	blockHash.SetBytes(data[offset : offset+32])
	return
}

func ParseL2Block(refL2Block *types.Block, genesis *Genesis) (refL1 eth.BlockID, refL2 eth.BlockID, parentL2 common.Hash, err error) {
	refL2 = eth.BlockID{Hash: refL2Block.Hash(), Number: refL2Block.NumberU64()}
	if refL2.Number <= genesis.L2.Number {
		if refL2.Hash != genesis.L2.Hash {
			err = fmt.Errorf("unexpected L2 genesis block: %s, expected %s", refL2, genesis.L2)
			return
		}
		refL1 = genesis.L1
		refL2 = genesis.L2
		parentL2 = common.Hash{}
		return
	}

	parentL2 = refL2Block.ParentHash()

	if refL2Block.NumberU64() == 0 {
		// TODO: set to genesis L1 block ID (after sanity checking we got the right L2 genesis block from the engine)
		refL1 = eth.BlockID{}
		return
	}
	txs := refL2Block.Transactions()
	if len(txs) == 0 || txs[0].Type() != types.DepositTxType {
		err = fmt.Errorf("l2 block is missing L1 info deposit tx, block hash: %s", refL2Block.Hash())
		return
	}
	refL1Nr, _, _, refL1Hash, err := ParseL1InfoDepositTxData(txs[0].Data())
	if err != nil {
		err = fmt.Errorf("failed to parse L1 info deposit tx from L2 block: %v", err)
		return
	}
	refL1 = eth.BlockID{Hash: refL1Hash, Number: refL1Nr}
	return
}
