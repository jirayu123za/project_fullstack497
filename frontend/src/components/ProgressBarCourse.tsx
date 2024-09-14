import React, { useState, useEffect } from 'react';
import axios from 'axios';

interface ProgressBarCourseProps {
  courseId: string; 
}

const ProgressBarCourse: React.FC<ProgressBarCourseProps> = ({ courseId }) => {
  const [totalAssignments, setTotalAssignments] = useState<number>(0);
  const [completedAssignments, setCompletedAssignments] = useState<number>(0);

  useEffect(() => {
    const fetchAssignments = async () => {
      try {
        const response = await axios.get(`http://localhost:3000/api/courses/${courseId}/assignments`);
        const assignments = response.data;


        setTotalAssignments(assignments.length);
        const completed = assignments.filter((assignment: any) => assignment.completed).length;
        setCompletedAssignments(completed);
      } catch (error) {
        console.error("Error fetching assignments:", error);
      }
    };

    fetchAssignments();
  }, [courseId]);

  const progressPercentage = totalAssignments > 0 ? (completedAssignments / totalAssignments) * 100 : 0;

  return (
    <div className="relative w-full h-8 bg-white border-2 border-[#E4C1F9] rounded-full">
      <div
        className="absolute top-0 left-0 h-full bg-[#E4C1F9] bg-opacity-50 rounded-full flex items-center"
        style={{ width: `${progressPercentage}%` }}
      >
        <div className="w-6 h-6 bg-[#E4C1F9] rounded-full absolute right-4 transform translate-x-1/2 shadow-md"></div>
        <span className="text-sm font-bold text-gray-800 ml-2 absolute right-7">{`${Math.round(progressPercentage)}%`}</span>
      </div>
    </div>
  );
};

export default ProgressBarCourse;
