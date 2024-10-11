import { useEffect, useState } from "react";
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
import { FaBars } from "react-icons/fa";
import AssignmentButton from "../components/AssignmentButton";
import { MdOutlineAttachFile } from "react-icons/md";
import { useNavigate, useParams } from "react-router-dom";
import axios from "axios";

export default function InstructorDashboard() {
  const navigate = useNavigate();
  const { course_id, assignment_id } = useParams();
  const icons = [dashicon, noticon, joinicon, exiticon];
  const links = ["/stddash", "/notifications", "/std_join_course"];
  
  const [isOpen, setIsOpen] = useState(false);
  const [user_group_name, setUserGroup] = useState("");
  const [profile_image, setProfileImage] = useState("");
  const [user_first_name, setFirstName] = useState("");
  const [user_id, setUserID] = useState("");
  const [students, setStudents] = useState([]);
  const [description, setDescription] = useState("");
  const [due_date, setDueDate] = useState("");
  const [uploadedFiles, setUploadedFiles] = useState<File[]>([]);
  const [isDuePassed, setIsDuePassed] = useState(false); 

  const toggleMenu = () => {
    setIsOpen(!isOpen);
  };

  const fetchPersonalData = async () => {      
    try {
      const res = await axios.get("/api/QueryPersonDataByUserID");
  
      if (res.data && res.data.user && res.data.user.length > 0) {
        const { user_image_url, user_first_name, user_id } = res.data.user[0];
  
        if (user_image_url) setProfileImage(user_image_url);
        if (user_first_name) setFirstName(user_first_name);
        if (user_id) setUserID(user_id);
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


          const now = new Date();
          setIsDuePassed(dateObj < now);
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
  
  const fetchSubmissions = async () => {
    try {
      const res = await axios.get(`/api/api/QuerySubmissionsStatus?course_id=${course_id}&assignment_id=${assignment_id}`);
      if (res.data) {
        const { submissions } = res.data;
        if (submissions) setStudents(submissions);
        console.log("fetchSubmissions", res.data);
      } else {
        console.warn("No data found in response");
      }
    } catch (error) {
      console.error("Error loading submissions:", error);
    }
  };

  useEffect(() => {
    fetchPersonalData();
    fetchPersonalUserGroup();
    fetchAssignmentDetails();
    fetchSubmissions();
  }, []);

  const handleLogout = async () => {
    try {
      const resp = await axios.post("/api/logout");
      if (resp.status === 200){
        navigate("/landing");
      }
    } catch (error) {
      console.error("Error during logout:", error);
    }
  };

  const handleFileUpload = (file: File) => {
    setUploadedFiles((prevFiles) => [...prevFiles, file]);
    console.log(`Uploaded: ${file.name}`);
  };

  const handleSubmit = async () => {
    try {
      const formData = new FormData();  
      uploadedFiles.forEach((file) => {
        formData.append("files", file);
      });
  
      formData.append("user_id", user_id);
      if (assignment_id) {
        formData.append("assignment_id", assignment_id);
      }
      formData.append("user_group_name", user_group_name);
      formData.append("user_name", user_first_name);
  
      const response = await axios.post("/api/api/CreateSubmission", formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      });
  
      console.log("File uploaded successfully:", response.data);
    } catch (error) {
      console.error("Error updating assignment:", error);
    }
  };

  return (
    <div className="bg-B1 flex items-center min-h-dvh min-w-full font-poppins">
      <div className="container mx-auto flex lg:flex-row gap-5 p-5">
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
                      readOnly 
                      onChange={(e) => setDueDate(e.target.value)}
                    />
                  </label>
                </div>
              </div>

              {!isDuePassed && (
                <div onClick={handleSubmit}>
                  <AssignmentButton text={"Submit"} color={"green"} />
                </div>
              )}
            </div>

            <div className="mt-5 flex gap-3 ">
              <div className="basis-5/6 h-full">
                <AssignmentDetail user_group_name={user_group_name} assignment_description={description} onChange={(e) => setDescription(e.target.value)}/>
                  <div className="mt-2">
                    <DropBox onFileUpload={handleFileUpload} isDuePassed={isDuePassed} />
                  </div>
              </div>

              <div className="basis-1/6">
                  <AssignmentSubmitted submissions={students} />
                  <div className="flex flex-col space-y-2">
                  {uploadedFiles.map((fileName, index) => (
                    <div
                      key={index}
                      className="flex items-center border-2 border-M1 rounded-lg p-2 w-[200px] h-[50px] mt-2 overflow-hidden gap-3"
                    >
                      <MdOutlineAttachFile size={25} color="#344B59" />
                      <p className="text-[#5A8FAA] truncate w-full overflow-hidden text-ellipsis whitespace-nowrap">{fileName.name}</p>
                    </div>
                  ))}
                </div>
              </div>
            </div>
          </div>
        </div>

        {/* Right */}
        <div
          className={`xl:block fixed inset-y-0 right-0 bg-white z-40 transition-transform duration-300 ${isOpen ? "translate-x-0" : "translate-x-full"}`}
          onMouseLeave={() => setIsOpen(false)}
        >
          <RightMain icons={icons} links={links} profile_image={profile_image} user_group_name={user_group_name} handleLogout={handleLogout}/>
          </div>
        <div className="hidden xl:block">
          <RightMain icons={icons} links={links} profile_image={profile_image} user_group_name={user_group_name} handleLogout={handleLogout}/>
        </div>
      </div>
    </div>
  );
}
