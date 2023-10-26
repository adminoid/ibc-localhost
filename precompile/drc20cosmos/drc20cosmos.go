package drc20cosmos

import (
	"bytes"
	"embed"
	"encoding/hex"
	"encoding/json"
	sdkerrors "github.com/adminoid/cosmos-sdk/types/errors"
	bankkeeper "github.com/adminoid/cosmos-sdk/x/bank/keeper"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"strings"

	"bitbucket.org/decimalteam/go-smart-node/x/coin/types"
	sdk "github.com/adminoid/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethante "github.com/evmos/evmos/v14/app/ante"
	"github.com/evmos/evmos/v14/x/evm/statedb"

	// todo -- there is doubtful code
	evmtypes "github.com/evmos/evmos/v14/x/evm/types"
	evm "github.com/evmos/evmos/v14/x/evm"
	//evm "github.com/ethereum/go-ethereum"
)

const (
	AddressForContractOwner = "0x2941b073ad6b59b1de4fc70c69e39a9e130b25ce"
	createBin               = "60806040523480156200001157600080fd5b506040516200299d3803806200299d833981810160405281019062000037919062000445565b826006908162000048919062000720565b5081600590816200005a919062000720565b506012600460006101000a81548160ff021916908360ff16021790555080600381905550600354600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550620000d5336200014860201b60201c565b3373ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60035460405162000137919062000818565b60405180910390a3505050620008de565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620001ba576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001b190620008bc565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620002e08262000295565b810181811067ffffffffffffffff82111715620003025762000301620002a6565b5b80604052505050565b60006200031762000277565b9050620003258282620002d5565b919050565b600067ffffffffffffffff821115620003485762000347620002a6565b5b620003538262000295565b9050602081019050919050565b60005b838110156200038057808201518184015260208101905062000363565b60008484015250505050565b6000620003a36200039d846200032a565b6200030b565b905082815260208101848484011115620003c257620003c162000290565b5b620003cf84828562000360565b509392505050565b600082601f830112620003ef57620003ee6200028b565b5b8151620004018482602086016200038c565b91505092915050565b6000819050919050565b6200041f816200040a565b81146200042b57600080fd5b50565b6000815190506200043f8162000414565b92915050565b60008060006060848603121562000461576200046062000281565b5b600084015167ffffffffffffffff81111562000482576200048162000286565b5b6200049086828701620003d7565b935050602084015167ffffffffffffffff811115620004b457620004b362000286565b5b620004c286828701620003d7565b9250506040620004d5868287016200042e565b9150509250925092565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200053257607f821691505b602082108103620005485762000547620004ea565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620005b27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000573565b620005be868362000573565b95508019841693508086168417925050509392505050565b6000819050919050565b600062000601620005fb620005f5846200040a565b620005d6565b6200040a565b9050919050565b6000819050919050565b6200061d83620005e0565b620006356200062c8262000608565b84845462000580565b825550505050565b600090565b6200064c6200063d565b6200065981848462000612565b505050565b5b8181101562000681576200067560008262000642565b6001810190506200065f565b5050565b601f821115620006d0576200069a816200054e565b620006a58462000563565b81016020851015620006b5578190505b620006cd620006c48562000563565b8301826200065e565b50505b505050565b600082821c905092915050565b6000620006f560001984600802620006d5565b1980831691505092915050565b6000620007108383620006e2565b9150826002028217905092915050565b6200072b82620004df565b67ffffffffffffffff811115620007475762000746620002a6565b5b62000753825462000519565b6200076082828562000685565b600060209050601f83116001811462000798576000841562000783578287015190505b6200078f858262000702565b865550620007ff565b601f198416620007a8866200054e565b60005b82811015620007d257848901518255600182019150602085019450602081019050620007ab565b86831015620007f25784890151620007ee601f891682620006e2565b8355505b6001600288020188555050505b505050505050565b62000812816200040a565b82525050565b60006020820190506200082f600083018462000807565b92915050565b600082825260208201905092915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000620008a460268362000835565b9150620008b18262000846565b604082019050919050565b60006020820190508181036000830152620008d78162000895565b9050919050565b6120af80620008ee6000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c8063893d20e8116100ad578063a9059cbb11610071578063a9059cbb1461035d578063b09f12661461038d578063d28d8852146103ab578063dd62ed3e146103c9578063f2fde38b146103f95761012c565b8063893d20e8146102a35780638da5cb5b146102c157806395d89b41146102df5780639dc29fac146102fd578063a457c2d71461032d5761012c565b806332424aa3116100f457806332424aa3146101eb578063395093511461020957806340c10f191461023957806370a0823114610269578063715018a6146102995761012c565b806306fdde0314610131578063095ea7b31461014f57806318160ddd1461017f57806323b872dd1461019d578063313ce567146101cd575b600080fd5b610139610415565b60405161014691906117c1565b60405180910390f35b6101696004803603810190610164919061187c565b6104a7565b60405161017691906118d7565b60405180910390f35b6101876104c5565b6040516101949190611901565b60405180910390f35b6101b760048036038101906101b2919061191c565b6104cf565b6040516101c491906118d7565b60405180910390f35b6101d56105a8565b6040516101e2919061198b565b60405180910390f35b6101f36105bf565b604051610200919061198b565b60405180910390f35b610223600480360381019061021e919061187c565b6105d2565b60405161023091906118d7565b60405180910390f35b610253600480360381019061024e919061187c565b610685565b60405161026091906118d7565b60405180910390f35b610283600480360381019061027e91906119a6565b610730565b6040516102909190611901565b60405180910390f35b6102a1610779565b005b6102ab6108cc565b6040516102b891906119e2565b60405180910390f35b6102c96108db565b6040516102d691906119e2565b60405180910390f35b6102e7610904565b6040516102f491906117c1565b60405180910390f35b6103176004803603810190610312919061187c565b610996565b60405161032491906118d7565b60405180910390f35b6103476004803603810190610342919061187c565b610a41565b60405161035491906118d7565b60405180910390f35b6103776004803603810190610372919061187c565b610b0e565b60405161038491906118d7565b60405180910390f35b610395610b2c565b6040516103a291906117c1565b60405180910390f35b6103b3610bba565b6040516103c091906117c1565b60405180910390f35b6103e360048036038101906103de91906119fd565b610c48565b6040516103f09190611901565b60405180910390f35b610413600480360381019061040e91906119a6565b610ccf565b005b60606006805461042490611a6c565b80601f016020809104026020016040519081016040528092919081815260200182805461045090611a6c565b801561049d5780601f106104725761010080835404028352916020019161049d565b820191906000526020600020905b81548152906001019060200180831161048057829003601f168201915b5050505050905090565b60006104bb6104b4610d70565b8484610d78565b6001905092915050565b6000600354905090565b60006104dc848484610f41565b61059d846104e8610d70565b61059885604051806060016040528060288152602001611fe560289139600260008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600061054e610d70565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546111cd9092919063ffffffff16565b610d78565b600190509392505050565b6000600460009054906101000a900460ff16905090565b600460009054906101000a900460ff1681565b600061067b6105df610d70565b8461067685600260006105f0610d70565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461123190919063ffffffff16565b610d78565b6001905092915050565b600061068f610d70565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461071c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071390611ae9565b60405180910390fd5b610726838361128f565b6001905092915050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b610781610d70565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461080e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161080590611ae9565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60006108d66108db565b905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606005805461091390611a6c565b80601f016020809104026020016040519081016040528092919081815260200182805461093f90611a6c565b801561098c5780601f106109615761010080835404028352916020019161098c565b820191906000526020600020905b81548152906001019060200180831161096f57829003601f168201915b5050505050905090565b60006109a0610d70565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610a2d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2490611ae9565b60405180910390fd5b610a378383611418565b6001905092915050565b6000610b04610a4e610d70565b84610aff856040518060600160405280602581526020016120336025913960026000610a78610d70565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546111cd9092919063ffffffff16565b610d78565b6001905092915050565b6000610b22610b1b610d70565b8484610f41565b6001905092915050565b60058054610b3990611a6c565b80601f0160208091040260200160405190810160405280929190818152602001828054610b6590611a6c565b8015610bb25780601f10610b8757610100808354040283529160200191610bb2565b820191906000526020600020905b815481529060010190602001808311610b9557829003601f168201915b505050505081565b60068054610bc790611a6c565b80601f0160208091040260200160405190810160405280929190818152602001828054610bf390611a6c565b8015610c405780601f10610c1557610100808354040283529160200191610c40565b820191906000526020600020905b815481529060010190602001808311610c2357829003601f168201915b505050505081565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b610cd7610d70565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610d64576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d5b90611ae9565b60405180910390fd5b610d6d816115bb565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610de7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dde90611b7b565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610e56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4d90611c0d565b60405180910390fd5b80600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610f349190611901565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610fb0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fa790611c9f565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361101f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161101690611d31565b60405180910390fd5b61108b8160405180606001604052806026815260200161200d60269139600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546111cd9092919063ffffffff16565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061112081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461123190919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516111c09190611901565b60405180910390a3505050565b6000838311158290611215576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161120c91906117c1565b60405180910390fd5b50600083856112249190611d80565b9050809150509392505050565b60008082846112409190611db4565b905083811015611285576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161127c90611e34565b60405180910390fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036112fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112f590611ea0565b60405180910390fd5b6113138160035461123190919063ffffffff16565b60038190555061136b81600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461123190919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161140c9190611901565b60405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611487576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161147e90611f32565b60405180910390fd5b6114f38160405180606001604052806022815260200161205860229139600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546111cd9092919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061154b816003546116e790919063ffffffff16565b600381905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516115af9190611901565b60405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361162a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161162190611fc4565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600061172983836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506111cd565b905092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561176b578082015181840152602081019050611750565b60008484015250505050565b6000601f19601f8301169050919050565b600061179382611731565b61179d818561173c565b93506117ad81856020860161174d565b6117b681611777565b840191505092915050565b600060208201905081810360008301526117db8184611788565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611813826117e8565b9050919050565b61182381611808565b811461182e57600080fd5b50565b6000813590506118408161181a565b92915050565b6000819050919050565b61185981611846565b811461186457600080fd5b50565b60008135905061187681611850565b92915050565b60008060408385031215611893576118926117e3565b5b60006118a185828601611831565b92505060206118b285828601611867565b9150509250929050565b60008115159050919050565b6118d1816118bc565b82525050565b60006020820190506118ec60008301846118c8565b92915050565b6118fb81611846565b82525050565b600060208201905061191660008301846118f2565b92915050565b600080600060608486031215611935576119346117e3565b5b600061194386828701611831565b935050602061195486828701611831565b925050604061196586828701611867565b9150509250925092565b600060ff82169050919050565b6119858161196f565b82525050565b60006020820190506119a0600083018461197c565b92915050565b6000602082840312156119bc576119bb6117e3565b5b60006119ca84828501611831565b91505092915050565b6119dc81611808565b82525050565b60006020820190506119f760008301846119d3565b92915050565b60008060408385031215611a1457611a136117e3565b5b6000611a2285828601611831565b9250506020611a3385828601611831565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680611a8457607f821691505b602082108103611a9757611a96611a3d565b5b50919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611ad360208361173c565b9150611ade82611a9d565b602082019050919050565b60006020820190508181036000830152611b0281611ac6565b9050919050565b7f42455032303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000611b6560248361173c565b9150611b7082611b09565b604082019050919050565b60006020820190508181036000830152611b9481611b58565b9050919050565b7f42455032303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b6000611bf760228361173c565b9150611c0282611b9b565b604082019050919050565b60006020820190508181036000830152611c2681611bea565b9050919050565b7f42455032303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b6000611c8960258361173c565b9150611c9482611c2d565b604082019050919050565b60006020820190508181036000830152611cb881611c7c565b9050919050565b7f42455032303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b6000611d1b60238361173c565b9150611d2682611cbf565b604082019050919050565b60006020820190508181036000830152611d4a81611d0e565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611d8b82611846565b9150611d9683611846565b9250828203905081811115611dae57611dad611d51565b5b92915050565b6000611dbf82611846565b9150611dca83611846565b9250828201905080821115611de257611de1611d51565b5b92915050565b7f536166654d6174683a206164646974696f6e206f766572666c6f770000000000600082015250565b6000611e1e601b8361173c565b9150611e2982611de8565b602082019050919050565b60006020820190508181036000830152611e4d81611e11565b9050919050565b7f42455032303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000611e8a601f8361173c565b9150611e9582611e54565b602082019050919050565b60006020820190508181036000830152611eb981611e7d565b9050919050565b7f42455032303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b6000611f1c60218361173c565b9150611f2782611ec0565b604082019050919050565b60006020820190508181036000830152611f4b81611f0f565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611fae60268361173c565b9150611fb982611f52565b604082019050919050565b60006020820190508181036000830152611fdd81611fa1565b905091905056fe42455032303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636542455032303a207472616e7366657220616d6f756e7420657863656564732062616c616e636542455032303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726f42455032303a206275726e20616d6f756e7420657863656564732062616c616e6365a2646970667358221220f3fcca04a355fc1eb8680f1da4227ca90709a8b82d7bf60ee452fd7ba47ce6f164736f6c63430008110033000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000001ed36a00000000000000000000000000000000000000000000000000000000000000006647361647361000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000066473616473610000000000000000000000000000000000000000000000000000"
	gasCostCreation         = 3026345
	contractCreateForCoin   = "contractCreateForCoin"
	contractAction          = "contractAction"

	firstIncrease = 5
)

