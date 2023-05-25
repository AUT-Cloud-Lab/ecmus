package connector

import (
	"github.com/amsen20/ecmus/internal/model"
	"github.com/amsen20/ecmus/logging"
)

type Connector interface {
	FindNodes() error
	FindDeployments() error
	GetClusterState() *model.ClusterState

	Deploy(pod *model.Pod) error
	DeletePod(pod *model.Pod) (bool, error)
	WatchSchedulingEvents() (<-chan *Event, error)
}

type EventType int64

const (
	POD_CREATED EventType = iota
	POD_CHANGED
	POD_DELETED
)

type Event struct {
	EventType EventType
	Pod       *model.Pod
	Node      *model.Node
	Status    model.PodStatus
}

var log = logging.Get()
