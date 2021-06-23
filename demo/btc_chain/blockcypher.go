package btc_chain

import (
	"github.com/af913337456/blockparser/types"
	"math/big"
)

/**
 * Copyright (C), 2019-2021
 * FileName: blockcypher
 * Author:   LinGuanHong
 * Date:     2021/6/23 4:54 下午
 * Description:
 */

type BlockCypherAPI struct {
}

func NewBlockCypherAPI() *BlockCypherAPI {
	return &BlockCypherAPI{}
}

func (b *BlockCypherAPI) GetParentHash(childHash string) (string, error) {
	// todo call API: getBlockByHash
	return "", nil
}
func (b *BlockCypherAPI) GetLatestBlockNumber() (*big.Int, error) {
	// todo call API: getLatestBlock
	return new(big.Int).SetInt64(10000), nil
}
func (b *BlockCypherAPI) GetBlockInfoByNumber(blockNumber *big.Int) (*types.ScannerBlockInfo, error) {
	// todo call API: getBlockByHeight
	return &types.ScannerBlockInfo{}, nil
}
func (b *BlockCypherAPI) Close() {

}
