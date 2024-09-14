import React from "react";

interface AssignmentDetailProps {
  role: string;
}

export default function AssignmentDetail({ role }: AssignmentDetailProps) {
  return (
    <div className="flex justify-center items-center h-screen">
      <textarea
        className="container border-2 border-B1 w-[893px] max-w-full max-h-[525px] rounded-lg p-6 text-M1"
        placeholder="Assignment details..."
        disabled={role === "student"}
      />
    </div>
  );
}
