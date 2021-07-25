package xscan

import "net/http"

type Client struct {
	endpoint string
	apiKey   string
	http     *http.Client
}

type Erc20TokenTransfersOpts struct {
	// startblock: starting blockNo to retrieve results
	StartBlock string
	// endblock: ending blockNo to retrieve results
	EndBlock string
	// sort order, asc or desc
	Sort string
	// To get paginated results use page=<page number>
	Page string
	// offset=<max records to return>
	Offset string
	// To get transfer events for a specific token contract, include the contractaddress paramete
	ContractAddress string
}

type Erc20TokenTransfersRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []Tx   `json:"result"`
}

type Tx struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	From              string `json:"from"`
	ContractAddress   string `json:"contractAddress"`
	To                string `json:"to"`
	Value             string `json:"value"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenDecimal      string `json:"tokenDecimal"`
	TransactionIndex  string `json:"transactionIndex"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
	Confirmations     string `json:"confirmations"`
}
