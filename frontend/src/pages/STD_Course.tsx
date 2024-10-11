import React, { useEffect, useState } from "react";
import LeftMain from "../components/LeftMain";
import noticon from "../icons/bxs_bell.png";
import joinicon from "../icons/material-symbols_join.png";
import dashicon from "../icons/mdi_human-welcome.png";
import exiticon from "../icons/vaadin_exit-o.png";
import RightMain from "../components/RightMain";
import Assicon from "../icons/ion_list.png";
import icon from "../icons/mdi_cog-box.png";
import { FaBars } from "react-icons/fa";
import axios from "axios";
import { useLocation, useNavigate, useParams } from "react-router-dom";
import ProgressBarCourse from "../components/ProgressBarCourse";
import UpcomingAssignment from "../components/UpcomingAssignment";
import AssignmentList from "../components/AssignmentList";

export default function STD_Course() {
  const location = useLocation();
  const navigate = useNavigate();
  const { course_id } = useParams();
  const course = location.state?.course;

  const icons = [dashicon, noticon, joinicon, exiticon];
  const links = ["/stddash", "/notifications", "/std_join_course"];
  const [isOpen, setIsOpen] = useState(false);
  const [profile_image, setProfileImage] = useState("");
  const [user_group_name, setUserGroup] = useState("");
  const [assignment_name, setAssignmentName] = useState("");  
  const [assignment_description, setAssignmentDescription] = useState("");
  const [upcomingAssignments, setUpcomingAssignments] = useState<Assignment[]>([]);
  const [due_date, setDueDate] = useState("");
  const [assignments, setAssignments] = useState<Assignment[]>([]);
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

    const fetchUpComingAssignments = async () => {
      try {
        const res = await axios.get("/api/api/QueryAssignmentsByUserIDSorted");
        if (res.data) {
          const { assignments } = res.data;
          console.log("upcoming assignments", assignments);
          if ( assignments ) setUpcomingAssignments(assignments);
        }else {
          console.warn("No data found in response");
        }
      } catch (error) {
        console.log("Error loading assignments:", error);
      }
    }

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
    fetchUpComingAssignments();
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

  const upcomingAssignmentsWithColor = upcomingAssignments.map((assignment) => {
    const course = courses.find((course) => course.course_id === assignment.course_id);
    const color = course ? course.course_color : "gray";
    return { 
      assignment_id: assignment.assignment_id,
      assignment_name: assignment.assignment_name,
      assignment_due_date: assignment.due_date,
      color 
    };
  });

  if (!course) {
    return <div>Loading...</div>;
  }

  return (
    <div className="bg-B1 flex items-center min-h-screen w-full font-poppins">
      <div className="container mx-auto flex flex-col lg:flex-row gap-5 p-5">
        {/* Left */}
        <div className="bg-white rounded-2xl flex flex-1 relative w-full lg:w-1/2 min-h-[900px] justify-between flex-col">
          <div>
            <LeftMain title={course.course_name} icon={icon} />
            <button
              className="absolute right-10 top-10 block xl:hidden"
              onClick={toggleMenu}
            >
              <FaBars size={40} color="#344B59" />
            </button>
            <div className="flex flex-col md:flex-row">
              <div className="basis-full md:basis-1/2 px-10 p-4">
                <div className="flex gap-5 items-center mb-4">
                  <div className="flex items-center text-E1 gap-3">
                    <img src={Assicon} alt="Assicon" />
                    <h2 className="text-xl">Assignment</h2>
                  </div>
                </div>
                {/* แสดงผล assignments */}
                <AssignmentList Assignment={assignmentsWithColor} showCourseName={false}/> 
              </div>
              <div className="basis-full md:basis-1/2 mt-3 md:mt-0 px-5">
              <UpcomingAssignment UpcomingAssignment={upcomingAssignmentsWithColor} /> 
              </div>
            </div>
          </div>
          <div className="mx-10 mb-4">
            <ProgressBarCourse course={course} />
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
    </div>
  );
}
