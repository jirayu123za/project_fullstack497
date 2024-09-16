import React from "react";
import { FaChild } from "react-icons/fa";

interface BoardProps {
  name: string;
}

const LeftMain: React.FC<BoardProps> = ({ name }) => {
  return (
    <div className="flex items-center p-4 flex-col justify-start">
      <div className="flex items-center w-full mb-2 border-b py-4  border-B1 font-poppins font-medium text-[32px] text-E1">
        <FaChild className=" mr-2" size={50} />
        <span>Welcome, {name}</span>
      </div>
    </div>
  );
};

export default LeftMain;
