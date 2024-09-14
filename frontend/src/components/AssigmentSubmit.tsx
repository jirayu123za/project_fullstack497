import React, { useState } from "react";

interface AssigmentSubmitProps {
  StdCode: string;
  Status: string;
}

export default function AssigmentSubmit({
  StdCode,
  Status,
}: AssigmentSubmitProps) {
  const [isHovered, setIsHovered] = useState(false);

  const [isVisible, setIsVisible] = useState(true);

  const handleClick = () => {
    setIsVisible(!isVisible);
  };

  return (
    <div className="flex justify-center items-center">
      <div
        className="flex gap-3 items-center"
        onMouseEnter={() => setIsHovered(true)}
        onMouseLeave={() => setIsHovered(false)}
      >
        {isHovered && (
          <div onClick={handleClick}>
            <svg
              width="29"
              height="29"
              viewBox="0 0 29 29"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M20.3906 8.60938L8.60938 20.3906M8.60938 8.60938L20.3906 20.3906"
                stroke="#344B59"
                stroke-width="4.96875"
                stroke-linecap="round"
                stroke-linejoin="round"
              />
            </svg>
          </div>
        )}

        <div>
          <svg
            width="29"
            height="29"
            viewBox="0 0 29 29"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <circle cx="14.5" cy="14.5" r="14.5" fill={Status} />
          </svg>
        </div>

        {isVisible && <div>{StdCode}</div>}
      </div>
    </div>
  );
}
