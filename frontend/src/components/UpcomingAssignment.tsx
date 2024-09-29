import React from "react";
import TitleElement from "./TitleElement";
import upcomingIcon from "../icons/carbon_event-schedule.png";

interface UpcomingAssignment {
  assignment_id: string;
  assignment_name: string;
  assignment_due_date: string;
  color: string;
}

interface UpcomingAssignmentProps {
  UpcomingAssignment: UpcomingAssignment[];
}

const getColorClass = (color: string) => {
  switch (color.toLowerCase()) {
    case "purple":
      return "bg-purple-300 text-purple-800";
    case "yellow":
      return "bg-yellow-300 text-yellow-800";
    case "green":
      return "bg-green-300 text-green-800";
    case "red":
      return "bg-red-300 text-red-800";
    case "pink":
      return "bg-pink-300 text-pink-800";
    case "blue":
      return "bg-blue-300 text-blue-800";
    case "orange":
      return "bg-orange-300 text-orange-800";
    case "brown":
      return "bg-brown-300 text-brown-800";
    default:
      return "bg-gray-300 text-gray-800";
  }
};

const UpcomingAssignment: React.FC<UpcomingAssignmentProps> = ({ UpcomingAssignment }) => {
  return (
    <div className="p-4 overflow-hidden font-poppins text-E1">
      <div className="mb-4">
        <TitleElement name="Upcoming Assignment" icon={upcomingIcon} />
      </div>
      <div className="max-h-80 overflow-y-auto"> {/* เพิ่ม class max-h และ overflow-y-auto */}
        {UpcomingAssignment.map((upcomingAssignment) => {
          const color = upcomingAssignment.color || "gray";

          return (
            <div key={upcomingAssignment.assignment_id} className="p-2 font-poppins">
              <div className="flex flex-col w-full max-w-full gap-3 font-poppins text-sm">
                <div className="text-base sm:text-lg font-base">{upcomingAssignment.assignment_name}</div>
                {/* Progress Bar */}
                <div className="w-full bg-white rounded-full overflow-hidden h-[23px] border border-[#D9D9D9]">
                  <div
                    className={`${getColorClass(color)} bg-opacity-60 text-end py-1 px-2 rounded-full font-semibold pr-3 h-full flex items-center justify-end`}
                    style={{ width: "50%" }} // Adjust this based on actual progress
                  >
                    <div className="text-xs sm:text-sm">{upcomingAssignment.assignment_due_date} days left</div>
                  </div>
                </div>
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
};


export default UpcomingAssignment;