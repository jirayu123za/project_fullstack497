import React, { useState, useEffect } from "react";
import axios from "axios";
import stdicon from "../icons/tdesign_member.png";
import { FiX } from "react-icons/fi";
import TitleElement from "./TitleElement";
import { FaPlusSquare } from "react-icons/fa";
import { useParams } from "react-router-dom";

interface Student {
  user_id: string;
  first_name: string;
  last_name: string;
  profile_image: string;
}

interface StudentListProps {
  user_group_name: string;
}

const StudentList: React.FC<StudentListProps> = ({ user_group_name }) => {
  const { course_id } = useParams();
  const [Students, setStudents] = useState<Student[]>([]);
  const [email, setEmail] = useState("");
  const [isPopupOpen, setIsPopupOpen] = useState(false);
  const [errorMessage, setErrorMessage] = useState<string | null>(null);

  useEffect(() => {
    fetchStudentsList();
  }, []);

  const fetchStudentsList = async () => {
    try {
      const res = await axios.get(`/api/api/QueryUsersEnrollment?course_id=${course_id}`);
      if (res.data && res.data.users) {
        const studentsData = res.data.users.map((user: any) => ({
          user_id: user.user_id,
          first_name: user.first_name,
          last_name: user.last_name,
          profile_image: user.profile_image_url,
        }));

        setStudents(studentsData);
        console.log("studentsData", studentsData);
      } else {
        console.warn("No data found in response");
      }
    } catch (error) {
      console.error("Error fetching students:", error);
    }
  };

  const handleRemoveStudent = async (user_id: string) => {
    try {
      console.log("Call func handleRemoveStudent with user_id:", user_id);
      
      const res = await axios.delete(`/api/api/DeleteUserEnrollment?course_id=${course_id}&user_id=${user_id}`);
      if (res.status === 200) {
        await fetchStudentsList();
        console.log("Student removed successfully");
      } else {
        setErrorMessage("Failed to remove student.");
      }
    } catch (error) {
      console.error("Error removing student:", error);
    }
  }

  const handleAddStudent = async () => {
    try {
      const res = await axios.post(`/api/api/CreateEnrollment?course_id=${course_id}`, { email });

      if (res.status === 201) {
        await fetchStudentsList();
        setIsPopupOpen(false);
        setEmail("");
        setErrorMessage(null);
        console.log("Student added successfully");
      } else {
        setErrorMessage("Failed to add student.");
      }
    } catch (error) {
      console.error("Error adding student:", error);
      setErrorMessage("An error occurred while adding the student.");
    }
  };

  return (
    <div className="max-w-full">
      <div className="flex items-center mb-4 gap-3">
        <TitleElement name={"Student List"} icon={stdicon} />
        {user_group_name !== "STUDENT" && (
          <div onClick={() => setIsPopupOpen(true)}>
            <FaPlusSquare size={40} color="#93B955" />
          </div>
        )}
      </div>
      <div className="overflow-y-auto max-h-[550px] scrollbar-hide w-full">
        <ul className="divide-y divide-gray-200">
          {Array.isArray(Students) &&
            Students.map((student) => (
              <li
                key={student.user_id}
                className="flex items-center justify-between p-2"
              >
                <div className="flex items-center">
                  <img
                    src={student.profile_image}
                    alt={student.first_name + " " + student.last_name}
                    className="w-10 h-10 rounded-full mr-10 object-cover"
                  />
                  <p className="text-lg">{student.first_name + " " + student.last_name}</p>
                </div>
                <FiX
                  className="text-red-500 cursor-pointer"
                  onClick={() => handleRemoveStudent(student.user_id)}
                />
              </li>
            ))}
        </ul>
      </div>

      {/* Popup Add student */}
      {isPopupOpen && (
        <div className="fixed inset-0 bg-gray-600 bg-opacity-75 flex justify-center items-center z-50">
          <div className="bg-white p-8 rounded-lg shadow-lg w-full max-w-md">
            <h2 className="text-xl mb-4">Add New Student</h2>
            <input
              type="email"
              className="border p-2 w-full mb-4"
              placeholder="Enter student email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
            {errorMessage && <p className="text-red-500">{errorMessage}</p>}
            <div className="flex justify-end gap-2">
              <button
                className="bg-red-500 text-white px-4 py-2 rounded"
                onClick={() => setIsPopupOpen(false)}
              >
                Cancel
              </button>
              <button
                className="bg-R4 text-white px-4 py-2 rounded"
                onClick={() => handleAddStudent()}
              >
                Add Student
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default StudentList;
