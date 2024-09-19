import React, { useState } from "react";
import RightMain from "../components/RightMain";
import LeftMain from "../components/LeftMain";
import noticon from "../icons/bxs_bell.png";
import createicon from "../icons/ion_create.png";
import dashicon from "../icons/mdi_human-welcome.png";
import exiticon from "../icons/vaadin_exit-o.png";
import dashboardicon from "../icons/E1_human-welcome.png";
import CourseList from "../components/CourseList";
import UpcomingElement from "../components/UpcomingElement";
import AssignmentList from "../components/AssignmentList";
import { FaBars } from "react-icons/fa";

export default function InstructorDashboard() {
  const icons = [dashicon, noticon, createicon, exiticon];
  const links = ["/dashboard", "/notifications", "/create"];
  const [isOpen, setIsOpen] = useState(false);

  const toggleMenu = () => {
    setIsOpen(!isOpen);
  };

  return (
    <div className="bg-B1 flex items-center min-h-dvh min-w-full">
      <div className="container mx-auto flex lg:flex-row gap-5 p-5">
        {/* Left */}
        <div className="bg-white rounded-2xl flex-1 relative ">
          <div>
            <LeftMain title="" icon={dashboardicon} />
            <button
              className="absolute right-10 top-10 block lg:hidden"
              onClick={toggleMenu}
            >
              <FaBars size={40} color="#344B59" />
            </button>
          </div>
          <div className="px-4 md:px-6 lg:px-10">
            <div className="mb-4">
              <CourseList />
              </div>
            <div className="flex flex-col lg:flex-row gap-4">
              <div className="lg:flex-1">
                <AssignmentList />
              </div>
              <div className="lg:flex-1">
                <UpcomingElement />
              </div>
            </div>
          </div>
        </div>
        {/* Right */}
        <div
          className={`lg:block fixed inset-y-0 right-0 bg-white z-40 transition-transform duration-300 ${
            isOpen ? "translate-x-0" : "translate-x-full"
          }`}
          onMouseLeave={() => setIsOpen(false)}
        >
          <RightMain icons={icons} links={links} />
        </div>
        <div className="hidden lg:block">
          <RightMain icons={icons} links={links} />
        </div>
      </div>
    </div>
  );
}