import React from "react";
import { IconType } from "react-icons"; 

interface HeaderProps {
  icon: IconType; 
  title: string; 
}

const LeftMain: React.FC<HeaderProps> = ({ icon: Icon, title }) => {
  return (
    <div className="flex items-center p-4 flex-col justify-start">
      <div className="flex items-center w-full mb-2 border-b py-4  border-B1 font-poppins font-medium text-[32px] text-E1">
      <Icon className="mr-2 text-lg text-344B59 " size={24} />

      <h2 className="text-lg font-semibold text-344B59">{title}</h2>
      </div>
    </div>
  );
};

export default LeftMain;
