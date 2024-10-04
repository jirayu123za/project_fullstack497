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
import { MdOutlineAttachFile } from "react-icons/md";
import { useParams } from "react-router-dom";
import axios from "axios";

export default function InstructorDashboard() {
  const icons = [dashicon, noticon, joinicon, exiticon];
  const links = ["/insdash", "/notifications", "/join", "/exit"];
  
  const [isOpen, setIsOpen] = useState(false);
  const [profileimage, setProfileimage] = useState("");
  const [students, setStudents] = useState<{ StdCode: string; Status: string }[]>([]);
  const [dueDate, setDueDate] = useState("");
  const [uploadedFiles, setUploadedFiles] = useState<string[]>([]);
  const [currentStudentId, setCurrentStudentId] = useState<string | null>(null); // เพิ่ม state เพื่อเก็บรหัสนักศึกษา
  const { assignmentId } = useParams(); // ดึง assignmentId จาก URL

  const toggleMenu = () => {
    setIsOpen(!isOpen);
  };

  // ฟังก์ชันเพื่อดึงรายละเอียด assignment และรหัสนักศึกษา
  const fetchAssignmentDetails = async () => {
    try {
      const response = await axios.get(`/api/assignments/${assignmentId}`); // ดึงข้อมูล assignment จาก assignmentId
      const assignmentData = response.data;
      setDueDate(assignmentData.due_date);
    } catch (error) {
      console.error("เกิดข้อผิดพลาดในการดึงข้อมูล assignment:", error);
    }
  };

  const fetchStudentId = async () => {
    try {
      const response = await axios.get("/api/student-id");
      setCurrentStudentId(response.data.studentId);
    } catch (error) {
      console.error("เกิดข้อผิดพลาดในการดึงรหัสนักศึกษา:", error);
    }
  };

  useEffect(() => {
    fetchAssignmentDetails();
    fetchStudentId();
  }, [assignmentId]);

  const handleFileUpload = (file: File) => {
    if (!currentStudentId) return; 

    setUploadedFiles((prevFiles) => [...prevFiles, file.name]);

    // อัปเดตสถานะของนักศึกษาเมื่ออัปโหลดไฟล์
    setStudents((prevStudents) => {
      const updatedStudents = prevStudents.map((student) =>
        student.StdCode === currentStudentId
          ? { ...student, Status: "not_submitted" }
          : student
      );
      return [
        ...updatedStudents,
        { StdCode: currentStudentId, Status: "submitted" },
      ];
    });

    console.log(`อัปโหลด: ${file.name}`);
  };

  const handleSubmit = async () => {
    try {
      const response = await axios.post("/api/submit-assignment", {
        dueDate,
        uploadedFiles,
        students,
      });
      console.log("ส่งงานสำเร็จ:", response.data);
    } catch (error) {
      console.error("เกิดข้อผิดพลาดในการส่งข้อมูล:", error);
    }
  };

  return (
    <div className="bg-B1 flex items-center min-h-dvh min-w-full font-poppins">
      <div className="container mx-auto flex lg:flex-row gap-5 p-5">
        {/* ซ้าย */}
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

          {/* กำหนด Due Date และปุ่มส่ง */}
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
                      value={dueDate}
                      onChange={(e) => setDueDate(e.target.value)}
                    />
                  </label>
                </div>
              </div>
              <div onClick={handleSubmit}>
                <AssignmentButton text={"Submit"} color={"green"} />
              </div>
            </div>

            <div className="mt-5 flex gap-3 ">
              <div className="basis-5/6 h-full">
                <AssignmentDetail role="Student" />
                <div className="mt-2">
                  <DropBox onFileUpload={handleFileUpload} />
                </div>
              </div>

              {/* สถานะของนักศึกษา */}
              <div className="basis-1/6">
                <AssignmentSubmitted students={students} />
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

        {/* ขวา */}
        <div
          className={`xl:block fixed inset-y-0 right-0 bg-white z-40 transition-transform duration-300 ${isOpen ? "translate-x-0" : "translate-x-full"}`}
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
