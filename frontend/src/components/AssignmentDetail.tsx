import React from "react";

interface AssignmentDetailProps {
  role: string;
}

export default function AssignmentDetail({ role }: AssignmentDetailProps) {
  return (
    <textarea
      className="container border-2 border-B1 min-w-full min-h-[500px] rounded-lg p-6 text-M1"
      placeholder="Assignment details..."
      disabled={role === "student"}
    />
  );
}
