import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { LuUsers } from 'react-icons/lu'; 
import { FiX } from 'react-icons/fi'; 

interface Friend {
  id: number;
  name: string;
  image: string;
}

const FriendList: React.FC = () => {
  const [friends, setFriends] = useState<Friend[]>([]);

  useEffect(() => {
    axios.get('http://localhost:3000/api/friends')
      .then(response => {
        setFriends(response.data);
      })
      .catch(error => {
        console.error('Error fetching friends:', error);
      });
  }, []);


  const handleRemove = (id: number) => {
    setFriends(friends.filter(friend => friend.id !== id));
  };

  return (
    <div className="p-4 max-w-full">

      <div className="flex items-center mb-4">
        <LuUsers className="mr-2 text-lg" />
        <h2 className="text-lg font-semibold">Friend List</h2>
      </div>
      <div className="overflow-y-auto max-h-60 scrollbar-hide w-1/5"> 
        <ul className="divide-y divide-gray-200">
          {friends.map((friend) => (
            <li key={friend.id} className="flex items-center justify-between p-2">
              <div className="flex items-center">
                <img
                  src={friend.image}
                  alt={friend.name}
                  className="w-10 h-10 rounded-full mr-3"
                />
                <p className="text-sm">{friend.name}</p>
              </div>
              <FiX
                className="text-red-500 cursor-pointer"
                onClick={() => handleRemove(friend.id)}
              />
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default FriendList;
