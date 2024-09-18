import React, { useState } from "react";
import { useLocation } from "react-router-dom";

interface HeaderProps {
  icon: string;
  title: string;
}

const LeftMain: React.FC<HeaderProps> = ({ icon, title }) => {
  const location = useLocation();
  const isDashboard = location.pathname.includes("/dashboard");
  return (
    <div className="flex items-center p-4 flex-col justify-start">
      <div className="flex flex-col sm:flex-row items-center w-full mb-2 border-b py-4  border-B1 font-poppins font-medium text-[32px] text-E1 gap-3">
        <img src={icon} alt="icon" width={53} height={53} />
        <h2 className="text-3xl font-medium text-344B59">
          {isDashboard ? `Welcome , ${title}` : title}
        </h2>
      </div>
    </div>
  );
};

export default LeftMain;
