# Events

The auction module emits the following events:
## MsgOpenAuction

| Type               | Attribute Key | Attribute Value   |
| :----------------- | :------------ | :---------------- |
| open_auction       | auction_id    | {auction_id}      |
| open_auction       | creator       | {creatorAddreses} |
| message            | module        | auction           |
| message            | sender        | {ownerAddress}    |


## MsgEditAuction

| Type               | Attribute Key | Attribute Value   |
| :----------------- | :------------ | :---------------- |
| edit_auction       | auction_id    | {auction_id}      |
| edit_auction       | owner         | {ownerAddress}    |
| message            | module        | auction           |
| message            | sender        | {ownerAddress}    |


## MsgCancelAuction

| Type               | Attribute Key | Attribute Value   |
| :----------------- | :------------ | :---------------- |
| cancel_auction     | auction_id    | {auction_id}      |
| cancel_auction     | owner         | {ownerAddress}    |
| message            | module        | auction           |
| message            | sender        | {ownerAddress}    |


## MsgOpenBid

| Type               | Attribute Key | Attribute Value   |
| :----------------- | :------------ | :---------------- |
| open_bid           | auction_id    | {auction_id}      |
| open_bid           | bidder        | {bidderAddress}   |
| message            | module        | auction           |
| message            | sender        | {bidderAddress}   |


## MsgCancelBid

| Type               | Attribute Key | Attribute Value   |
| :----------------- | :------------ | :---------------- |
| cancel_bid         | auction_id    | {auction_id}      |
| cancel_bid         | owner         | {bidderAddress}   |
| message            | module        | auction           |
| message            | sender        | {bidderAddress}   |


## MsgWithdraw

| Type               | Attribute Key | Attribute Value    |
| :----------------- | :------------ | :----------------- |
| withdraw           | auction_id    | {auction_id}       |
| withdraw           | recipient     | {recipientAddress} |
| message            | module        | auction            |
| message            | sender        | {recipientAddress} |
