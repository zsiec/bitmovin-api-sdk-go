package model
import (
	"time"
)

type KubernetesCluster struct {
	// Name of the resource. Can be freely chosen by the user.
	Name string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description string `json:"description,omitempty"`
	// Creation timestamp expressed in UTC: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Modified timestamp expressed in UTC: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource
	Id string `json:"id,omitempty"`
	// Shows if the Bitmovin Agent is alive
	Online *bool `json:"online,omitempty"`
	// Shows if the Kubernetes cluster is accessible by the Bitmovin Agent
	Connected *bool `json:"connected,omitempty"`
	AgentDeploymentDownloadUrl string `json:"agentDeploymentDownloadUrl,omitempty"`
}

