import { useEffect, useState } from "react";
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
import StudentList from "../components/StudentList";
import axios from "axios";
import { useLocation, useNavigate, useParams } from "react-router-dom";
import AssignmentList from "../components/AssignmentList";

export default function INS_Course() {
  const navigate = useNavigate();
  const location = useLocation();
  const course = location.state?.course;
  const icons = [dashicon, noticon, createicon, exiticon];
  const links = ["/stddash", "/notifications", "/create"];
  const { course_id } = useParams();
  const [assignments, setAssignments] = useState<Assignment[]>([]);
  const [profile_image, setProfileImage] = useState("");
  const [user_group_name, setUserGroup] = useState("");
  const [isPopupOpen, setIsPopupOpen] = useState(false);
  const [isOpen, setIsOpen] = useState(false);
  const [assignment_name, setAssignmentName] = useState("");  
  const [assignment_description, setAssignmentDescription] = useState("");
  const [due_date, setDueDate] = useState("");
  const [courses, setCourses] = useState<Course[]>([]);

  interface Course {
    course_id: string;
    course_color: string;
    course_name: string;
  }

  interface Assignment {
    assignment_id: string;
    course_id: string;
    assignment_name: string;
    due_date: string;
  }

  useEffect(() => {
    console.log("course_id from URL params:", course_id);
  }, [course_id]);

  const toggleMenu = () => {
    setIsOpen(!isOpen);
  };

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


  const handleDelete = async () => {
    console.log("Call handleDelete course"); 

    if (window.confirm("Are you sure you want to delete this course?")) {
      try {
        const resp = await axios.delete(`/api/api/DeleteCourse?course_id=${course_id}`);

        if (resp.status === 200) {
          console.log(resp.data);
          navigate("/insdash");
        }
      } catch (error) {
        console.log("Error deleting course:", error);
      }
    }
  };

  const handleAddAssignment = async () => {
    try {
      const response = await axios.post(`/api/api/CreateAssignment?course_id=${course_id}`, {
        assignment_name: assignment_name,
        assignment_description: assignment_description,
        due_date: due_date,
      });

      console.log("Assignment added:", response.data);
      
      setIsPopupOpen(false);
      setAssignmentName("");
      setAssignmentDescription("");
      setDueDate("");
    } catch (error) {
      console.error("Error adding assignment:", error);
    }
  };

  useEffect(() => {
    const fetchPersonalData = async () => {      
      try {
        const res = await axios.get("/api/QueryPersonDataByUserID");
    
        if (res.data && res.data.user && res.data.user.length > 0) {
          const { user_image_url } = res.data.user[0];
    
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

    // Assignments by course_id
    const fetchAssignments = async () => {
      try {
        const res = await axios.get(`/api/api/QueryAssignmentsByCourseID?course_id=${course_id}`);
        if (res.data) {          
          const { assignments } = res.data;
          console.log("assignments", assignments);
          if ( assignments ) setAssignments(assignments);
        }else {
          console.warn("No data found in response");
        }
      } catch (error) {
        console.log("Error loading assignments:", error);
      }
    };

    const fetchCourses = async () => {
      try {
        const res = await axios.get("/api/QueryCourseByUserID");
        if (res.data) {
          const { courses } = res.data;
          console.log("courses", courses);
          if (courses) setCourses(courses);
        } else {
          console.warn("No data found in response");
        }
      } catch (error) {
        console.error("Error loading courses:", error);
      }
    };

    fetchCourses() ;
    fetchPersonalUserGroup();
    fetchPersonalData();
    fetchAssignments();
  }, [course_id]);

  const assignmentsWithColor = assignments.map((assignment) => {
    const course = courses.find((course) => course.course_id === assignment.course_id);
    const color = course ? course.course_color : "gray";
    const course_name = course ? course.course_name : "Unknown Course";
    return {
      assignment_id: assignment.assignment_id,
      assignment_name: assignment.assignment_name,
      assignment_due_date: assignment.due_date,
      color,
      course_name: course_name,
      course_id: assignment.course_id
    };
  });

  return (
    <div className="bg-B1 flex items-center min-h-screen w-full font-poppins">
      <div className="container mx-auto flex flex-col lg:flex-row gap-5 p-5">
        {/* Left */}
        <div className="bg-white rounded-2xl flex-1 relative w-full lg:w-1/2 min-h-[900px]">
          <div>
            <LeftMain title={course.course_name} icon={icon} />
            <button
              className="absolute right-10 top-10 block xl:hidden"
              onClick={toggleMenu}
            >
              {""}
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

                <AssignmentList Assignment={assignmentsWithColor} showCourseName={false} user_group_name={user_group_name}/> 

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
                    </div>{" "}
                    <input
                      className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                      id="code"
                      type="text"
                      defaultValue={course.course_id}
                      readOnly
                    />
                  </label>
                </div>
                <div className="mt-5">
                  <StudentList user_group_name={user_group_name}/>
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
          <RightMain icons={icons} links={links} profile_image={profile_image} user_group_name={user_group_name} handleLogout={handleLogout}/>
          </div>
        <div className="hidden xl:block">
          <RightMain icons={icons} links={links} profile_image={profile_image} user_group_name={user_group_name} handleLogout={handleLogout}/>
        </div>
      </div>

      {isPopupOpen && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
          <div className="bg-white p-8 rounded-lg shadow-lg">
            <h2 className="text-2xl mb-4">Add New Assignment</h2>
            <label className="block mb-2">
              Assignment Name:
              <input
                type="text"
                value={assignment_name}
                onChange={(e) => setAssignmentName(e.target.value)}
                className="border p-2 w-full"
              />
            </label>
            <label className="block mb-2">
              Assignment Description:
              <input
                type="text"
                value={assignment_description}
                onChange={(e) => setAssignmentDescription(e.target.value)}
                className="border p-2 w-full"
              />
            </label>
            <label className="block mb-2">
              Due Date:
              <input
                type="date"
                value={due_date}
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
