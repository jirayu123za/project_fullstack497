import React, { useEffect, useState } from "react";
import RightMain from "../components/RightMain";
import LeftMain from "../components/LeftMain";
import noticon from "../icons/bxs_bell.png";
import joinicon from "../icons/material-symbols_join.png";
import dashicon from "../icons/mdi_human-welcome.png";
import exiticon from "../icons/vaadin_exit-o.png";
import dashboardicon from "../icons/E1_human-welcome.png";
import CourseList from "../components/CourseList";
import UpcomingElement from "../components/UpcomingElement";
import AssignmentList from "../components/AssignmentList";
import { FaBars } from "react-icons/fa";

export default function STD_Dashboard() {
  const icons = [dashicon, noticon, joinicon, exiticon];
  const links = ["/stddash", "/notifications", "/stdcreate", "/exit"];

  const [isOpen, setIsOpen] = useState(false);
  const [profileimage, setProfileimage] = useState("");
  const [courses, setCourses] = useState([]);
  const [firstname, setFirstname] = useState("");

  const toggleMenu = () => {
    setIsOpen(!isOpen);
  };

  useEffect(() => {
    const fetchPersonaldata = async () => {
      try {
        const response = await fetch("/data.json");
        const data = await response.json();
        setProfileimage(data.profileimage);
        setFirstname(data.firstName);
      } catch (error) {
        console.error("Error loading email:", error);
      }
    };

    const fetchCourses = async () => {
      try {
        const response = await fetch("/coursedetail.json");
        const data = await response.json();
        setCourses(data); // เก็บข้อมูล courses
      } catch (error) {
        console.error("Error loading courses:", error);
      }
    };

    fetchPersonaldata();
    fetchCourses();
  }, []);

  return (
    <div className="bg-B1 flex items-center min-h-dvh min-w-full">
      <div className="container mx-auto flex lg:flex-row gap-5 p-5">
        {/* Left */}
        <div className="bg-white rounded-2xl flex-1 relative min-h-[900px]">
          <div>
            <LeftMain title={firstname} icon={dashboardicon} />
            <button
              className="absolute right-10 top-10 block xl:hidden"
              onClick={toggleMenu}
            >
              <FaBars size={40} color="#344B59" />
            </button>
          </div>
          <div className="px-4 md:px-6 lg:px-10">
            <div className="mb-4">
              <CourseList courses={courses} role="student" />
            </div>
            <div className="flex flex-col lg:flex-row gap-4">
              <div className="lg:flex-1">
                <AssignmentList courses={courses} />
              </div>
              <div className="lg:flex-1">
                <UpcomingElement courses={courses} />
              </div>
            </div>
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
