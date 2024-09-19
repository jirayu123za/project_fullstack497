import React, { useEffect, useState } from "react";
import TitleElement from "./TitleElement";
import UpcomingAssignment from "./UpcomingAssignment";
import axios from "axios";
import upcomingIcon from "../icons/carbon_event-schedule.png";

interface Assignment {
  assignmentName: string;
  dueDate: string;
  percentage: number;
  color: string;
  timeleft: number;
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

  /*
  const [assignments, setAssignments] = useState<
    { percentage: string; color: string; timeleft: number }[]
  >([]);
  */
  const [assignments, setAssignments] = useState<Assignment[]>([]);

  useEffect(() => {
    axios
      .get("api/api/QueryAssignmentByUserID")
      .then((response) => {
        const assignmentsData = response.data.assignments.map((assignment: any) => {
          // Calculate the number of days left
          const dueDate = new Date(assignment.due_date);
          const today = new Date();
          const timeDiff = dueDate.getTime() - today.getTime();
          const daysLeft = Math.ceil(timeDiff / (1000 * 3600 * 24));

          // Determine the progress percentage (example logic)
          const percentage = Math.min((30 - daysLeft) / 30 * 100, 100);
          const color = daysLeft <= 3 ? "red" : daysLeft <= 7 ? "yellow" : "green";
          console.log("Hello" ,response.data);

          return {
            assignmentName: assignment.assignment_name,
            dueDate: assignment.due_date,
            percentage: percentage,
            color: color,
            timeleft: daysLeft,
          };
        });

        // Sort assignments by time left
        const sortedAssignments = assignmentsData.sort(
          (a: Assignment, b: Assignment) => a.timeleft - b.timeleft
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
                assignmentName={assignment.assignmentName} 
                percentage={assignment.percentage.toString()}
                color={assignment.color}
                timeleft={assignment.timeleft}
            />
          ))}
        </div>
      </div>
    </div>
  );
}