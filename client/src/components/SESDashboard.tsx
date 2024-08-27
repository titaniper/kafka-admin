'use client';

import { useState, useEffect } from 'react';
import axios from 'axios';
import Link from 'next/link';

export default function ConsumerGroups() {
  const [consumerGroups, setConsumerGroups] = useState([]);
  const [keyword, setKeyword] = useState('');
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchConsumerGroups();
  }, []);

  const fetchConsumerGroups = async (searchKeyword = '') => {
    try {
      const response = await axios.get('/api/consumer-groups', {
        params: { keyword: searchKeyword }
      });
      setConsumerGroups(response.data.data);
      setError(null);
    } catch (err) {
      setError('Error fetching consumer groups');
      console.error(err);
    }
  };

  const handleSearch = (e) => {
    e.preventDefault();
    fetchConsumerGroups(keyword);
  };

  return (
      <div className="container mx-auto p-4 bg-gray-100 min-h-screen">
        <h1 className="text-2xl font-bold mb-4 text-gray-800">Kafka Consumer Groups</h1>

        <form onSubmit={handleSearch} className="mb-4">
          <input
              type="text"
              value={keyword}
              onChange={(e) => setKeyword(e.target.value)}
              placeholder="Filter consumer groups"
              className="border p-2 mr-2 rounded text-gray-800 bg-white"
          />
          <button type="submit" className="bg-blue-500 text-white p-2 rounded hover:bg-blue-600">
            Search
          </button>
        </form>

        {error && <p className="text-red-500 mb-4">{error}</p>}

        <ul className="list-disc pl-5 text-gray-800">
          {consumerGroups.map((group, index) => (
              <li key={index} className="mb-1">
                <Link href={`/consumer-groups/${encodeURIComponent(group)}`} className="text-blue-500 hover:underline">
                  {group}
                </Link>
              </li>
          ))}
        </ul>
      </div>
  );
}