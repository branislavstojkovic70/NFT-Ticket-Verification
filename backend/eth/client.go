package eth

var (
	infraURL   = "https://sepolia.infura.io/v3/8aaddcea52c24faeac2b2f6830528e93"
	client     *ethclient.Client
	chainID    = big.NewInt(11155111)
)
