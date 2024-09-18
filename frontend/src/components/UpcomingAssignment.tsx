import React from "react";

interface UpcomingAssignmentProps {
  percentage: string;
  color: string;
  timeleft: number;
}

export default function UpcomingAssignment({
  percentage,
  color,
  timeleft,
}: UpcomingAssignmentProps) {
  const RandomColor =
    {
      purple: "bg-R1",
      yellow: "bg-R2",
      pink: "bg-R3",
      green: "bg-R4",
    }[color] || "bg-gray-500";

  return (
    <div className="p-2 font-poppins">
      <div className="flex flex-col w-full max-w-full gap-3 font-poppins text-sm">
        <div className="text-base sm:text-lg font-base">Assignment 1</div>
        {/* Loading */}
        <div className="w-full bg-white rounded-full overflow-hidden h-[23px] border border-[#D9D9D9]">
          <div
            className={`${RandomColor} bg-opacity-60 text-white text-end py-1 px-2 rounded-full font-semibold pr-3 h-full flex items-center justify-end`}
            style={{ width: `${percentage}` }}
          >
            <div className="text-xs sm:text-sm">{timeleft} days left</div>
          </div>
        </div>
      </div>
    </div>
  );
}
