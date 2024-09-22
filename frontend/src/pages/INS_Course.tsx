import React, { useEffect, useState } from "react";
import LeftMain from "../components/LeftMain";
import noticon from "../icons/bxs_bell.png";
import createicon from "../icons/ion_create.png";
import dashicon from "../icons/mdi_human-welcome.png";
import exiticon from "../icons/vaadin_exit-o.png";
import RightMain from "../components/RightMain";
import TitleElement from "../components/TitleElement";
import Assicon from "../icons/ion_list.png";
import icon from "../icons/mdi_cog-box.png";
import { FaBars, FaPlusSquare } from "react-icons/fa";
import AssignmentButton from "../components/AssignmentButton";
import codeicon from "../icons/material-symbols_code.png";
import dateicon from "../icons/material-symbols-light_update.png";
import FriendList from "../components/FriendList";
import axios from "axios";
import { useNavigate, useParams } from "react-router-dom";

export default function INS_Course() {
  const icons = [dashicon, noticon, createicon, exiticon];
  const links = ["/dashboard", "/notifications", "/create", "/exit"];
  const [isOpen, setIsOpen] = useState(false);
  const [courses, setCourses] = useState([]);
  const [profileimage, setProfileimage] = useState("");
  const { course_id } = useParams();
  const [isPopupOpen, setIsPopupOpen] = useState(false); // ใช้สำหรับแสดง popup
  const [title, setTitle] = useState("");
  const [dueDate, setDueDate] = useState("");

  const navigate = useNavigate();

  const toggleMenu = () => {
    setIsOpen(!isOpen);
  };

  const handleDelete = () => {
    if (window.confirm("Are you sure you want to delete this course?")) {
      console.log("Deleting course");
      navigate("/dashboard");
    }
  };

  // const handleClick = async (assignmentId: string) => {
  //   try {
  //     await axios.post("/api/assignments", { assignment_id: assignmentId });
  //     navigate("/assignment");
  //   } catch (error) {
  //     console.error("Error sending assignment ID:", error);
  //   }
  // };

  // useEffect(() => {
  //   const fetchAssignments = async () => {
  //     try {
  //       // const response = await fetch(`/api/assignments?course_id=${course_id}`);
  //       const response = await fetch("/assignment.json");
  //       const data = await response.json();
  //       // console.log(data.assignments);
  //       // console.log(data);
  //       setCourses(data);
  //     } catch (error) {
  //       console.error("Error fetching assignments:", error);
  //     }
  //   };

  //   if (course_id) {
  //     fetchAssignments();
  //   }
  //   const fetchPersonaldata = async () => {
  //     try {
  //       const response = await fetch("/data.json");
  //       const data = await response.json();
  //       setProfileimage(data.profileimage);
  //     } catch (error) {
  //       console.error("Error loading email:", error);
  //     }
  //   };

  //   fetchPersonaldata();
  //   fetchAssignments();
  // }, [course_id]);

  const handleAddAssignment = async () => {
    try {
      // const response = await axios.post("/api/addAssignment", {
      //   title: title,
      //   due_date: dueDate,
      // });

      // mock การตอบกลับของ axios.post โดยใช้ Promise.resolve()
      const mockResponse = Promise.resolve({
        data: {
          title: title,
          due_date: dueDate,
        },
      });
      const response = await mockResponse;

      console.log("Assignment added:", response.data);
      // ปิด popup และล้างค่าหลังจากบันทึกสำเร็จ
      setIsPopupOpen(false);
      setTitle("");
      setDueDate("");
      // โหลดข้อมูลใหม่หลังจากเพิ่ม assignment สำเร็จ
      fetchAssignments();
    } catch (error) {
      console.error("Error adding assignment:", error);
    }
  };

  const fetchAssignments = async () => {
    try {
      const response = await fetch("/assignment.json");
      const data = await response.json();
      setCourses(data);
    } catch (error) {
      console.error("Error fetching assignments:", error);
    }
  };

  useEffect(() => {
    if (course_id) {
      fetchAssignments();
    }

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
        <div className="bg-white rounded-2xl flex-1 relative w-full lg:w-1/2 min-h-[900px]">
          <div>
            <LeftMain title={courses.course_name} icon={icon} />
            <button
              className="absolute right-10 top-10 block xl:hidden"
              onClick={toggleMenu}
            >
              <FaBars size={40} color="#344B59" />
            </button>
            <div className="flex flex-col md:flex-row">
              <div className="basis-full md:basis-1/2 px-10">
                <div className="flex gap-5 items-center mb-4">
                  <div className="flex items-center text-E1 gap-3">
                    <img src={Assicon} alt="Assicon" />
                    <h2 className="text-xl">Assignment</h2>
                  </div>
                  <div onClick={() => setIsPopupOpen(true)}>
                    <FaPlusSquare size={40} color="#93B955" />
                  </div>
                </div>
                {/* แสดง Assignment ที่ดึงมาจาก API */}
                {Array.isArray(courses.assignments) &&
                  courses.assignments.map((assignment) => {
                    return (
                      <div
                        key={assignment.assignment_id}
                        className={`flex items-center justify-between border-4 p-2.5 mb-2 rounded-lg shadow-sm h-[87px] cursor-pointer ${getColorClass(
                          courses.color
                        )}`}
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
              <div className="basis-full md:basis-1/2 px-10 mt-3 md:mt-0">
                <div className="flex justify-end gap-3">
                  <div onClick={handleDelete}>
                    <AssignmentButton text={"Delete"} color={"red"} />
                  </div>
                </div>
                <div className="flex mt-4">
                  <label
                    htmlFor="code"
                    className="flex items-center w-full gap-3"
                  >
                    <div className="flex-shrink-0">
                      <TitleElement name={"Code"} icon={codeicon} />
                    </div>
                    <input
                      className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                      id="code"
                      type="text"
                      defaultValue={courses.course_id}
                      readOnly
                    />
                  </label>
                </div>
                <div className="mt-5">
                  <FriendList />
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

      {isPopupOpen && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
          <div className="bg-white p-8 rounded-lg shadow-lg">
            <h2 className="text-2xl mb-4">Add New Assignment</h2>
            <label className="block mb-2">
              Title:
              <input
                type="text"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                className="border p-2 w-full"
              />
            </label>
            <label className="block mb-2">
              Due Date:
              <input
                type="date"
                value={dueDate}
                onChange={(e) => setDueDate(e.target.value)}
                className="border p-2 w-full"
              />
            </label>
            <div className="flex justify-end gap-4 mt-5">
              <button
                className="bg-gray-300 p-2 rounded"
                onClick={() => setIsPopupOpen(false)}
              >
                Cancel
              </button>
              <button
                className="bg-R4 text-white p-2 rounded"
                onClick={handleAddAssignment}
              >
                Add Assignment
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
