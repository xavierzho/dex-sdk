package dexsdk

// ChainId Public Blockchain Known Identify
type ChainId int64

const (
	EthMain ChainId = 1  // ethereum mainnet
	Ropsten ChainId = 3  // ethereum Ropsten testnet
	Rinkeby ChainId = 4  // ethereum Rinkeby testnet
	Goerli  ChainId = 5  // ethereum Goerli testnet
	Kovan   ChainId = 42 // ethereum Kovan testnet

	BscMain ChainId = 56 // bsc mainnet
	BscTest ChainId = 97 // bsc testnet

	OkcMain ChainId = 66 // okc mainnet
	OkcTest ChainId = 65 // okc testnet

	HecoMain ChainId = 128 // heco mainnet
	HecoTest ChainId = 256 // heco testnet

	PolygonMain ChainId = 137   // polygon mainnet
	Mumbai      ChainId = 80001 // polygon mumbai testnet

	Ganache ChainId = 8545 // local testnet
)

// String Name of the Chain id
func (chain ChainId) String() string {
	switch chain {
	case EthMain:
		return "Ethereum Mainnet"
	case Ropsten:
		return "Ethereum Ropsten Testnet"
	case Rinkeby:
		return "Ethereum Rinkeby Testnet"
	case Goerli:
		return "Ethereum Goerli Testnet"
	case Kovan:
		return "Ethereum Kovan Testnet"
	case BscMain:
		return "Bsc Mainnet"
	case BscTest:
		return "Bsc Testnet"
	case HecoMain:
		return "Heco Chain Mainnet"
	case HecoTest:
		return "Heco Chain Testnet"
	case PolygonMain:
		return "Polygon Mainnet"
	case Mumbai:
		return "Polygon Mumbai Testnet"
	case Ganache:
		return "Ganache Testnet"
	case OkcMain:
		return "OKC Mainnet"
	case OkcTest:
		return "OKC Testnet"
	default:
		return ""
	}
}

// Symbol of the chain id
func (chain ChainId) Symbol() string {
	switch chain {
	case EthMain, Ropsten, Rinkeby, Goerli, Kovan, Ganache:
		return "ETH"
	case BscMain, BscTest:
		return "BNB"
	case HecoMain, HecoTest:
		return "HT"
	case PolygonMain, Mumbai:
		return "MATIC"
	case OkcTest, OkcMain:
		return "OKT"
	default:
		return ""
	}
}

// Explorer of the chain id
func (chain ChainId) Explorer() string {
	switch chain {
	case EthMain:
		return "https://etherscan.io/"
	case Ropsten:
		return "https://ropsten.etherscan.io/"
	case Rinkeby:
		return "https://rinkeby.etherscan.io/"
	case Goerli:
		return "https://goerli.etherscan.io/"
	case Kovan:
		return "https://kovan.etherscan.io/"
	case BscMain:
		return "https://bscscan.com/"
	case BscTest:
		return "https://testnet.bscscan.com/"
	case HecoMain:
		return "https://hecoinfo.com/"
	case HecoTest:
		return "https://testnet.hecoinfo.com/"
	case PolygonMain:
		return "https://polygonscan.com/"
	case Mumbai:
		return "https://mumbai.polygonscan.com/"
	case OkcMain:
		return "https://www.oklink.com/en/okc/"
	case OkcTest:
		return "https://www.oklink.com/en/okc-test/"
	default:
		return ""
	}
}

// IsTestnet ...
func (chain ChainId) IsTestnet() bool {
	switch chain {
	case Ropsten, Rinkeby, Goerli, Kovan, Ganache, BscTest, HecoTest, Mumbai:
		return true
	default:
		return false
	}
}
