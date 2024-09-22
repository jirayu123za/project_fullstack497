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
  const [isPopupOpen, setIsPopupOpen] = useState(false); // สำหรับควบคุม popup
  const [email, setEmail] = useState(""); // สำหรับ input email
  const [errorMessage, setErrorMessage] = useState<string | null>(null);

  const fetchFriends = () => {
    axios
      .get("friend.json") // ควรเปลี่ยนให้เป็น API endpoint ที่ดึงข้อมูลรายชื่อนักเรียน
      .then((response) => {
        setFriends(response.data);
      })
      .catch((error) => {
        console.error("Error fetching friends:", error);
      });
  };

  useEffect(() => {
    fetchFriends(); // ดึงข้อมูลนักเรียนเมื่อ component mount
  }, []);

  const handleRemove = (id: number) => {
    setFriends(friends.filter((friend) => friend.id !== id));
  };

  const handleAddFriend = async () => {
    try {
      // เรียก API เพื่อเช็ค email ใน database
      const response = await axios.post("/api/checkEmail", { email });
      const { data } = response;

      if (data.exists) {
        // หาก email มีใน database จะเรียก API เพื่อเพิ่มนักเรียนใหม่
        const addResponse = await axios.post("/api/addFriend", { email });
        if (addResponse.status === 200) {
          // หลังจากเพิ่มสำเร็จ ให้ fetch รายชื่อนักเรียนมาใหม่
          fetchFriends();
          setIsPopupOpen(false); // ปิด popup
          setEmail(""); // ล้างค่า input
          setErrorMessage(null); // ล้าง error message
        } else {
          setErrorMessage("Failed to add friend.");
        }
      } else {
        setErrorMessage("Email not found in database.");
      }
    } catch (error) {
      console.error("Error adding friend:", error);
      setErrorMessage("Error adding friend.");
    }
  };

  return (
    <div className="max-w-full">
      <div className="flex items-center mb-4 gap-3">
        <TitleElement name={"Student List"} icon={stdicon} />
        <div onClick={() => setIsPopupOpen(true)}>
          <FaPlusSquare size={40} color="#93B955" />
        </div>
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

      {/* Popup สำหรับการเพิ่มนักเรียน */}
      {isPopupOpen && (
        <div className="fixed inset-0 bg-gray-600 bg-opacity-75 flex justify-center items-center z-50">
          <div className="bg-white p-8 rounded-lg shadow-lg w-full max-w-md">
            <h2 className="text-xl mb-4">Add New Student</h2>
            <input
              type="email"
              className="border p-2 w-full mb-4"
              placeholder="Enter student email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
            {errorMessage && <p className="text-red-500">{errorMessage}</p>}
            <div className="flex justify-end gap-2">
              <button
                className="bg-red-500 text-white px-4 py-2 rounded"
                onClick={() => setIsPopupOpen(false)}
              >
                Cancel
              </button>
              <button
                className="bg-R4 text-white px-4 py-2 rounded"
                onClick={handleAddFriend}
              >
                Add Student
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default FriendList;
