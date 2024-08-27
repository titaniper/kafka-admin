export interface ConsumerGroup {
    groupId: string;
}

export interface ConsumerGroupDetails {
    inherit: string;
    groupId: string;
    members: number;
    topics: number;
    simple: boolean;
    partitionAssignor: string;
    state: string;
    coordinator: Coordinator;
    consumerLag: number;
    partitions: PartitionInfo[];
}

export interface Coordinator {
    id: number;
    host: string;
    port: number;
    bytesInPerSec: number | null;
    bytesOutPerSec: number | null;
    partitionsLeader: number | null;
    partitions: number | null;
    inSyncPartitions: number | null;
    partitionsSkew: number | null;
    leadersSkew: number | null;
}

export interface PartitionInfo {
    topic: string;
    partition: number;
    currentOffset: number;
    endOffset: number;
    consumerLag: number;
    consumerId: string | null;
    host: string | null;
}