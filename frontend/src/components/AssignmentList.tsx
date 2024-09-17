import React, { useState, useEffect } from "react";
import { FaTrash, FaDownload, FaList } from "react-icons/fa";
import axios from "axios";
import TitleElement from "./TitleElement";
import Assicon from "../icons/ion_list.png";

interface Assignment {
  id: string;
  title: string;
  colorClass: string;
  iconColorClass: string;
}

const AssignmentList: React.FC = () => {
  const [assignments, setAssignments] = useState<Assignment[]>([]);

  useEffect(() => {
    axios
    .get("api/api/QueryAssignmentByUserID")
    .then((response) => response.data)
    .then((data) => {
      console.log("Fetched courses:", data);
    
        const fetchedAssignments = data.assignments.map((assignment: any, index: number) => ({
          id: assignment.AssignmentID,
          title: assignment.assignment_name,
          colorClass: getColorClasses(index),
          iconColorClass: getColorClasses(index),
        }));
        setAssignments(fetchedAssignments);
      })
      .catch((error) => {
        console.error("Error fetching assignments:", error);
      });
  }, []);

  const getColorClasses = (index: number) => {
    const colorClasses = [
      {
        colorClass: "border border-R1 border-opacity-60",
        iconColorClass: "text-R1",
      },
      {
        colorClass: "border border-R2 border-opacity-60",
        iconColorClass: "text-R2",
      },
      {
        colorClass: "border border-R4 border-opacity-60",
        iconColorClass: "text-R4",
      },
      {
        colorClass: "border border-R3 border-opacity-60",
        iconColorClass: "text-R3",
      },
    ];
    return colorClasses[index % colorClasses.length];
  };

  return (
    <div className="p-4 max-h-[402px] overflow-hidden font-poppins text-E1">
      <div className="flex items-center mb-4">
        <TitleElement name="Assignment" icon={Assicon} />
      </div>
      <div className="max-h-[350px] overflow-y-scroll scrollbar-hide">
        {assignments.slice(0, 10).map((assignment) => (
          <div
            key={assignment.id}
            className={`flex items-center justify-between border-4 p-2.5 mb-2 rounded-lg shadow-sm w-[460px] h-[87px] ${assignment.colorClass}`}
          >
            <p className="text-2xl">{assignment.title}</p>
            <div className="flex space-x-2">
              <FaDownload
                className={`cursor-pointer ${assignment.iconColorClass}`}
              />
              <FaTrash
                className={`cursor-pointer ${assignment.iconColorClass}`}
              />
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default AssignmentList;
