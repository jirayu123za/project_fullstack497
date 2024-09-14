import React from "react";
import { FaChild } from "react-icons/fa"; 

interface BoardProps {
  name: string;
}

const LeftMain: React.FC<BoardProps> = ({ name }) => {
  return (
    <div className="flex items-center p-4 h-[900px] flex-col justify-start">
      <div className="flex items-center w-full mb-2 border-b">
        <FaChild className="text-black-500 mr-2" size={20} />
        <span className="text-gray-800">Welcome, {name}</span>
      </div>
    </div>
  );
};

export default LeftMain;
