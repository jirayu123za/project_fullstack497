import React from "react";
import RightMain from "../components/RightMain";
import LeftMain from "../components/LeftMain";
import noticon from "../icons/bxs_bell.png";
import joinicon from "../icons/material-symbols_join.png";
import dashicon from "../icons/mdi_human-welcome.png";
import exiticon from "../icons/vaadin_exit-o.png";
import CourseList from "../components/CourseList";
import UpcomingElement from "../components/UpcomingElement";
import AssignmentList from "../components/AssignmentList";

export default function InstructorDashboard() {
  const icons = [dashicon, noticon, joinicon, exiticon];
  const links = ["/dashboard", "/notifications", "/join", "/exit"];

  return (
    <div className="bg-B1 h-screen flex justify-center items-center gap-8">
      <div className="w-[1174px] h-[900px] bg-white rounded-2xl">
        <LeftMain name="Natacha" />
        <div
          className="mb-16 mx-10
        "
        >
          <CourseList />
        </div>
        <div className="flex justify-between mx-10">
          <AssignmentList />
          <UpcomingElement />
        </div>
      </div>
      <RightMain icons={icons} links={links} />
    </div>
  );
}