func createEthMessage(tx *ethtypes.Transaction) ethtypes.Message {
	return ethtypes.NewMessage(
		common.HexToAddress(AddressForContractOwner),
		tx.To(),
		tx.Nonce(),
		tx.Value(),
		tx.Gas(),
		new(big.Int).Set(tx.GasPrice()),
		new(big.Int).Set(tx.GasFeeCap()),
		new(big.Int).Set(tx.GasTipCap()),
		tx.Data(),
		tx.AccessList(),
		false,
	)
}

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

type Drc20Cosmos struct {
	abi.ABI
	ctx        sdk.Context
	evmKeeper  ethante.EVMKeeper
	bankKeeper bankkeeper.Keeper
	stateDB    *statedb.StateDB
	evm        evm.EVM
	cfg        *evmtypes.EVMConfig
	Coin       types.Coin
}

// NewDrc20Cosmos create instance of contract
func NewDrc20Cosmos(ctx sdk.Context,
	evmKeeper ethante.EVMKeeper,
	bankKeeper bankkeeper.Keeper,
	coinAction types.Coin,
) (*Drc20Cosmos, error) {
	abiBz, err := f.ReadFile("abi.json")
	if err != nil {
		return nil, err
	}

	newAbi, err := abi.JSON(bytes.NewReader(abiBz))
	if err != nil {
		return nil, err
	}

	params := evmKeeper.GetParams(ctx)
	ethCfg := params.ChainConfig.EthereumConfig(evmKeeper.ChainID())
	baseFee := evmKeeper.GetBaseFee(ctx, ethCfg)

	contractCreateTx := &ethtypes.AccessListTx{
		GasPrice: nil,
		Gas:      53000,
		To:       nil,
		Data:     []byte("empty code"),
		Nonce:    1,
	}
	ethTx := ethtypes.NewTx(contractCreateTx)
	ethMsg := &evmtypes.MsgEthereumTx{}
	ethMsg.FromEthereumTx(ethTx)
	ethMsg.From = AddressForContractOwner

	coreMsg := createEthMessage(ethMsg.AsTransaction())

	cfg := &evmtypes.EVMConfig{
		ChainConfig: ethCfg,
		Params:      params,
		CoinBase:    common.Address{},
		BaseFee:     baseFee,
	}

	stateNewDB := statedb.New(ctx, evmKeeper, statedb.NewEmptyTxConfig(common.BytesToHash(ctx.HeaderHash().Bytes())))
	evmNew := evmKeeper.NewEVM(ctx, coreMsg, cfg, evmtypes.NewNoOpTracer(), stateNewDB)

	// nonce := stateDB.GetNonce(common.HexToAddress("0x2941b073ad6b59b1de4fc70c69e39a9e130b25ce"))

	// stateDB.SetNonce(common.HexToAddress("0x2941b073ad6b59b1de4fc70c69e39a9e130b25ce"), nonce)
	// ret, _, leftoverGas, vmErr = evm.Create(sender, msg.Data(), leftoverGas, msg.Value())
	// stateDB.SetNonce(sender.Address(), msg.Nonce()+1)

	return &Drc20Cosmos{
		ABI:        newAbi,
		ctx:        ctx,
		evmKeeper:  evmKeeper,
		bankKeeper: bankKeeper,
		stateDB:    stateNewDB,
		evm:        evmNew,
		cfg:        cfg,
		Coin:       coinAction,
	}, nil
}

