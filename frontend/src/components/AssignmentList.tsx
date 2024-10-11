import React from "react";
import { MdOutlineFileDownload, MdDateRange } from "react-icons/md";
import { useNavigate } from "react-router-dom";

interface Assignment {
  assignment_id: string;
  course_id: string;
  assignment_name: string;
  assignment_due_date: string;
  color: string;
  course_name: string;
}

interface AssignmentListProps {
  Assignment: Assignment[];
  showCourseName?: boolean;
  user_group_name: string;
}

const AssignmentList: React.FC<AssignmentListProps> = ({ Assignment, showCourseName = true, user_group_name }) => {
  const navigate = useNavigate();

  const groupedAssignments = Assignment.reduce((acc, assignment) => {
    if (!acc[assignment.course_name]) {
      acc[assignment.course_name] = [];
    }
    acc[assignment.course_name].push(assignment);
    return acc;
  }, {} as Record<string, Assignment[]>);

  const handleClick = async (assignment_id: string, course_id: string) => {
    try {
      console.log("Assignment ID:", assignment_id, "course ID:", course_id, "Assignment", Assignment);
      if (user_group_name === "STUDENT") {
        navigate(`/STDcourse/${course_id}/assignment/${assignment_id}`);
      } else if (user_group_name === "INSTRUCTOR") {
        navigate(`/course/${course_id}/assignment/${assignment_id}`);
      } else {
        console.error("Invalid user group name");
      }
    } catch (error) {
      console.error("Error sending assignment ID:", error);
    }
  };

  const getColorClass = (color: string) => {
    switch (color.toLowerCase()) {
      case "purple":
        return {
          borderColor: "border-purple-200",
          iconColor: "text-purple-500",
          titlecolor: "text-purple-300",
        };
      case "yellow":
        return {
          borderColor: "border-yellow-200",
          iconColor: "text-yellow-500",
          titlecolor: "text-yellow-300",
        };
      case "green":
        return {
          borderColor: "border-green-200",
          iconColor: "text-green-500",
          titlecolor: "text-green-300",
        };
      case "red":
        return {
          borderColor: "border-red-200",
          iconColor: "text-red-500",
          titlecolor: "text-red-300",
        };
      case "pink":
        return {
          borderColor: "border-pink-200",
          iconColor: "text-pink-500",
          titlecolor: "text-pink-300",
        };
      case "blue":
        return {
          borderColor: "border-blue-200",
          iconColor: "text-blue-500",
          titlecolor: "text-blue-300",
        };
      case "orange":
        return {
          borderColor: "border-orange-200",
          iconColor: "text-orange-500",
          titlecolor: "text-orange-300",
        };
      case "brown":
        return {
          borderColor: "border-brown-200",
          iconColor: "text-brown-500",
          titlecolor: "text-brown-300",
        };
      default:
        return {
          borderColor: "border-gray-200",
          iconColor: "text-gray-500",
          titlecolor: "text-gray-300",
        };
    }
  };

  return (
    <div className="p-4 overflow-hidden font-poppins text-E1">
      <div className="max-h-[400px] overflow-y-scroll scrollbar-hide">
        {Object.entries(groupedAssignments).map(([courseName, assignments]) => (
          <div key={courseName} className="mb-4">
            {showCourseName && <h2 className="text-xl font-bold mb-2">{courseName}</h2>}
            {assignments.map((assignment, index) => {
              const colorClasses = getColorClass(assignment.color || "gray");
              return (
                <div
                  key={`${assignment.assignment_id}-${index}`}
                  onClick={() => handleClick(assignment.assignment_id, assignment.course_id)}
                  className={`flex items-center justify-between border-4 p-2.5 mb-2 rounded-lg shadow-sm h-[87px] cursor-pointer ${colorClasses.borderColor}`}
                >
                  <p className="text-xl">{assignment.assignment_name}</p>
                  <div className="flex space-x-3">
                    <MdDateRange
                      className={`cursor-pointer ${colorClasses.iconColor}`}
                      size={25}
                      onClick={() =>
                        alert(`Assignment Due Date : ${assignment.assignment_due_date}`)
                      }
                    />
                    <MdOutlineFileDownload
                      className={`cursor-pointer ${colorClasses.iconColor}`}
                      size={25}
                    />
                  </div>
                </div>
              );
            })}
          </div>
        ))}
      </div>
    </div>
  );
};

export default AssignmentList;
