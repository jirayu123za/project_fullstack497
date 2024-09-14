import React from 'react';
import { LuBox } from 'react-icons/lu'; 

const CourseTitle: React.FC = () => {
  return (
    <div className="flex items-center text-black mb-4">
      <LuBox className="mr-2" size={24} /> 
      <h2 className="text-xl font-bold">Course</h2>
    </div>
  );
};

export default CourseTitle;
