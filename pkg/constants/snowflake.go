package constants

const (
	DefaultDataCenterID = int64(0)
	DefaultWorkerID     = int64(0)
)

const (
	WorkerOfUserService = 1 + int64(iota)
	WorkerOfOrderService
)
