import React, { useState, useEffect } from "react";
import { FaTrash, FaDownload, FaList } from "react-icons/fa";
import axios from "axios";

interface Assignment {
  id: number;
  title: string;
  colorClass: string; 
  iconColorClass: string; 
}

const AssignmentList: React.FC = () => {
  const [assignments, setAssignments] = useState<Assignment[]>([]);

  useEffect(() => {
    axios.get("http://localhost:3000/api/assignments")
      .then((response) => {
        const fetchedAssignments = response.data.map((assignment: any, index: number) => {
          const { colorClass, iconColorClass } = getColorClasses(index);
          return {
            ...assignment,
            colorClass, 
            iconColorClass 
          };
        });
        setAssignments(fetchedAssignments);
      })
      .catch((error) => {
        console.error("Error fetching assignments:", error);
      });
  }, []);

  const getColorClasses = (index: number) => {
    const colorClasses = [
      { colorClass: "border border-purple-300", iconColorClass: "text-purple-300" },
      { colorClass: "border border-yellow-300", iconColorClass: "text-yellow-300" },
      { colorClass: "border border-green-300", iconColorClass: "text-green-300" },
      { colorClass: "border border-pink-300", iconColorClass: "text-pink-300" },
    ];
    return colorClasses[index % colorClasses.length];
  };

  return (
    <div className="p-4 max-h-72 overflow-hidden">
      <div className="flex items-center mb-4">
        <FaList className="text-gray-600 mr-2" /> 
        <h2 className="text-base font-bold">Assignment List</h2>
      </div>
      <div className="max-h-60 overflow-y-scroll scrollbar-hide">
        {assignments.slice(0, 10).map((assignment) => (
          <div
            key={assignment.id}
            className={`flex items-center justify-between border-4 p-2.5 mb-2 rounded-lg shadow-sm w-80 ${assignment.colorClass}`}
          >
            <p>{assignment.title}</p>
            <div className="flex space-x-2">
              <FaDownload className={`cursor-pointer ${assignment.iconColorClass}`} />
              <FaTrash className={`cursor-pointer ${assignment.iconColorClass}`} />
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default AssignmentList;
