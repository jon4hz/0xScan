package xscan

import "context"

// GetErc20TokenTransfers gets a list of "ERC-20 - Token Transfer Events" by Address
func (c *Client) GetErc20TokenTransfers(ctx context.Context, address string, opts *Erc20TokenTransfersOpts) (*Erc20TokenTransfersRes, int, error) {
	module := "account"
	action := "tokentx"
	var dataRes Erc20TokenTransfersRes

	queries := make(map[string]interface{})
	queries["address"] = address

	if opts != nil {
		if opts.StartBlock != "" {
			queries["startblock"] = opts.StartBlock
		}
		if opts.EndBlock != "" {
			queries["endblock"] = opts.EndBlock
		}
		if opts.Sort != "" {
			queries["sort"] = opts.Sort
		}
		if opts.Page != "" {
			queries["page"] = opts.Page
		}
		if opts.Offset != "" {
			queries["offset"] = opts.Offset
		}
		if opts.ContractAddress != "" {
			queries["contractaddress"] = opts.ContractAddress
		}
	}

	statusCode, err := c.doRequest(ctx, module, action, &dataRes, queries)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
