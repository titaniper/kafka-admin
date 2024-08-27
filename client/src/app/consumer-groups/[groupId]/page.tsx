'use client';

import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Link from 'next/link';
import { useParams } from 'next/navigation';
import { ConsumerGroupDetails, PartitionInfo } from '@/types/consumerGroup';

export default function ConsumerGroupDetailsPage() {
    const [groupDetails, setGroupDetails] = useState<ConsumerGroupDetails | null>(null);
    const [error, setError] = useState<string | null>(null);
    const { groupId } = useParams();

    useEffect(() => {
        fetchGroupDetails();
    }, [groupId]);

    const fetchGroupDetails = async () => {
        try {
            const response = await axios.get<ConsumerGroupDetails>(`/api/consumer-groups/${encodeURIComponent(groupId as string)}`);
            setGroupDetails(response.data);
            setError(null);
        } catch (err) {
            setError('Error fetching consumer group details');
            console.error(err);
        }
    };

    const handleResetOffset = async (partition: PartitionInfo) => {
        try {
            await axios.post('/api/consumer-groups/reset-offset', {
                groupId,
                topic: partition.topic,
                partition: partition.partition
            });
            alert(`Offset reset successfully for topic: ${partition.topic}, partition: ${partition.partition}`);
            fetchGroupDetails(); // Refresh the details after reset
        } catch (err) {
            console.error('Error resetting offset:', err);
            alert('Failed to reset offset');
        }
    };

    if (error) {
        return <div className="text-red-500">{error}</div>;
    }

    if (!groupDetails) {
        return <div className="text-gray-800">Loading...</div>;
    }

    return (
        <div className="container mx-auto p-4 bg-gray-100 min-h-screen">
            <Link href="/consumer-groups" className="text-blue-500 hover:underline mb-4 inline-block">
                &larr; Back to Consumer Groups
            </Link>
            <h1 className="text-2xl font-bold mb-4 text-gray-800">Consumer Group: {groupDetails.groupId}</h1>
            <div className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
                <p className="text-gray-800"><strong>Members:</strong> {groupDetails.members}</p>
                <p className="text-gray-800"><strong>Topics:</strong> {groupDetails.topics}</p>
                <p className="text-gray-800"><strong>State:</strong> {groupDetails.state}</p>
                <p className="text-gray-800"><strong>Partition Assignor:</strong> {groupDetails.partitionAssignor}</p>
                <p className="text-gray-800"><strong>Consumer Lag:</strong> {groupDetails.consumerLag}</p>

                <h2 className="text-xl font-bold mt-4 mb-2 text-gray-800">Coordinator</h2>
                <p className="text-gray-800"><strong>Host:</strong> {groupDetails.coordinator.host}</p>
                <p className="text-gray-800"><strong>Port:</strong> {groupDetails.coordinator.port}</p>

                <h2 className="text-xl font-bold mt-4 mb-2 text-gray-800">Partitions</h2>
                <table className="min-w-full bg-white">
                    <thead>
                    <tr>
                        <th className="py-2 px-4 border-b text-gray-800">Topic</th>
                        <th className="py-2 px-4 border-b text-gray-800">Partition</th>
                        <th className="py-2 px-4 border-b text-gray-800">Current Offset</th>
                        <th className="py-2 px-4 border-b text-gray-800">End Offset</th>
                        <th className="py-2 px-4 border-b text-gray-800">Consumer Lag</th>
                        <th className="py-2 px-4 border-b text-gray-800">Action</th>
                    </tr>
                    </thead>
                    <tbody>
                    {groupDetails.partitions.map((partition, index) => (
                        <tr key={index}>
                            <td className="py-2 px-4 border-b text-gray-800">{partition.topic}</td>
                            <td className="py-2 px-4 border-b text-gray-800">{partition.partition}</td>
                            <td className="py-2 px-4 border-b text-gray-800">{partition.currentOffset}</td>
                            <td className="py-2 px-4 border-b text-gray-800">{partition.endOffset}</td>
                            <td className="py-2 px-4 border-b text-gray-800">{partition.consumerLag}</td>
                            <td className="py-2 px-4 border-b text-gray-800">
                                <button
                                    onClick={() => handleResetOffset(partition)}
                                    className="bg-red-500 text-white px-2 py-1 rounded text-sm hover:bg-red-600"
                                >
                                    Reset Offset
                                </button>
                            </td>
                        </tr>
                    ))}
                    </tbody>
                </table>
            </div>
        </div>
    );
}