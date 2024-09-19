import { useState, useEffect } from "react";
import { useLocation } from "react-router-dom";
import axios from "axios";

interface RightMainProps {
  icons: string[];
  links: string[];
}

export default function RightMain({ icons, links }: RightMainProps) {
  //const location = useLocation();
  //const role = location.state?.role || localStorage.getItem("role") || "Role";
  const [role, setRole] = useState("");

  const handleLogout = async () => {
    try {
      const response = await fetch("api/logout", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
      });

      if (response.ok) {
        console.log("Logout successful");
        console.log(response);
        window.location.href = "/landing";
      } else {
        console.error("Logout failed:", response.statusText);
      }
    } catch (error) {
      console.error("Error during logout:", error);
    }
  };

  useEffect(() => {
    axios
      .get("api/api/QueryUserGroupByUserID")
      .then((response) => response.data)
      .then((data) => {
        console.log("Fetched user group:", data);
        setRole(data.group_name);
      })
      .catch((error) => {
        console.error("Error fetching user group:", error);
      });
  }, []);

  return (
    <div className="bg-white rounded-xl w-[150px] h-full">
      <div className="flex flex-col items-center h-full justify-between py-6">
        <div className="flex flex-col justify-center items-center gap-10 ">
          <div>
            <img
              className="w-28 h-28 p-2 rounded-full ring-4 ring-B1"
              src="https://img.freepik.com/free-photo/portrait-man-laughing_23-2148859448.jpg?size=338&ext=jpg&ga=GA1.1.2008272138.1725667200&semt=ais_hybrid"
              alt="Profile image"
            />
          </div>

          <div className="flex flex-col gap-10 w-fit">
            {icons.map((icon, index) => (
              <a
                href={links[index]}
                key={index}
                className="inline-block"
                onClick={index === icons.length - 1 ? handleLogout : undefined}
              >
                <img
                  src={icon}
                  alt={`icon-${index}`}
                  className="w-14 h-14 inline-block"
                />
              </a>
            ))}
          </div>
        </div>
        {/* Role button */}
        <button
          className="text-white font-medium text-xl bg-E1 px-4 py-2 rounded-full w-[130px]"
          disabled
        >
          {role}
        </button>
      </div>
    </div>
  );
}
