package payment

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Verifier struct {
	client *ethclient.Client
	to     common.Address
	price  *big.Int
	store  *Store
}

func NewVerifier(rpc, to string, price int64, store *Store) (*Verifier, error) {
	c, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}
	return &Verifier{
		client: c,
		to:     common.HexToAddress(to),
		price:  big.NewInt(price),
		store:  store,
	}, nil
}

func (v *Verifier) Verify(txHash string) error {
	if v.store.Used(txHash) {
		return errors.New("tx already used")
	}

	tx, _, err := v.client.TransactionByHash(
		context.Background(),
		common.HexToHash(txHash),
	)
	if err != nil {
		return err
	}

	if tx.To() == nil || *tx.To() != v.to {
		return errors.New("wrong receiver")
	}

	if tx.Value().Cmp(v.price) < 0 {
		return errors.New("underpayment")
	}

	v.store.Mark(txHash)
	return nil
}
