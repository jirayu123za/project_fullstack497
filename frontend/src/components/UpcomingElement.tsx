import React, { useEffect, useState } from "react";
import TitleElement from "./TitleElement";
import UpcomingAssignment from "./UpcomingAssignment";
import axios from "axios";
import upcomingIcon from "../icons/carbon_event-schedule.png";

interface Assignment {
  percentage: string;
  color: string;
}

export default function UpcomingElement() {
  // const assignments = [
  //   { percentage: "70%", color: "green", timeleft: 3 },
  //   { percentage: "50%", color: "purple", timeleft: 6 },
  //   { percentage: "30%", color: "yellow", timeleft: 7 },
  //   { percentage: "70%", color: "green", timeleft: 2 },
  //   { percentage: "50%", color: "purple", timeleft: 1 },
  //   { percentage: "30%", color: "yellow", timeleft: 11 },
  // ];

  const [assignments, setAssignments] = useState<
    { percentage: string; color: string; timeleft: number }[]
  >([]);

  useEffect(() => {
    axios
      .get("/upcoming.json")
      .then((response) => {
        const sortedAssignments = response.data.sort(
          (a: { timeleft: number }, b: { timeleft: number }) =>
            a.timeleft - b.timeleft
        );
        setAssignments(sortedAssignments);
      })
      .catch((error) => {
        console.error("Error fetching data:", error);
      });
  }, []);

  return (
    <div className="p-4 overflow-hidden font-poppins text-E1">
      <TitleElement name="Upcoming Assignment" icon={upcomingIcon} />
      {/* Loop assignment */}
      <div className="max-h-[400px] overflow-y-auto scrollbar-hide">
        <div>
          {assignments.map((assignment, index) => (
            <UpcomingAssignment
              key={index}
              percentage={assignment.percentage}
              color={assignment.color}
              timeleft={assignment.timeleft}
            />
          ))}
        </div>
      </div>
    </div>
  );
}
