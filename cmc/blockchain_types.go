package cmc

import "time"

type BlockchainStatisticsLatest struct {
	Data map[string]struct {
		Id                  int       `json:"id"`
		Slug                string    `json:"slug"`
		Symbol              string    `json:"symbol"`
		BlockRewardStatic   float64   `json:"block_reward_static"`
		ConsensusMechanism  string    `json:"consensus_mechanism"`
		Difficulty          string    `json:"difficulty"`
		Hashrate24H         string    `json:"hashrate_24h"`
		PendingTransactions int       `json:"pending_transactions"`
		ReductionRate       string    `json:"reduction_rate"`
		TotalBlocks         int       `json:"total_blocks"`
		TotalTransactions   string    `json:"total_transactions"`
		Tps24H              float64   `json:"tps_24h"`
		FirstBlockTimestamp time.Time `json:"first_block_timestamp"`
	} `json:"data"`
	Status Status `json:"status"`
}
