package services

import (
	"github.com/smartcontractkit/chainlink/store"
	"github.com/smartcontractkit/chainlink/store/models"
)

func ExportedExecuteRunAtBlock(
	run models.JobRun,
	store *store.Store,
	overrides models.RunResult,
	blockNumber *models.IndexableBlockNumber,
) (models.JobRun, error) {
	run.Overrides = overrides
	return executeRunAtBlock(run, store, blockNumber)
}

func ExportedChannelForRun(jr JobRunner, runID string) chan<- store.RunRequest {
	return jr.channelForRun(runID)
}

func ExportedResumeSleepingRuns(jr JobRunner) error {
	return jr.resumeSleepingRuns()
}

func ExportedWorkerCount(jr JobRunner) int {
	return jr.workerCount()
}
