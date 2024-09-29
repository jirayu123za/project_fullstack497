import React, { useEffect, useState } from "react";
import LeftMain from "../components/LeftMain";
import noticon from "../icons/bxs_bell.png";
import joinicon from "../icons/material-symbols_join.png";
import dashicon from "../icons/mdi_human-welcome.png";
import exiticon from "../icons/vaadin_exit-o.png";
import RightMain from "../components/RightMain";
import Assicon from "../icons/ion_list.png";
import icon from "../icons/mdi_cog-box.png";
import { FaBars } from "react-icons/fa";
import axios from "axios";
import { useLocation, useNavigate, useParams } from "react-router-dom";
import UpcomingElement from "../components/UpcomingElement";
import ProgressBarCourse from "../components/ProgressBarCourse";

export default function STD_Course() {
  const icons = [dashicon, noticon, joinicon, exiticon];
  const links = ["/stddash", "/notifications", "/stdcreate", "/exit"];
  const [isOpen, setIsOpen] = useState(false);
  const [profileimage, setProfileimage] = useState("");
  const { course_id } = useParams();  // ใช้ course_id จาก URL
  const location = useLocation();
  const course = location.state?.course;

  const navigate = useNavigate();  // ใช้ useNavigate สำหรับเปลี่ยนหน้า

  const toggleMenu = () => {
    setIsOpen(!isOpen);
  };

  // ฟังก์ชันสำหรับการเปลี่ยนไปยังหน้า STD_Assignment
  const handleNavigateToAssignment = (assignment_id: string) => {
    navigate(`/stdassignment/${assignment_id}`, {
      state: { assignment_id, course_id } // ส่ง assignment_id และ course_id ไปพร้อมกัน
    });
  };

  useEffect(() => {
    const fetchPersonaldata = async () => {
      try {
        const response = await fetch("/data.json");
        const data = await response.json();
        setProfileimage(data.profileimage);
      } catch (error) {
        console.error("Error loading profile image:", error);
      }
    };

    fetchPersonaldata();
  }, [course_id]);

  const getColorClass = (color: string) => {
    switch (color.toLowerCase()) {
      case "purple":
        return "border-purple-300 text-purple-400";
      case "yellow":
        return "border-yellow-300 text-yellow-400";
      case "green":
        return "border-green-300 text-green-400";
      case "red":
        return "border-red-300 text-red-400";
      case "pink":
        return "border-pink-300 text-pink-400";
      case "blue":
        return "border-blue-300 text-blue-400";
      case "orange":
        return "border-orange-300 text-orange-400";
      case "brown":
        return "border-brown-300 text-brown-400";
      default:
        return "border-gray-300 text-gray-400";
    }
  };

  return (
    <div className="bg-B1 flex items-center min-h-screen w-full font-poppins">
      <div className="container mx-auto flex flex-col lg:flex-row gap-5 p-5">
        {/* Left */}
        <div className="bg-white rounded-2xl flex flex-1 relative w-full lg:w-1/2 min-h-[900px] justify-between flex-col">
          <div>
            <LeftMain title={course.course_name} icon={icon} />
            <button
              className="absolute right-10 top-10 block xl:hidden"
              onClick={toggleMenu}
            >
              <FaBars size={40} color="#344B59" />
            </button>
            <div className="flex flex-col md:flex-row">
              <div className="basis-full md:basis-1/2 px-10 p-4">
                <div className="flex gap-5 items-center mb-4">
                  <div className="flex items-center text-E1 gap-3">
                    <img src={Assicon} alt="Assicon" />
                    <h2 className="text-xl">Assignment</h2>
                  </div>
                </div>
                {/* แสดงผล assignments */}
                {Array.isArray(course?.assignments) &&
                  course.assignments.map((assignment) => {
                    return (
                      <div
                        key={assignment.assignment_id}
                        className={`flex items-center justify-between border-4 p-2.5 mb-2 rounded-lg shadow-sm h-[87px] cursor-pointer ${getColorClass(
                          course.color // ใช้ color จาก course
                        )}`}
                        onClick={() => handleNavigateToAssignment(assignment.assignment_id)}  // เมื่อกด จะเรียกฟังก์ชันนำทาง
                      >
                        <div className="flex flex-col">
                          <p className="text-xl">{assignment.title}</p>
                          <p className="text-sm text-gray-600">
                            Due Date: {assignment.due_date}
                          </p>
                        </div>
                      </div>
                    );
                  })}
              </div>
              <div className="basis-full md:basis-1/2 mt-3 md:mt-0 px-5">
                <UpcomingElement courses={course} />
              </div>
            </div>
          </div>
          <div className="mx-10 mb-4">
            <ProgressBarCourse course={course} />
          </div>
        </div>
        {/* Right */}
        <div
          className={`xl:block fixed inset-y-0 right-0 bg-white z-40 transition-transform duration-300 ${
            isOpen ? "translate-x-0" : "translate-x-full"
          }`}
          onMouseLeave={() => setIsOpen(false)}
        >
          <RightMain icons={icons} links={links} profileimage={profileimage} />
        </div>
        <div className="hidden xl:block">
          <RightMain icons={icons} links={links} profileimage={profileimage} />
        </div>
      </div>
    </div>
  );
}
