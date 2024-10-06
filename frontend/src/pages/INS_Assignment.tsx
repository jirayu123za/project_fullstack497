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
import axios from "axios";
import { useNavigate, useParams } from 'react-router-dom';

export default function InstructorDashboard() {
  const icons = [dashicon, noticon, joinicon, exiticon];
  const links = ["/insdash", "/notifications", "/join", "/exit"];
  const navigate = useNavigate();
  const { course_id, assignment_id } = useParams();
  const [isOpen, setIsOpen] = useState(false);
  const [user_group_name, setUserGroup] = useState("");
  const [profile_image, setProfileImage] = useState("");
  const [students, setStudents] = useState([]);
  const [description, setDescription] = useState("");
  const [due_date, setDueDate] = useState("");
  const [uploadedFiles, setUploadedFiles] = useState<string[]>([]);

  const toggleMenu = () => {
    setIsOpen(!isOpen);
  };

  const handleFileUpload = (file: File) => {
    setUploadedFiles((prevFiles) => [...prevFiles, file.name]);
    console.log(`Uploaded: ${file.name}`);
  };

  const handleUpdate = async () => {
    try {
      const response = await axios.put(`/api/api/UpdateAssignmentByCourseIDAndAssignmentID?course_id=${course_id}&assignment_id=${assignment_id}`, {
        due_date,
        assignment_description: description,
        //uploadedFiles,
        //students,
      });
      console.log("Updated successfully:", response.data);
    } catch (error) {
      console.error("Error updating assignment:", error);
    }
  };

  const handleDelete = async () => {
    console.log("call delete assignment");
    
    if (window.confirm("Are you sure you want to delete this assignment?")) {
      try {
        const response = await axios.delete(`/api/api/DeleteAssignmentByCourseIDAndAssignmentID?course_id=${course_id}&assignment_id=${assignment_id}`);
        
        if (response.status === 200) {
          navigate("/insdash");
        }
        console.log("Deleted successfully:", response.data);
      } catch (error) {
        console.error("Error deleting assignment:", error);
      }
    }
  };

  useEffect(() => {
    console.log("log API assignment due_date:", due_date);
    console.log("log API assignment description:", description);
  }, [due_date, description]);

  useEffect(() => {
    const fetchPersonalData = async () => {      
      try {
        const res = await axios.get("/api/QueryPersonDataByUserID");
    
        if (res.data && res.data.user && res.data.user.length > 0) {
          const { user_image_url, /*user_first_name*/ } = res.data.user[0];
    
          if (user_image_url) setProfileImage(user_image_url);
          //if (user_first_name) setFirstName(user_first_name);
          console.log(res.data);
          
        } else {
          console.warn("No data found in response");
        }
      } catch (error) {
        console.error("Error loading personal data:", error);
      }
    };

    const fetchPersonalUserGroup = async () => {
      try {
        const res = await axios.get("/api/api/QueryUserGroupByUserID");
        if (res.data) {
          const { user_group_name } = res.data;
          if (user_group_name) setUserGroup(user_group_name);
          console.log(res.data);
          
        } else {
          console.warn("No data found in response");
        }
      } catch (error) {
        console.error("Error loading user_group_name:", error);
      }
    };

    const fetchAssignmentDetails = async () => {
      try {
        const res = await axios.get(`/api/api/QueryAssignmentsByCourseIDAndAssignmentID?course_id=${course_id}&assignment_id=${assignment_id}`);
        if (res.data) {
          const { due_date, assignment_description } = res.data.assignment;

          if (due_date) {
            const dateObj = new Date(due_date);
            const formattedDate = dateObj.toISOString().split('T')[0];
            setDueDate(formattedDate);
          }
          if (assignment_description) setDescription(assignment_description);
    
          console.log("fetchAssignmentDetails", res.data);
        } else {
          console.warn("No data found in response");
        }
      } catch (error) {
        console.error("Error loading assignment details:", error);
      }
    };

    fetchPersonalData();
    fetchPersonalUserGroup();
    fetchAssignmentDetails();
  }, []);

  // Fetch students' assignment status from backend
  // useEffect(() => {
  //   const fetchStudents = async () => {
  //     try {
  //       const response = await axios.get("/api/course/students"); // แก้ URL ให้ตรงกับ backend ของคุณ
  //       setStudents(response.data);
  //     } catch (error) {
  //       console.error("Error fetching student data:", error);
  //     }
  //   };

  //   fetchStudents();
  // }, []);

  useEffect(() => {
    const fetchStudents = async () => {
      try {
        const response = await fetch("/students.json");
        const data = await response.json();
        const submittedStudents = data.students.filter((student: { Status: string; }) => student.Status === "submitted");
        setStudents(submittedStudents);
      } catch (error) {
        console.error("Error fetching student data:", error);
      }
    };

    fetchStudents();
  }, []);

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

          {/* Due Date and Buttons */}
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
                      value={due_date}
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

            {/* Assignment Detail, DropBox, and Student List */}
            <div className="mt-5 flex gap-3 ">
              <div className="basis-5/6 h-full">
                <div className="flex items-start gap-5 ">
                  <AssignmentDetail user_group_name={user_group_name} assignment_description={description} onChange={(e) => setDescription(e.target.value)}/>
                </div>
                <div className="mt-2">
                  <DropBox onFileUpload={handleFileUpload} />
                </div>
              </div>

              {/* Student Status */}
              <div className="basis-1/6">
                <div className="้">
                  <AssignmentSubmitted students={students} />
                </div>
                <div className="flex flex-col space-y-2 ">
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
          <RightMain icons={icons} links={links} profile_image={profile_image} user_group_name={user_group_name}/>
        </div>
        <div className="hidden xl:block">
          <RightMain icons={icons} links={links} profile_image={profile_image} user_group_name={user_group_name}/>
        </div>
      </div>
    </div>
  );
}
