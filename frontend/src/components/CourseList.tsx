import React, { useState, useEffect } from "react";
import axios from "axios";
import TitleElement from "./TitleElement";

interface Course {
  id: number;
  name: string;
  code: string;
  assignmentsCount: number;
  imageUrl: string;
  colorClass: string;
}

const CourseList: React.FC = () => {
  const [courses, setCourses] = useState<Course[]>([]);

  useEffect(() => {
    axios
      .get("http://localhost:3000/api/courses")
      .then((response) => {
        const fetchedCourses = response.data.map((course: any, index: number) => ({
          ...course,
          colorClass: getColorClass(index),
        }));
        setCourses(fetchedCourses);
      })
      .catch((error) => {
        console.error("Error fetching courses:", error);
      });
  }, []);

  const getColorClass = (index: number) => {
    const colors = [
      "bg-yellow-200 text-yellow-800",
      "bg-green-200 text-green-800",
      "bg-pink-200 text-pink-800",
      "bg-purple-200 text-purple-800",
    ];
    return colors[index % colors.length];
  };

  return (
    <div className="p-4 max-w-full">
      <TitleElement />
      <div className="overflow-x-auto scrollbar-hide w-6/12 mt-4">
        <div className="flex space-x-4 flex-nowrap whitespace-nowrap" style={{ width: 'calc(4 * 20rem)' }}>
          {courses.map((course) => (
            <div
              key={course.id}
              className="inline-block border border-gray-200 p-5 rounded-lg shadow-md flex flex-col items-center justify-center bg-white w-72 h-3/5"
            >
              <img
                src={course.imageUrl}
                alt="Course"
                className="w-48 h-24 object-cover mb-2 rounded-md"
              />
              <p className="text-left text-sm font-medium w-full">{course.name}</p>
              <p className="text-left text-sm font-medium w-full">({course.code})</p>
              <span
                className={`block text-xs text-center mt-2 px-2 py-1 rounded-full ${course.colorClass}`}
              >
                Assignment release: {course.assignmentsCount}
              </span>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default CourseList;
