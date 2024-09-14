import React from "react";
import AssignmentSubmit from "./AssigmentSubmit";

interface AssignmentSubmittedProps {
  ConfigAssignment: { StdCode: string; Status: string }[];
}

export default function AssignmentSubmitted({
  ConfigAssignment,
}: AssignmentSubmittedProps) {
  return (
    <div className="flex justify-center items-center h-screen">
      <div className="bg-white border-2 border-B1 w-[214px] h-[528px] rounded-xl p-6 flex flex-col space-y-5 overflow-y-auto">
        {ConfigAssignment.map((config, index) => (
          <AssignmentSubmit
            key={index}
            StdCode={config.StdCode}
            Status={config.Status}
          />
        ))}
      </div>
    </div>
  );
}
