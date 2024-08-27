package models

type ConsumerGroupDetailsResponse struct {
	Inherit           string          `json:"inherit"`
	GroupID           string          `json:"groupId"`
	Members           int             `json:"members"`
	Topics            int             `json:"topics"`
	Simple            bool            `json:"simple"`
	PartitionAssignor string          `json:"partitionAssignor"`
	State             string          `json:"state"`
	Coordinator       Coordinator     `json:"coordinator"`
	ConsumerLag       int64           `json:"consumerLag"`
	Partitions        []PartitionInfo `json:"partitions"`
}

type Coordinator struct {
	ID               int      `json:"id"`
	Host             string   `json:"host"`
	Port             int      `json:"port"`
	BytesInPerSec    *float64 `json:"bytesInPerSec"`
	BytesOutPerSec   *float64 `json:"bytesOutPerSec"`
	PartitionsLeader *int     `json:"partitionsLeader"`
	Partitions       *int     `json:"partitions"`
	InSyncPartitions *int     `json:"inSyncPartitions"`
	PartitionsSkew   *float64 `json:"partitionsSkew"`
	LeadersSkew      *float64 `json:"leadersSkew"`
}

type PartitionInfo struct {
	Topic         string  `json:"topic"`
	Partition     int32   `json:"partition"`
	CurrentOffset int64   `json:"currentOffset"`
	EndOffset     int64   `json:"endOffset"`
	ConsumerLag   int64   `json:"consumerLag"`
	ConsumerID    *string `json:"consumerId"`
	Host          *string `json:"host"`
}
