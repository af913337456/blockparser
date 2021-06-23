package btc_chain

import (
	"github.com/af913337456/blockparser/types"
	"github.com/go-xorm/xorm"
)

/**
 * Copyright (C), 2019-2021
 * FileName: mysql
 * Author:   LinGuanHong
 * Date:     2021/6/23 4:54 下午
 * Description: 使用 MySQL 进行数据存储的例子，非常简洁的代码
 */

type BtcBlockScannerMysqlStorage struct {
	db      *xorm.Engine
	session *xorm.Session
}

func NewBtcBlockScannerMysqlStorage(db *xorm.Engine) *BtcBlockScannerMysqlStorage {
	return &BtcBlockScannerMysqlStorage{
		db: db,
	}
}

func (b *BtcBlockScannerMysqlStorage) GetDbLastBlock() (*types.Block, error) {
	latestBlock := &types.Block{}
	if _, err := b.db.SQL("select from block order by id desc limit 1").Get(latestBlock); err != nil {
		return nil, err
	}
	return latestBlock, nil
}
func (b *BtcBlockScannerMysqlStorage) GetDbBlockByHash(blockHash string) (*types.Block, error) {
	targetBlock := &types.Block{}
	if _, err := b.db.SQL("select from block where block_hash=?", blockHash).Get(targetBlock); err != nil {
		return nil, err
	}
	return targetBlock, nil
}
func (b *BtcBlockScannerMysqlStorage) RecordBlock(block *types.Block, updateModel, commitAfterOpt bool) error {
	_, err := b.db.Insert(block)
	return err
}
func (b *BtcBlockScannerMysqlStorage) HandleForkEvent(info *types.BlockForkInfo) error {
	// todo according your own logic to judge then return
	return nil
}
func (b *BtcBlockScannerMysqlStorage) TransactionHandler(block *types.ScannerBlockInfo, dbTx interface{}, blockTxs interface{}) error {
	// do something, such as push the data to kafka queue
	return nil
}
func (b *BtcBlockScannerMysqlStorage) TxOpen() (interface{}, error) {
	b.session = b.db.NewSession()
	return nil, nil
}
func (b *BtcBlockScannerMysqlStorage) TxCommit() error {
	return b.session.Commit()
}
func (b *BtcBlockScannerMysqlStorage) TxRollBack() error {
	return b.session.Rollback()
}
func (b *BtcBlockScannerMysqlStorage) TxClose() {
	b.session.Close()
}
func (b *BtcBlockScannerMysqlStorage) Close() error {
	b.TxClose()
	return b.db.Close()
}
