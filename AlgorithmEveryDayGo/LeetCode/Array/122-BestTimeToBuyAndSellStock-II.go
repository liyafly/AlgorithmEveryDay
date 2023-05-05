package leetCode

// BestTimeToBuyAndSellStockII
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfit(prices []int) int {
	// 贪心算法
	// 只要股票价格呈递增趋势，那么在每一天买进，最后一天卖出
	if len(prices) < 2 {
		return 0
	}
	profit := 0
	for i := 1; i < len(prices); i++ {
		// 如果股票价格呈递增趋势，那么在每一天买进，最后一天卖出
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			profit += diff
		}
	}
	return profit
}
