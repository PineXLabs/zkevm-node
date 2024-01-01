package polygonzkevm

import "github.com/ethereum/go-ethereum/crypto"

// ToDo: parse from config file as a future improvement
const RollupMode = "validium"

func convert2ValidiumBatchData(batches []PolygonZkEVMBatchData) []PolygonZkEVMBatchDataValidium {
	var Validiumbatches []PolygonZkEVMBatchDataValidium
	for _, batchData := range batches {
		Validiumbatches = append(Validiumbatches, PolygonZkEVMBatchDataValidium{
			TransactionsHash:   crypto.Keccak256Hash(batchData.Transactions),
			GlobalExitRoot:     batchData.GlobalExitRoot,
			Timestamp:          batchData.Timestamp,
			MinForcedTimestamp: batchData.MinForcedTimestamp,
		})
	}
	return Validiumbatches
}

func convert2ValidiumForcedBatchData(forceBatches []PolygonZkEVMForcedBatchData) []PolygonZkEVMForcedBatchDataValidium {
	var Validiumbatches []PolygonZkEVMForcedBatchDataValidium
	for _, forceBatch := range forceBatches {
		Validiumbatches = append(Validiumbatches, PolygonZkEVMForcedBatchDataValidium{
			TransactionsHash:   crypto.Keccak256Hash(forceBatch.Transactions),
			GlobalExitRoot:     forceBatch.GlobalExitRoot,
			MinForcedTimestamp: forceBatch.MinForcedTimestamp,
		})
	}
	return Validiumbatches
}
