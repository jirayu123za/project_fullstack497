import LeftMain from "../components/LeftMain";
import RightMain from "../components/RightMain";
import AssignmentDetail from "../components/AssignmentDetail";
import DropBox from "../components/DropBox";
import AssignmentSubmitted from "../components/AssignmentSubmitted";
import noticon from "../icons/bxs_bell.png";
import joinicon from "../icons/material-symbols_join.png";
import dashicon from "../icons/mdi_human-welcome.png";
import exiticon from "../icons/vaadin_exit-o.png";
import Assign from "../icons/ion_list.png";
import { useEffect, useState } from "react";
import { FaBars } from "react-icons/fa";
import AssignmentButton from "../components/AssignmentButton";
import TitleElement from "../components/TitleElement";
import dateicon from "../icons/material-symbols-light_update.png";
import { MdOutlineAttachFile } from "react-icons/md";

export default function STD_Assignment() {
  const icons = [dashicon, noticon, joinicon, exiticon];
  const links = ["/insdash", "/notifications", "/join", "/exit"];

  const [isOpen, setIsOpen] = useState(false);
  const [profileimage, setProfileimage] = useState("");
  // const [courses, setCourses] = useState([]);
  const [dueDate, setDueDate] = useState("");

  const toggleMenu = () => {
    setIsOpen(!isOpen);
  };

  const handleUpdate = () => {
    console.log("Updating course with due date:", dueDate);
  };

  const [uploadedFiles, setUploadedFiles] = useState<string[]>([]);

  //   const handleUpdate = () => {
  //     console.log("Update Assignment Details");
  //   };

  const handleDelete = () => {
    console.log("Delete Assignment");
  };

  const handleFileUpload = (file: File) => {
    setUploadedFiles((prevFiles) => [...prevFiles, file.name]);
    console.log(`Uploaded: ${file.name}`);
  };

  useEffect(() => {
    const fetchPersonaldata = async () => {
      try {
        const response = await fetch("/data.json");
        const data = await response.json();
        setProfileimage(data.profileimage);
      } catch (error) {
        console.error("Error loading email:", error);
      }
    };

    fetchPersonaldata();
  }, []);

  const ConfigAssignment = [
    { StdCode: "640610629", Status: "#E61616" },
    { StdCode: "640629042", Status: "#E61616" },
    { StdCode: "633934788", Status: "#E61616" },
    { StdCode: "634124894", Status: "#E61616" },
    { StdCode: "123408895", Status: "#E61616" },
    { StdCode: "640610629", Status: "#E61616" },
    { StdCode: "640629042", Status: "#E61616" },
    { StdCode: "633934788", Status: "#E61616" },
    { StdCode: "634124894", Status: "#E61616" },
    { StdCode: "123408895", Status: "#E61616" },
    { StdCode: "640610629", Status: "#E61616" },
    { StdCode: "640629042", Status: "#E61616" },
    { StdCode: "633934788", Status: "#E61616" },
    { StdCode: "634124894", Status: "#E61616" },
    { StdCode: "123408895", Status: "#E61616" },
  ];

  return (
    <div className="bg-B1 flex items-center min-h-dvh min-w-full font-poppins">
      <div className="container mx-auto flex lg:flex-row gap-5 p-5">
        {/* Left */}
        <div className="bg-white rounded-2xl flex-1 relative min-h-[900px]">
          <div>
            <LeftMain title="Assignment" icon={Assign} />
            <button
              className="absolute right-10 top-10 block xl:hidden"
              onClick={toggleMenu}
            >
              <FaBars size={40} color="#344B59" />
            </button>
          </div>
          {/* code here */}
          <div className="px-4 md:px-6 lg:px-10">
            <div className="flex justify-between items-center">
              <div>
                <div className="flex mt-3">
                  <label
                    htmlFor="date"
                    className="flex items-center w-full gap-3"
                  >
                    <div className="flex-shrink-0">Due Date</div>
                    <input
                      className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                      id="date"
                      type="date"
                      onChange={(e) => setDueDate(e.target.value)}
                    />
                  </label>
                </div>
              </div>
              <div className="flex gap-3">
                <div onClick={handleUpdate}>
                  <AssignmentButton text={"Update"} color={"yellow"} />
                </div>
                <div onClick={handleDelete}>
                  <AssignmentButton text={"Delete"} color={"red"} />
                </div>
              </div>
            </div>
            <div className="mt-5 flex gap-3 ">
              <div className="basis-5/6 h-full">
                <div className="flex items-start gap-5 ">
                  <AssignmentDetail role="Instructor" />
                </div>
                <div className="mt-2">
                  <DropBox onFileUpload={handleFileUpload} />
                </div>
              </div>
              <div className="basis-1/6">
                <div className="">
                  <AssignmentSubmitted ConfigAssignment={ConfigAssignment} />
                </div>
                <div className="flex flex-col space-y-2">
                  {uploadedFiles.map((fileName, index) => (
                    <div
                      key={index}
                      className="flex items-center border-2 border-M1 rounded-lg p-2 w-[200px] h-[50px] mt-2 overflow-hidden gap-3"
                    >
                      <MdOutlineAttachFile size={25} color="#344B59" />
                      <p className="text-[#5A8FAA]">{fileName}</p>
                    </div>
                  ))}
                </div>
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
