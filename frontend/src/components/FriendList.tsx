import React, { useState, useEffect } from "react";
import axios from "axios";
import stdicon from "../icons/tdesign_member.png";
import { FiX } from "react-icons/fi";
import TitleElement from "./TitleElement";
import { FaPlusSquare } from "react-icons/fa";

interface Friend {
  id: number;
  name: string;
  image: string;
}

const FriendList: React.FC = () => {
  const [friends, setFriends] = useState<Friend[]>([]);

  useEffect(() => {
    axios
      .get("friend.json")
      .then((response) => {
        setFriends(response.data);
      })
      .catch((error) => {
        console.error("Error fetching friends:", error);
      });
  }, []);

  const handleRemove = (id: number) => {
    setFriends(friends.filter((friend) => friend.id !== id));
  };

  return (
    <div className="max-w-full">
      <div className="flex items-center mb-4 gap-3">
        <TitleElement name={"Student List"} icon={stdicon} />
        <FaPlusSquare size={40} color="#93B955" />
      </div>
      <div className="overflow-y-auto max-h-[550px] scrollbar-hide w-full">
        <ul className="divide-y divide-gray-200">
          {Array.isArray(friends) &&
            friends.map((friend) => (
              <li
                key={friend.id}
                className="flex items-center justify-between p-2"
              >
                <div className="flex items-center">
                  <img
                    src={friend.image}
                    alt={friend.name}
                    className="w-10 h-10 rounded-full mr-10 object-cover"
                  />
                  <p className="text-lg">{friend.name}</p>
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
