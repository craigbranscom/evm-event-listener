package utils

const (
	//event signatures
	EVENT_SIG_DAO_FACTORY_CREATED      = "DAOFactoryCreated(address)"
	EVENT_SIG_NFT_FACTORY_CREATED      = "NFTFactoryCreated(address)"
	EVENT_SIG_PROPOSAL_FACTORY_CREATED = "ProposalFactoryCreated(address)"
	EVENT_SIG_CONTRACT_PRODUCED        = "ContractProduced(address,address,address)"
	EVENT_SIG_MEMBER_REGISTERED        = "MemberRegistered(address,address,uint256,address)"
	EVENT_SIG_PROPOSAL_CREATED         = "ProposalCreated(address,address,uint256,address)"
	EVENT_SIG_TRANSFER                 = "Transfer(address,address,uint256)"
	EVENT_SIG_STATUS_UPDATED           = "StatusUpdated(string)"
	EVENT_SIG_VOTE_CAST                = "VoteCast(address,uint256,uint8)"
	EVENT_SIG_BAY_CREATED              = "BayCreated(address,address)"
	EVENT_SIG_ASK_CREATED              = "AskCreated(bytes32,address,address,uint256,address,uint256)"
	EVENT_SIG_BID_CREATED              = "BidCreated(bytes32,address,address,uint256,address,uint256)"
	EVENT_SIG_ORDER_FULFILLED          = "OrderFulfilled(bytes32)"
	EVENT_SIG_ASK_CANCELLED            = "AskCancelled(bytes32)"
	EVENT_SIG_BID_CANCELLED            = "BidCancelled(bytes32)"

	//event hashes (keccak256 of respective signature)
	EVENT_HASH_DAO_FACTORY_CREATED      = "0xaab2cc7f2123749388ff257fd1a3f45420dc9ba35a1d2feafcb9c44d63a4887a"
	EVENT_HASH_PROPOSAL_FACTORY_CREATED = "0x35d71b99bdb343ca2dd2814592bbf63b5a379cf9b151185dcd862a0199dd8753"
	EVENT_HASH_CONTRACT_PRODUCED        = "0x5db1f34edd85623861e21337b01e1c2035c3383b68b101d961bc2337c580a962"
	EVENT_HASH_MEMBER_REGISTERED        = "0x6d6a6e32530d4895a149a76159bcea283d7bbd32170c5e5a95b7c02a9098c38a"
	EVENT_HASH_PROPOSAL_CREATED         = "0xa33abb4d0226d76a56e89b18b3c8f3dec914fce06449b6ff646962c19ddc971e"
	EVENT_HASH_TRANSFER                 = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	EVENT_HASH_STATUS_UPDATED           = "0x69f380a4c820112d1c3e49f6e19e1d33c66d98a47f9c8ad3221a82b79b0bfcbc"
	EVENT_HASH_VOTE_CAST                = "0x251f168c3d37dca4f3922f4aa12852da0c3de8ef70710e0cf58eba85805168a2"
	EVENT_HASH_BAY_CREATED              = "0xafcc4615a5c6ce81e1ee1ecc9f150c249104ed97f69f737f6eb8fec74b8a91e4"
	EVENT_HASH_ASK_CREATED              = "0x22649edb01e89b0ff40ea9cbf2302950808ee082b372ffed1423e6e6cef8f49a"
	EVENT_HASH_BID_CREATED              = "0xb70335faa97c8bf1fd6727d01143b2c03fd75096ed19e11a7cd4764a89ae21ef"
	EVENT_HASH_ORDER_FULFILLED          = "0x48224e98144457f56f46e15959589d8155a9e29ca79439ab9ba37f60561b5c56"
	EVENT_HASH_ASK_CANCELLED            = "0xa7e78ab2d839285f692f95318e21738ed76fe005b13097399f19eb043f46cbdd"
	EVENT_HASH_BID_CANCELLED            = "0xb56dc4096011ba5fd2e46e5c3e7b04dec423b5e7b5fce9a17a419d77c832177c"
)

// eventSignature := []byte("BayCreated(address,address)")
// hash := crypto.Keccak256Hash(eventSignature)
// fmt.Println(">>>>>>>", hash.Hex())
