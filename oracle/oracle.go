package oracle

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	//"math/big"
)

type Client interface {
	Config() *Config
}

type Config struct {
	URL     string
	Address common.Address
}

type Broker struct {
	cfg      *Config
	client   *ethclient.Client
	contract *ChainlinkOracle
}

func New(ctx context.Context, cfg *Config) (*Broker, error) {
	if cfg.URL == "" {
		return nil, errors.New("must provide a INFURA URL value")
	}

	bkr := &Broker{
		cfg: cfg,
	}

	client, err := ethclient.Dial(cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("could not initialize client: %v", err)
	}

	instance, err := NewChainlinkOracle(cfg.Address, client)

	if err != nil {
		log.Fatal(err)
	}

	bkr.client = client
	bkr.contract = instance

	return bkr, nil
}

func (bkr *Broker) Config() *Config {
	return bkr.cfg
}

func (bkr *Broker) Client() *ethclient.Client {
	return bkr.client
}

func (bkr *Broker) getGas(user common.Address) (*big.Int, error) {
	gas, err := bkr.contract.GetLatestAnswer(nil)
	if err != nil {
		return nil, err
	}
	return gas, err
}
