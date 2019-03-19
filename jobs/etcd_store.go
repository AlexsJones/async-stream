package jobs

import (
	"github.com/BeameryHQ/async-stream/kvstore"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func NewEtcdJobStore(cli *clientv3.Client, path string, consumerName string, runningNoUpdate time.Duration) JobStore {
	kv := kvstore.NewEtcdStore(cli)
	return NewJobStore(kv, path, consumerName, runningNoUpdate)
}
