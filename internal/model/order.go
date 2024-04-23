package model

import (
	"time"

	"github.com/enrique/cron-bridge/internal/blockchain/theta"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status string

const (
	Created   Status = "created"
	Pending   Status = "pending"
	Approved  Status = "approved"
	Rejected  Status = "rejected"
	Paid      Status = "paid"
	Refund    Status = "refund"
	InTransit Status = "in transit" // Added for the bridging process
	Completed Status = "completed"  // Added to mark the end of the process
)

type StatusChange struct {
	Status    Status    `bson:"status" json:"status"`
	ChangedAt time.Time `bson:"changed_at" json:"changed_at"`
}

type Chain string

const (
	Ethereum Chain = "ethereum"
	Theta    Chain = "theta"
)

type Contract struct {
	Name            string `bson:"name,omitempty" json:"name,omitempty" validate:"required"`
	Symbol          string `bson:"symbol,omitempty" json:"symbol,omitempty" validate:"required"`
	Decimal         int    `bson:"decimal,omitempty" json:"decimal,omitempty" validate:"required"`
	ContractAddress string `bson:"contract_address,omitempty" json:"contract_address,omitempty" validate:"required"`
	Chain           string `bson:"chain,omitempty" json:"chain,omitempty" validate:"required"`
}

// http://www.thetascan.io/api/contract/?contract=0x3da3d8cde7b12cd2cbb688e2655bcacd8946399d
type Order struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" validate:"required"`
	PublicOrderId         primitive.ObjectID `bson:"public_order_id,omitempty" json:"public_order_id,omitempty" validate:"required"`
	Sequence              int64              `bson:"sequence,omitempty" json:"sequence,omitempty" validate:"required"`
	CustomerWalletAddress string             `bson:"customer_wallet_address,omitempty" json:"customer_wallet_address,omitempty" validate:"required"`
	TokenAmount           float64            `bson:"token_amount,omitempty" json:"token_amount,omitempty" validate:"required,token_amount"`
	OrderWalletAddress    string             `bson:"order_wallet_address,omitempty" json:"order_wallet_address,omitempty" validate:"required"`
	OrderStatus           Status             `bson:"order_status,omitempty" json:"order_status,omitempty" validate:"required"`
	StatusChanges         []StatusChange     `bson:"status_changes,omitempty" json:"status_changes,omitempty"`
	FromChain             Chain              `bson:"from_chain,omitempty" json:"from_chain,omitempty"`
	ToChain               Chain              `bson:"to_chain,omitempty" json:"to_chain,omitempty"`
	From                  Contract           `bson:"from_contract,omitempty" json:"from_contract,omitempty"`
	To                    Contract           `bson:"to_contract,omitempty" json:"to_contract,omitempty"`
	CreatedAt             time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt             time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	Prices                theta.Price        `bson:"prices,omitempty" json:"prices,omitempty"`
	EthGasPriceInUSD      float64            `bson:"eth_gas_price_usd,omitempty" json:"eth_gas_price_usd,omitempty"`
	EthGasPriceInReplay   float64            `bson:"eth_gas_price_in_replay_token,omitempty" json:"eth_gas_price_in_replay_token,omitempty"`
	TotalOrderAmount      float64            `bson:"total_order_amount,omitempty" json:"total_order_amount,omitempty"`
}
