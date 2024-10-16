import React from "react";

interface TitleElementProps {
  name: string;
  icon: string;
}

const TitleElement: React.FC<TitleElementProps> = ({ name, icon }) => {
  return (
    <div className="flex items-center text-E1 gap-3 ">
      <img src={icon} alt="icon" />
      <h2 className="text-xl whitespace-nowrap">{name}</h2>
    </div>
  );
};

export default TitleElement;
