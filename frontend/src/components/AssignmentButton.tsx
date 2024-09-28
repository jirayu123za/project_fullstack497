import React from "react";

interface AssignmentButtonProps {
  text: string;
  color: string;
}

export default function AssignmentButton({
  text,
  color,
}: AssignmentButtonProps) {
  const colorClass =
    {
      red: "bg-[#F80202]",
      green: "bg-[#93B955]",
      yellow: "bg-[#FFD45E]",
    }[color] || "bg-gray-500";

  return (
    <div>
      <button
        className={`font-semibold text-white w-32 h-8 rounded-xl ${colorClass} hover:opacity-75`}
      >
        {text}
      </button>
    </div>
  );
}
