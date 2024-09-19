import React, { useState, useEffect } from "react";
import axios from "axios";
import TitleElement from "./TitleElement";
import box from "../icons/Vector.png";

interface Course {
  id: string;
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
      .get("api/api/QueryCourseByUserID")
      .then((response) => response.data)
      .then((data) => {
        console.log("Fetched courses:", data);

        const fetchedCourses = data.courses.map((course: any, index: number) => ({
          id: course.course_id,
          name: course.course_name,
          code: course.course_code,
          assignmentsCount: course.Assignments ? course.Assignments.length : 0,
          imageUrl: "https://placehold.co/10x10",
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

  const previewbox = (
    <div className="inline-block border border-gray-200 p-5 rounded-lg shadow-md flex flex-col items-center justify-center bg-white w-72 h-[210px]">
      Please add course first
    </div>
  );

  return (
    <div className="p-4 max-w-full">
      <TitleElement name="Course" icon={box} />
      <div className="overflow-x-auto scrollbar-hide mt-4">
        <div
          className="flex space-x-4 flex-nowrap whitespace-nowrap w-32"
          // style={{ width: "calc(4 * 20rem)" }}
        >
          {courses.length === 0
            ? previewbox
            : courses.map((course) => (
                <div
                  key={course.id}
                  className="inline-block border border-gray-200 p-5 rounded-lg shadow-md flex flex-col items-center justify-center bg-white w-72 h-3/5"
                >
                  <img
                    src={course.imageUrl}
                    alt="Course"
                    className="w-48 h-24 object-cover mb-2 rounded-md"
                  />
                  <p className="text-left text-sm font-medium w-full">
                    {course.name}
                  </p>
                  <p className="text-left text-sm font-medium w-full">
                    ({course.code})
                  </p>
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
