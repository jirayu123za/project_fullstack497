import React, { useState, useEffect } from "react";
import axios from "axios";
import { useLocation } from "react-router-dom";

interface HeaderProps {
  icon: string;
  title: string;
}

const LeftMain: React.FC<HeaderProps> = ({ icon, title }) => {
  const [name, setName] = useState("");

  useEffect(() => {
    axios
      .get("api/api/QueryNameByUserID")
      .then((response) => response.data)
      .then((data) => {
        console.log("Fetched courses:", data);
        setName(data.name);
      })
      .catch((error) => {
        console.error("Error fetching courses:", error);
      });
  }, []);

  const location = useLocation();
  const isDashboard =
    location.pathname.includes("/insdash") ||
    location.pathname.includes("/stddash");
  return (
    <div className="flex items-center p-4 flex-col justify-start">
       <div className="flex flex-col sm:flex-row items-center w-full mb-2 border-b py-4  border-B1 font-poppins font-medium text-[32px] text-E1 gap-3">
        <img src={icon} alt="icon" width={53} height={53} />
        <h2 className="text-3xl font-medium text-344B59"> {isDashboard ? `Welcome , ${title}` : title}</h2>
        <h2 className="text-3xl font-medium text-344B59">{name}</h2>
      </div>
    </div>
  );
};

export default LeftMain;
