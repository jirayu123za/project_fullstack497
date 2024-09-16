import React from "react";
import { useLocation } from "react-router-dom";

interface RightMainProps {
  icons: string[];
  links: string[];
}

export default function RightMain({ icons, links }: RightMainProps) {
  const location = useLocation();
  const role = location.state?.role || "Role";

  return (
    <div>
      <div className="bg-white rounded-xl h-[900px] w-[200px] flex flex-col items-center justify-between py-8">
        <div className="flex flex-col justify-center items-center gap-10">
          {/* Profile image */}
          <div>
            <img
              className="w-28 h-28 p-2 rounded-full ring-4 ring-B1"
              src="https://img.freepik.com/free-photo/portrait-man-laughing_23-2148859448.jpg?size=338&ext=jpg&ga=GA1.1.2008272138.1725667200&semt=ais_hybrid"
              alt="Profileimage"
            />
          </div>
          {/* Icons */}
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
          className="text-white font-medium text-xl bg-E1 w-36 h-11 rounded-full"
          disabled
        >
          {role}
        </button>
      </div>
    </div>
  );
}
