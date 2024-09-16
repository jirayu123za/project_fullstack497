import React from "react";

interface HeaderProps {
  icon: string;
  title: string;
}

const LeftMain: React.FC<HeaderProps> = ({ icon, title }) => {
  return (
    <div className="flex items-center p-4 flex-col justify-start">
      <div className="flex items-center w-full mb-2 border-b py-4  border-B1 font-poppins font-medium text-[32px] text-E1 gap-3">
        <img src={icon} alt="icon" width={53} height={53} />
        <h2 className="text-3xl font-medium text-344B59">Welcome ,{title}</h2>
      </div>
    </div>
  );
};

export default LeftMain;
