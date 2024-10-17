import React, { useState, useEffect } from "react";
import axios from "axios";

// interface ProgressBarCourseProps {
//   courseId: string;
// }

interface ProgressBarCourseProps {
  course: { assignments: Array<{ completed: boolean }> }; // กำหนดประเภทของ props
}

const ProgressBarCourse: React.FC<ProgressBarCourseProps> = ({ course }) => {
  const [totalAssignments, setTotalAssignments] = useState<number>(0);
  const [completedAssignments, setCompletedAssignments] = useState<number>(0);

  // useEffect(() => {
  //   const fetchAssignments = async () => {
  //     try {
  //       const response = await axios.get(
  //         `http://localhost:3000/api/courses/${courseId}/assignments`
  //       );
  //       const courseData = response.data;

  //       // เข้าถึง assignments
  //       const assignments = courseData.assignments || [];

  //       setTotalAssignments(assignments.length);

  //       const completed = assignments.filter(
  //         (assignment: any) => assignment.completed
  //       ).length;

  //       setCompletedAssignments(completed);
  //     } catch (error) {
  //       console.error("Error fetching assignments:", error);
  //     }
  //   };

  //   fetchAssignments();
  // }, []);

  useEffect(() => {
    console.log(course);
    if (course?.assignments) {
      const assignments = course.assignments;

      setTotalAssignments(assignments.length);
      console.log(assignments);
      const completed = assignments.filter(
        (assignment) => assignment.complete === true
      ).length;
      console.log(completed);
      setCompletedAssignments(completed);
    }
  }, [course]);

  const progressPercentage =
    totalAssignments > 0 ? (completedAssignments / totalAssignments) * 100 : 0;

  const getColorClass = (color: string) => {
    switch (color.toLowerCase()) {
      case "purple":
        return "border-purple-300 ";
      case "yellow":
        return "border-yellow-300 ";
      case "green":
        return "border-green-300 ";
      case "red":
        return "border-red-300 ";
      case "pink":
        return "border-pink-300 ";
      case "blue":
        return "border-blue-300 ";
      case "orange":
        return "border-orange-300 ";
      case "brown":
        return "border-brown-300 ";
      default:
        return "border-gray-300";
    }
  };

  const getBgColorClass = (color: string) => {
    switch (color.toLowerCase()) {
      case "purple":
        return "bg-purple-300 ";
      case "yellow":
        return "bg-yellow-300 ";
      case "green":
        return "bg-green-300 ";
      case "red":
        return "bg-red-300 ";
      case "pink":
        return "bg-pink-300 ";
      case "blue":
        return "bg-blue-300 ";
      case "orange":
        return "bg-orange-300 ";
      case "brown":
        return "bg-brown-300 ";
      default:
        return "bg-gray-300";
    }
  };

  return (
    <div
      className={`relative w-full h-8 bg-white border-2 font-poppins ${getColorClass(
        course.color
      )} rounded-full`}
    >
      <div
        className={`absolute top-0 left-0 h-full ${getBgColorClass(
          course.color
        )} bg-opacity-50 rounded-full flex items-center `}
        style={{ width: `${progressPercentage}%` }}
      >
        <div
          className={`w-6 h-6 ${getBgColorClass(
            course.color
          )} rounded-full absolute right-1 ${
            progressPercentage === 0 ? "left-1" : ""
          }`}
        ></div>
        <span
          className={`text-sm font-semibold text-gray-800 ml-2 absolute right-7 ${
            progressPercentage === 0 ? "left-7" : ""
          }`}
        >{`${Math.round(progressPercentage)}%`}</span>
      </div>
    </div>
  );
};

export default ProgressBarCourse;
