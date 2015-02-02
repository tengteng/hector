package algo

import (
	"core"
)

type Clustering interface {
	Init(params map[string]string)
	Cluster(dataset core.DataSet)
}
