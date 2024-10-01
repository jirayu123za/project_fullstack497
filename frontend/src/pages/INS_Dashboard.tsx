import { useEffect, useState } from "react";
import RightMain from "../components/RightMain";
import LeftMain from "../components/LeftMain";
import noticon from "../icons/bxs_bell.png";
import createicon from "../icons/ion_create.png";
import dashicon from "../icons/mdi_human-welcome.png";
import exiticon from "../icons/vaadin_exit-o.png";
import dashboardicon from "../icons/E1_human-welcome.png";
import CourseList from "../components/CourseList";
import UpcomingAssignment from "../components/UpcomingAssignment";
import AssignmentList from "../components/AssignmentList";
import { FaBars } from "react-icons/fa";
import axios from "axios";
import TitleElement from "../components/TitleElement";
import Assicon from "../icons/ion_list.png";

export default function InstructorDashboard() {
  const icons = [dashicon, noticon, createicon, exiticon];
  const links = ["/insdash", "/notifications", "/create", "/exit"];
  const [isOpen, setIsOpen] = useState(false);
  const [first_name, setFirstName] = useState("");
  const [profile_image, setProfileImage] = useState("");
  //const [courses, setCourses] = useState([]);
  //const [assignments, setAssignments] = useState([]);
  //const [upcomingAssignments, setUpcomingAssignments] = useState([]);
  const [courses, setCourses] = useState<Course[]>([]);
  const [assignments, setAssignments] = useState<Assignment[]>([]);
  const [upcomingAssignments, setUpcomingAssignments] = useState<Assignment[]>([]);
  const [user_group_name, setUserGroup] = useState("");

  interface Course {
    course_id: string;
    course_code: string;
    course_name: string;
    course_color: string;
    course_image: string;
    Assignment: Assignment[];
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

  useEffect(() => {
    const fetchPersonalData = async () => {      
      try {
        const res = await axios.get("/api/QueryPersonDataByUserID");
    
        if (res.data && res.data.user && res.data.user.length > 0) {
          const { user_image_url, user_first_name } = res.data.user[0];
    
          if (user_image_url) setProfileImage(user_image_url);
          if (user_first_name) setFirstName(user_first_name);
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

    const fetchCourses = async () => {
      try {
        const res = await axios.get("/api/QueryCourseByUserID");
        if (res.data) {
          const { courses } = res.data;
          console.log("courses", courses);
          
          if ( courses ) setCourses(courses);
        } else {
          console.warn("No data found in response");
        }
      } catch (error) {
        console.error("Error loading courses:", error);
      }
    };

    const fetchAssignments = async () => {
      try {
        const res = await axios.get("/api/api/QueryAssignmentByUserID");
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
          console.log("assignments", assignments);
          if ( assignments ) setUpcomingAssignments(assignments);
        }else {
          console.warn("No data found in response");
        }
      } catch (error) {
        console.log("Error loading assignments:", error);
      }
    }

    fetchPersonalData();
    fetchPersonalUserGroup();
    fetchCourses();
    fetchAssignments();
    fetchUpComingAssignments();
  }, []);

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

  const assignmentsWithColor = assignments.map((assignment) => {
    const course = courses.find((course) => course.course_id === assignment.course_id);
    const color = course ? course.course_color : "gray";
    const course_name = course ? course.course_name : "Unknown Course";
    return {
      assignment_id: assignment.assignment_id,
      assignment_name: assignment.assignment_name,
      due_date: assignment.due_date,
      color,
      course_name: course_name,
    };
  });

  return (
    <div className="bg-B1 flex items-center min-h-dvh min-w-full">
      <div className="container mx-auto flex lg:flex-row gap-5 p-5">

        {/* Left */}
        <div className="bg-white rounded-2xl flex-1 relative min-h-[900px]">
          <div>
            <LeftMain title={first_name} icon={dashboardicon} />
            <button
              className="absolute right-10 top-10 block xl:hidden"
              onClick={toggleMenu}
            >
              <FaBars size={40} color="#344B59" />
            </button>
          </div>

          <div className="px-4 md:px-6 lg:px-10">
            <div className="mb-4">
              {/*Component Courses */}
              <CourseList courses={courses} user_group_name="instructor" />
            </div>
            <div className="flex flex-col lg:flex-row gap-4">
              <div className="flex-1 p-4 overflow-hidden font-poppins text-E1">
                  <TitleElement name="Assignment" icon={Assicon} />
                {/*Component Assignments list */}
                <AssignmentList Assignment={assignmentsWithColor} /> 
              </div>
              <div className="lg:flex-1">
                {/*Component Upcoming Assignments */}
                <UpcomingAssignment UpcomingAssignment={upcomingAssignmentsWithColor} /> 
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
          {/*Component Right Main */}
          <RightMain icons={icons} links={links} profile_image={profile_image} user_group_name={user_group_name}/>
        </div>
        <div className="hidden xl:block">
          {/*Component Right Main */}
          <RightMain icons={icons} links={links} profile_image={profile_image} user_group_name={user_group_name}/>
        </div>
      </div>
    </div>
  );
}