package blockchain

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	epassport "github.com/web3vote/Vote2024/go/EPassport"

	VoteGo "github.com/web3vote/Vote2024/go/Vote"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	EPassport *epassport.EPassportSession
	Vote *VoteGo.VoteSession
}

type Config struct {
	PrivateKey       string `env:"PRIVATE_KEY,notEmpty"`
	Gateway          string `env:"GATEWAY,notEmpty"`
	AccountAddress   string `env:"ACCOUNT_ADDRESS,notEmpty"`
	//PassportAddress  string `env:"PASSPORT_ADDRESS,notEmpty"`
	SingletonAddress string `env:"SINGLETON_ADDRESS,notEmpty"`
	FactoryAddress   string `env:"FACTORY_ADDRESS,notEmpty"`
	EPassportAddress string `env:"EPASSPORT_ADDRESS,notEmpty"`
	VoteAddress 	 string `env:"VOTE_ADDRESS,notEmpty"`
}

func NewClient(config Config) (*Client, error) {
	// Connecting to blockchain network
	client, err := ethclient.Dial(config.Gateway) // load from local .env file
	if err != nil {
		return nil, fmt.Errorf("could not connect to Ethereum gateway: %v\n", err)
	}

	// setting up private key in proper format
	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		return nil, err
	}

	// Creating an auth transactor
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(137))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(3))
	accountAddress := common.HexToAddress(config.AccountAddress)
	balance, err := client.BalanceAt(ctx, accountAddress, nil) //our balance
	cancel()
	if err != nil {
		return nil, err
	}

	log.Printf("Balance of the validator bot: %d\n", balance)


	/*
	// Setting up Passport Contract    OBSOLETE
	passportCenter, err := passport.NewPassport(common.HexToAddress(config.PassportAddress), client)
	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate a TGPassport contract: %v", err)
	}
	*/



	// new epassport contract
	epassportCenter, err := epassport.NewEPassport(common.HexToAddress(config.EPassportAddress), client)
	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate a EPassport contract: %v", err)
	}


	// new vote contract
	voteCenter, err := VoteGo.NewVote(common.HexToAddress(config.VoteAddress), client)
		if err != nil {
			return nil, fmt.Errorf("Failed to instantiate a Vote contract: %v", err)
		}


		// Wrap the Passport contract instance into a session
		epassport := &epassport.EPassportSession{
			Contract: epassportCenter,
			CallOpts: bind.CallOpts{
				Pending: true,
				From:    auth.From,
				Context: context.Background(),
			},
			TransactOpts: bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasLimit: 0,   // 0 automatically estimates gas limit
				GasPrice: nil, // nil automatically suggests gas price
				Context:  context.Background(),
			},
		}


		// Wrap the Passport contract instance into a session
		vote := &VoteGo.VoteSession{
					Contract: voteCenter,
					CallOpts: bind.CallOpts{
						Pending: true,
						From:    auth.From,
						Context: context.Background(),
					},
					TransactOpts: bind.TransactOpts{
						From:     auth.From,
						Signer:   auth.Signer,
						GasLimit: 0,   // 0 automatically estimates gas limit
						GasPrice: nil, // nil automatically suggests gas price
						Context:  context.Background(),
					},
				}

	return &Client{
		//Passport:  passport,
		EPassport: epassport,
		Vote: vote,

	}, nil
}



func (client *Client) GetHash(input string) ([32]byte, error){
	output,err := client.EPassport.GetKeccakHash(input)
	if err != nil {
		var emptyHash [32]byte
		return emptyHash, err
	} else {
		return output,nil
	}
}



/*
func (client *Client) IsRegistered(user_id int64) bool {
	passport_address, err := client.Passport.GetPassportWalletByID(user_id)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("check that user with this id:", user_id)
	log.Println("have associated wallet address:", passport_address)

	if passport_address == common.HexToAddress("0x0000000000000000000000000000000000000000") {
		log.Println("passport is null, user is not registred")
		return false
	} else {
		return true
	}
}
*/


/*
func (client *Client) CheckItemCreated(ctx context.Context, fileID string, start time.Time) (bool, error) {
	iter, err := client.Signleton.Contract.FilterItemCreated(&bind.FilterOpts{
		Start: uint64(start.Unix()),
	}, []string{fileID})
	if err != nil {
		return false, err
	}

	return iter.Next(), iter.Error()
}

// Returns count of remaining items, if remaining == 0 all files was created.
func (client *Client) CheckItemsCreated(ctx context.Context, fileIDs []string, start time.Time) (int, error) {
	count := len(fileIDs)
	iter, err := client.Signleton.Contract.FilterItemCreated(&bind.FilterOpts{
		Start: uint64(start.Unix()),
	}, fileIDs)
	if err != nil {
		return count, err
	}

	for iter.Next() {
		count -= 1
	}
	return count, iter.Error()
}
*/
