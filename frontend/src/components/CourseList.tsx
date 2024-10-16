import React from "react";
import TitleElement from "./TitleElement";
import box from "../icons/Vector.png";
import { useNavigate } from "react-router-dom";

interface Assignment {
  assignment_id: string;
  assignment_name: string;
  due_date: string;
  //complete: boolean;
}

interface Course {
  course_id: string;
  course_name: string;
  course_code: string;
  course_color: string;
  course_image: string;
  Assignment: Assignment[];
}

interface CourseListProps {
  courses: Course[];
  user_group_name: string;
}

const CourseList: React.FC<CourseListProps> = ({ courses, user_group_name }) => {
  const navigate = useNavigate();

  const handleClick = (course: Course) => {
    if (user_group_name == "student") {
      navigate(`/stdcourse/${course.course_id}`, { state: { course } });
    } else if (user_group_name == "instructor") {
      navigate(`/course/${course.course_id}`, { state: { course } });
    } else {
      console.log("Error role");
    }
  };

  const getColorClass = (color: string) => {
    switch (color.toLowerCase()) {
      case "purple":
        return "bg-purple-200 text-purple-800";
      case "yellow":
        return "bg-yellow-200 text-yellow-800";
      case "green":
        return "bg-green-200 text-green-800";
      case "red":
        return "bg-red-200 text-red-800";
      case "pink":
        return "bg-pink-200 text-pink-800";
      case "blue":
        return "bg-blue-200 text-blue-800";
      case "orange":
        return "bg-orange-200 text-orange-800";
      case "brown":
        return "bg-brown-200 text-brown-800";
      default:
        return "bg-gray-200 text-gray-800";
    }
  };

  const getColorBorderClass = (color: string) => {
    switch (color.toLowerCase()) {
      case "purple":
        return "border-purple-200 ";
      case "yellow":
        return "border-yellow-200 ";
      case "green":
        return "border-green-200 ";
      case "red":
        return "border-red-200 ";
      case "pink":
        return "border-pink-200 ";
      case "blue":
        return "border-blue-200 ";
      case "orange":
        return "border-orange-200 ";
      case "brown":
        return "border-brown-200 ";
      default:
        return "border-gray-200 ";
    }
  };

  const previewbox = (
    <div className="border border-gray-200 p-5 rounded-lg shadow-md flex flex-col items-center justify-center bg-white w-72 h-[210px]">
      Please add course first
    </div>
  );

  return (
    <div className="p-4 max-w-full">
      <TitleElement name="Course" icon={box} />
      <div className="overflow-x-auto scrollbar-hide mt-4">
        <div className="flex space-x-4 flex-nowrap whitespace-nowrap w-32">
          {courses.length === 0
            ? previewbox
            : courses.map((course) => (
                <div
                  key={course.course_id}
                  className={`inline-block border-4 ${getColorBorderClass(
                    course.course_color
                  )} p-5 rounded-lg shadow-md flex flex-col items-center justify-center bg-white w-72 h-3/5 cursor-pointer`}
                  onClick={() => handleClick(course)}
                >
                  <img
                    src={course.course_image}
                    alt="Course"
                    className="w-48 h-24 object-cover mb-2 rounded-md"
                  />
                  <p className="text-left text-sm font-medium w-full">
                    {course.course_name}
                  </p>
                  <p className="text-left text-sm font-medium w-full">
                    ({course.course_code})
                  </p>
                  <span
                    className={`block text-xs text-center mt-2 px-2 py-1 rounded-full ${getColorClass(
                      course.course_color
                    )}`}
                  >
                    Assignments count: {course.Assignment ? course.Assignment.length : 0}
                  </span>
                </div>
              ))}
        </div>
      </div>
    </div>
  );
};

export default CourseList;