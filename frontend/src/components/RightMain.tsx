import React from "react";
import { useLocation } from "react-router-dom";

interface RightMainProps {
  icons: string[];
  links: string[];
  profileimage: string;
}

export default function RightMain({
  icons,
  links,
  profileimage,
}: RightMainProps) {
  const location = useLocation();
  const role = location.state?.role || "Role";

  return (
    <div className="bg-white rounded-xl w-[150px] h-full">
      <div className="flex flex-col items-center h-full justify-between py-6">
        <div className="flex flex-col justify-center items-center gap-10 ">
          <div>
            <img
              className="w-28 h-28 p-2 rounded-full ring-4 ring-B1 object-cover"
              src={profileimage}
              alt="Profile image"
            />
          </div>

          <div className="flex flex-col gap-10 w-fit">
            {icons.map((icon, index) => (
              <a href={links[index]} key={index} className="inline-block">
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