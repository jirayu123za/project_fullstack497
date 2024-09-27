import React, { useEffect, useState } from "react";
import TitleElement from "./TitleElement";
import UpcomingAssignment from "./UpcomingAssignment";
import upcomingIcon from "../icons/carbon_event-schedule.png";

interface Assignment {
  assignment_id: string;
  title: string;
  due_date: string;
}

interface Course {
  course_id: string;
  color: string;
  assignments: Assignment[];
}

interface UpcomingElementProps {
  courses: Course[];
}

export default function UpcomingElement({ courses }: UpcomingElementProps) {
  const [assignments, setAssignments] = useState<
    { title: string; color: string; timeleft: number }[]
  >([]);

  useEffect(() => {
    let upcomingAssignments = [];

    if (Array.isArray(courses)) {
      // หาก courses เป็น array
      upcomingAssignments = courses.flatMap((course) =>
        course.assignments.map((assignment) => {
          // Calculate time left
          const timeleft = calculateTimeLeft(assignment.due_date);
          return {
            title: assignment.title,
            color: course.color,
            timeleft: timeleft,
          };
        })
      );
    } else if (typeof courses === "object" && courses !== null) {
      // หาก courses เป็น object
      const courseAssignments = courses.assignments.map((assignment) => {
        const timeleft = calculateTimeLeft(assignment.due_date);
        return {
          title: assignment.title,
          color: courses.color,
          timeleft: timeleft,
        };
      });

      upcomingAssignments = courseAssignments;
    }

    const sortedAssignments = upcomingAssignments.sort(
      (a, b) => a.timeleft - b.timeleft
    );
    setAssignments(sortedAssignments);
  }, [courses]);

  const calculateTimeLeft = (dueDate: string) => {
    const currentDate = new Date();
    const dueDateObj = new Date(dueDate);
    const timeDiff = dueDateObj.getTime() - currentDate.getTime();
    const daysLeft = Math.ceil(timeDiff / (1000 * 3600 * 24));
    return daysLeft;
  };

  return (
    <div className="p-4 overflow-hidden font-poppins text-E1">
      <div className="mb-4">
        <TitleElement name="Upcoming Assignment" icon={upcomingIcon} />
      </div>
      {/* Loop through assignments */}
      <div className="max-h-[400px] overflow-y-auto scrollbar-hide">
        <div>
          {assignments.map((assignment, index) => (
            <UpcomingAssignment
              key={index}
              color={assignment.color}
              timeleft={assignment.timeleft}
              title={assignment.title}
            />
          ))}
        </div>
      </div>
    </div>
  );
}
