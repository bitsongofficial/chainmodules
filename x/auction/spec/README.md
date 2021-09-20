# Auction Specification

## Abstract
This specification describes how to open a new auction on the chain. Auction module is needed for users to buy and sell their NFTs through a marketplace.
There are 3 auction types supported.(Single Auction, Open Edition Auction, Limited Edition Auction)

### Single Auction
The auction type where the winner can be the highest bid.

### Open Edition Auction
The auction type where anyone can get the NFT with any price, even free.

### Limited Edition Auction
The auction type where limited number of highest bidders can be the winners.

## Contents

1. **[State](01_state.md)**
    - [Auction](01_state.md#Auction)
    - [Bid](01_state.md#Bid)
2. **[Messages](02_messages.md)**
    - [MsgOpenAuction](02_messages.md#MsgOpenAuction)
    - [MsgEditAuction](02_messages.md#MsgEditAuction)
    - [MsgCancelAuction](02_messages.md#MsgCancelAuction)
    - [MsgOpenBid](02_messages.md#MsgOpenBid)
    - [MsgCancelBid](02_messages.md#MsgCancelBid)
    - [MsgWithdraw](02_messages.md#MsgWithdraw)
3. **[Events](03_events.md)**
    - [MsgIssueFanToken](03_events.md#MsgIssueFanToken)
    - [MsgEditFanToken](03_events.md#MsgEditFanToken)
    - [MsgMintFanToken](03_events.md#MsgMintFanToken)
    - [MsgBurnFanToken](03_events.md#MsgBurnFanToken)
    - [MsgTransferFanTokenOwner](03_events.md#MsgTransferFanTokenOwner)
4. **[Future Improvements](04_future_improvements.md)**