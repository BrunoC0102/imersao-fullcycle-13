package transformer

import (
	"github.com/BrunoC0102/imersao-fullcycle-13/internal/market/dto"
	"github.com/BrunoC0102/imersao-fullcycle-13/internal/market/entity"
)

func TransformInput(input dto.TradeInput) *entity.Order {
	asset := entity.NewAsset(input.AssetId, input.AssetId, 1000)
	investor := entity.NewInvestor(input.InvestorId)
	order := entity.NewOrder(input.OrderId, investor, asset, input.Shares, input.Price, input.OrderType)
	if input.CurrentShares > 0 {
		assetPosition := entity.NewInvestorAssetPosition(input.AssetId, input.CurrentShares)
		investor.AddAssetPosition(assetPosition)
	}
	return order
}

func TransformOutput(order *entity.Order) *dto.OrderOutput {
	output := &dto.OrderOutput{
		OrderId:    order.ID,
		InvestorId: order.Investor.ID,
		AssetId:    order.Asset.ID,
		OrderType:  order.OrderType,
		Status:     order.Status,
		Partial:    order.PendingShares,
		Shares:     order.Shares,
	}

	var transactionsOutput []*dto.TransactionOutput
	for _, t := range order.Transactions {
		transactionOutput := &dto.TransactionOutput{
			TransactionId: t.ID,
			BuyerId:       t.BuyingOrder.ID,
			SellerId:      t.SellingOrder.ID,
			AssetId:       t.SellingOrder.Asset.ID,
			Price:         t.Price,
			Shares:        t.Shares,
		}
		transactionsOutput = append(transactionsOutput, transactionOutput)
	}
	output.TransactionOutput = transactionsOutput
	return output
}