// CreateContractIfNotSet creation contract if not not to coin
func (drc *Drc20Cosmos) CreateContractIfNotSet() (bool, error) {

	sender := vm.AccountRef(common.HexToAddress(AddressForContractOwner))

	if drc.Coin.Drc20Address == "" {
		inputContract, err := hex.DecodeString(createBin)
		if err != nil {
			return false, err
		}

		// receive nonce for owner address for new contract
		nonce := drc.stateDB.GetNonce(common.HexToAddress(AddressForContractOwner))

		contractCreateTx := &ethtypes.AccessListTx{
			GasPrice: big.NewInt(0),
			Gas:      params.TxGasContractCreation,
			To:       nil,
			Data:     inputContract,
			Nonce:    nonce,
		}

		ethTx := ethtypes.NewTx(contractCreateTx)

		// access list preparation is moved from ante handler to here, because it's needed when `ApplyMessage` is called
		// under contexts where ante handlers are not run, for example `eth_call` and `eth_estimateGas`.
		if rules := drc.cfg.ChainConfig.Rules(big.NewInt(drc.ctx.BlockHeight()), drc.cfg.ChainConfig.MergeNetsplitBlock != nil); rules.IsBerlin {
			drc.stateDB.PrepareAccessList(common.HexToAddress(AddressForContractOwner), nil, drc.evm.ActivePrecompiles(rules), ethTx.AccessList())
		}

		drc.stateDB.SetNonce(common.HexToAddress(AddressForContractOwner), nonce)

		ret, contractAddr, _, vmErr := drc.evm.Create(sender, inputContract, gasCostCreation, big.NewInt(0))
		drc.stateDB.SetNonce(sender.Address(), nonce+1)

		drc.ctx.Logger().With(contractAddr.Hex()).Info(contractAddr.Hex())
		drc.ctx.Logger().With(ret).Info("Result create contract")

		if vmErr != nil {
			drc.ctx.Logger().Info(vmErr.Error())
			//drc.ctx.Logger().With(vmErr, sender.Address().Hex()).Info("failed to encode log vmErr %T")
			//return false, sdkerrors.ErrUnknownRequest.Wrapf("failed to encode log vmErr %T", vmErr)
		}

		// The dirty states in `StateDB` is either committed or discarded after return
		if err := drc.stateDB.Commit(); err != nil {
			drc.ctx.Logger().Info(vmErr.Error())
			//return false, sdkerrors.ErrUnknownRequest.Wrapf("failed to encode log Commit %T", err)
		}

		txLogAttrs := make([]sdk.Attribute, len(evmtypes.NewLogsFromEth(drc.stateDB.Logs())))
		for i, log := range drc.stateDB.Logs() {
			value, err := json.Marshal(log)
			if err != nil {
				return false, sdkerrors.ErrUnknownRequest.Wrapf("failed to encode log %T", err)
			}
			txLogAttrs[i] = sdk.NewAttribute(evmtypes.AttributeKeyTxLog, string(value))
		}

		// emit events
		drc.ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				contractCreateForCoin,
				txLogAttrs...,
			),
			sdk.NewEvent(
				evmtypes.AttributeKeyContractAddress,
				sdk.NewAttribute(contractAction, contractAddr.Hex()),
			),
		})

		drc.Coin.Drc20Address = strings.ToLower(contractAddr.Hex())
	}
	drc.ctx.Logger().Info(drc.Coin.Drc20Address)

	return true, nil
}
