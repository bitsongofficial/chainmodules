package simulation

// DONTCOVER

import (
	"bytes"
	"fmt"

	gogotypes "github.com/gogo/protobuf/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/bitsongofficial/chainmodules/x/auction/types"
)

// NewDecodeStore unmarshals the KVPair's Value to the corresponding auction type
func NewDecodeStore(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.PrefixAuctionById):
			var auctionA, auctionB types.Auction
			cdc.MustUnmarshalBinaryBare(kvA.Value, &auctionA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &auctionB)
			return fmt.Sprintf("%v\n%v", auctionA, auctionB)
		case bytes.Equal(kvA.Key[:1], types.PrefixAuctionsByOwner):
			var ownerA, ownerB gogotypes.Value
			cdc.MustUnmarshalBinaryBare(kvA.Value, &ownerA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &ownerB)
			return fmt.Sprintf("%v\n%v", ownerA, ownerB)
		case bytes.Equal(kvA.Key[:1], types.PrefixBidsByAuctionId):
			var bidA, bidB types.Bid
			cdc.MustUnmarshalBinaryBare(kvA.Value, &bidA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &bidB)
			return fmt.Sprintf("%v\n%v", bidA, bidB)
		case bytes.Equal(kvA.Key[:1], types.PrefixBidsByBidder):
			var bidderA, bidderB gogotypes.StringValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &bidderA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &bidderB)
			return fmt.Sprintf("%v\n%v", bidderA, bidderB)
		default:
			panic(fmt.Sprintf("invalid %s key prefix %X", types.ModuleName, kvA.Key[:1]))
		}
	}
}
