package signature

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/pkg/errors"
	"math/big"
	"reflect"
	"strconv"
	"time"
)

type StakeAndBakePermit struct {
	Owner    common.MixedcaseAddress `json:"owner"`
	Spender  common.MixedcaseAddress `json:"to"`
	Value    math.Decimal256         `json:"value"`
	Nonce    big.Int                 `json:"nonce"`
	Deadline int64                   `json:"deadline"`

	LBTC    common.MixedcaseAddress `json:"lbtc"`
	ChainId math.HexOrDecimal256    `json:"chainId"`
}

func NewStakeAndBakePermitFromJson(jsonTypedData string) (*StakeAndBakePermit, error) {
	var data apitypes.TypedData

	if err := json.Unmarshal([]byte(jsonTypedData), &data); err != nil {
		return nil, err
	}

	p := &StakeAndBakePermit{
		ChainId: math.HexOrDecimal256(big.Int{}),
	}

	owner, err := getParam[string](data, "owner")
	if err != nil {
		return nil, err
	}
	if err := p.SetOwnerFromString(owner); err != nil {
		return nil, err
	}

	spender, err := getParam[string](data, "spender")
	if err != nil {
		return nil, err
	}
	if err := p.SetSpenderFromString(spender); err != nil {
		return nil, err
	}

	val, err := getParam[string](data, "value")
	if err != nil {
		return nil, err
	}
	if err := p.SetValueFromString(val); err != nil {
		return nil, err
	}

	nonce, err := getParam[string](data, "nonce")
	if err != nil {
		return nil, err
	}
	if err := p.SetNonceFromString(nonce); err != nil {
		return nil, err
	}

	deadline, err := getParam[string](data, "deadline")
	if err != nil {
		return nil, err
	}
	if err := p.SetDeadlineFromString(deadline); err != nil {
		return nil, err
	}

	lbtc, err := common.NewMixedcaseAddressFromString(data.Domain.VerifyingContract)
	if err != nil {
		return nil, err
	}
	p.LBTC = *lbtc

	p.ChainId = *data.Domain.ChainId

	return p, nil
}

func (p *StakeAndBakePermit) SetOwnerFromString(val string) error {
	ownerAddr, err := common.NewMixedcaseAddressFromString(val)
	if err != nil {
		return err
	}
	p.Owner = *ownerAddr

	return nil
}

func (p *StakeAndBakePermit) SetSpenderFromString(val string) error {
	spenderAddr, err := common.NewMixedcaseAddressFromString(val)
	if err != nil {
		return err
	}
	p.Spender = *spenderAddr

	return nil
}

func (p *StakeAndBakePermit) SetValueFromString(val string) error {
	value, ok := new(big.Int).SetString(val, 10)
	if !ok {
		return errors.New("invalid value string")
	}
	p.Value = (math.Decimal256)(*value)

	return nil
}

func (p *StakeAndBakePermit) SetNonceFromString(val string) error {
	nonce, ok := new(big.Int).SetString(val, 10)
	if !ok {
		return errors.New("invalid nonce string")
	}
	p.Nonce = *nonce

	return nil
}

func (p *StakeAndBakePermit) SetDeadlineFromString(val string) error {
	deadline, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return err
	}
	p.Deadline = deadline
	return nil
}

func (p StakeAndBakePermit) Validate() error {
	// TODO: implement more validations

	// make sure expiry time is in the future
	if p.Deadline < time.Now().Unix() {
		return errors.New("expiry time is in the past")
	}
	return nil
}

// ToTypedData converts the tx to a EIP-712 Typed Data structure for signing
func (p StakeAndBakePermit) ToTypedData() apitypes.TypedData {
	var domainType = []apitypes.Type{
		{Name: "verifyingContract", Type: "address"},
		{Name: "chainId", Type: "uint256"},
	}

	permitTypedData := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": domainType,
			"Permit": []apitypes.Type{
				{Name: "owner", Type: "address"},
				{Name: "spender", Type: "address"},
				{Name: "value", Type: "uint256"},
				{Name: "nonce", Type: "uint256"},
				{Name: "deadline", Type: "uint256"},
			},
		},
		Domain: apitypes.TypedDataDomain{
			VerifyingContract: p.LBTC.Address().Hex(),
			ChainId:           &p.ChainId,
		},
		PrimaryType: "Permit",
		Message: apitypes.TypedDataMessage{
			"owner":    p.Owner.Address().Hex(),
			"spender":  p.Spender.Address().Hex(),
			"value":    p.Value.String(),
			"nonce":    fmt.Sprintf("%d", p.Nonce.Uint64()),
			"deadline": fmt.Sprintf("%d", p.Deadline),
		},
	}
	return permitTypedData
}

func getParam[K any](d apitypes.TypedData, msgKey string) (v K, err error) {
	if val, ok := d.Message[msgKey]; ok {
		if v, ok := val.(K); ok {
			return v, nil
		} else {
			err = errors.Errorf("%s has wrong type (%s)", msgKey, reflect.TypeOf(val).String())
		}
	} else {
		err = errors.Errorf("%s not presented", msgKey)
	}
	return
}
