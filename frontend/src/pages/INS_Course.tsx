import React from "react";
import LeftMain from "../components/LeftMain";
import noticon from "../icons/bxs_bell.png";
import createicon from "../icons/ion_create.png";
import dashicon from "../icons/mdi_human-welcome.png";
import exiticon from "../icons/vaadin_exit-o.png";
import RightMain from "../components/RightMain";
import TitleElement from "../components/TitleElement";
import Assicon from "../icons/ion_list.png";

export default function INS_Course() {
  const icons = [dashicon, noticon, createicon, exiticon];
  const links = ["/dashboard", "/notifications", "/create", "/exit"];
  return (
    <div className="bg-B1 h-screen flex justify-center items-center gap-8">
      <div className="w-[1174px] h-[900px] bg-white rounded-2xl">
        <LeftMain title="Fullstack" icon={dashicon} />
        <div>
          <TitleElement name="Assignment" icon={Assicon} />
        </div>
      </div>
      <RightMain icons={icons} links={links} />
    </div>
  );
}
