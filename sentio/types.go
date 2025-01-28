package sentio

type SentioProjectConfig struct {
	StartTimestamp int64
	EndTimestamp   int64
	// sentio `point_update` table uses different naming conventions for LBTC balance depending on the project
	LbtcBalanceFieldName string
}
