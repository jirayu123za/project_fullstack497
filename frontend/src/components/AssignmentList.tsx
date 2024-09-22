import React, { useState, useEffect } from "react";
import { MdOutlineFileDownload, MdDateRange } from "react-icons/md";
import { useNavigate } from "react-router-dom";
import TitleElement from "./TitleElement";
import Assicon from "../icons/ion_list.png";
import axios from "axios";

interface Assignment {
  assignment_id: string;
  title: string;
  due_date: string;
}

interface Course {
  course_id: string;
  course_name: string;
  color: string;
  assignments: Assignment[];
}

interface AssignmentListProps {
  courses: Course[];
}

const AssignmentList: React.FC<AssignmentListProps> = ({ courses }) => {
  const navigate = useNavigate();

  const handleClick = async (assignmentId: string) => {
    try {
      await axios.post("/api/assignments", { assignment_id: assignmentId });
      navigate("/assignment");
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
      <div className="flex items-center mb-4">
        <TitleElement name="Assignment" icon={Assicon} />
      </div>
      <div className="max-h-[400px] overflow-y-scroll scrollbar-hide">
        {courses.map((course) => {
          const colorClasses = getColorClass(course.color); // Get color for border and icons

          return (
            <div key={course.course_id}>
              <h2
                className={`${colorClasses.titlecolor} text-xl font-bold mb-2`}
              >
                {course.course_name}
              </h2>
              {course.assignments.map((assignment) => (
                <div
                  key={assignment.assignment_id}
                  onClick={() => handleClick(assignment.assignment_id)} // Handle assignment click
                  className={`flex items-center justify-between border-4 p-2.5 mb-2 rounded-lg shadow-sm h-[87px] cursor-pointer ${colorClasses.borderColor}`}
                >
                  <p className="text-xl">{assignment.title}</p>
                  <div className="flex space-x-3">
                    <MdDateRange
                      className={`cursor-pointer ${colorClasses.iconColor}`}
                      size={25}
                    />
                    <MdOutlineFileDownload
                      className={`cursor-pointer ${colorClasses.iconColor}`}
                      size={25}
                    />
                  </div>
                </div>
              ))}
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default AssignmentList;
