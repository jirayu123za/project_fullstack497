import React, { useState } from "react";
//import axios from "axios";
import Background from "../img/SignupBack.png";
import { IoArrowBackCircle } from "react-icons/io5";
import { useNavigate } from "react-router-dom";

export default function INS_Create() {
  const [course_name, setCourseName] = useState("");
  const [course_code, setCourseCode] = useState("");
  //const [term, setTerm] = useState("");
  const [image_url, setImageUrl] = useState("");
  const [color, setColor] = useState("Purple");
  const navigate = useNavigate();

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();

    const formData = {
      course_name,
      course_code,
      //term,
      image_url,
      color,
    };
    console.log(formData);

    try {
      const response = await fetch("api/api/CreateCourse", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ course_name, course_code, /*term,*/ image_url, color
         }),
      });
      console.log(course_name, course_code, /*term,*/ image_url, color);
      if (response.ok) {
        const result = await response.json();
        console.log("POST API success: ", result);
        navigate("/insdash");
      } else {
        console.log("POST API failed: ", response);
      }
    } catch (error) {
      console.error("Error during POST API: ", error);
    }
  };

  return (
    <div
      className="flex justify-center items-center h-screen font-poppins"
      style={{
        backgroundImage: `url(${Background})`,
        backgroundSize: "cover",
        backgroundRepeat: "no-repeat",
      }}
    >
      <div className="w-full max-w-4xl mx-auto px-4">
        <form
          className="bg-white border-4 border-B1 shadow-lg rounded-lg px-6 pt-10 pb-8 mb-4"
          onSubmit={handleSubmit} // เรียก handleSubmit เมื่อกด submit
        >
          {/* Course Name */}
          <div className="mb-4 flex flex-col sm:flex-row gap-5 items-center">
            <label
              className="block text-gray-700 text-xl mb-2 sm:mb-0 sm:w-1/4"
              htmlFor="course_name"
            >
              Course Name:
            </label>
            <input
              className="shadow appearance-none border rounded w-full sm:w-3/4 py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="course_name"
              type="text"
              value={course_name}
              onChange={(e) => setCourseName(e.target.value)}
            />
          </div>

          {/* Course Code */}
          <div className="mb-4 flex flex-col sm:flex-row gap-5 items-center">
            <label
              className="block text-gray-700 text-xl mb-2 sm:mb-0 sm:w-1/4"
              htmlFor="course_code"
            >
              Course Code:
            </label>
            <input
              className="shadow appearance-none border rounded w-full sm:w-3/4 py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="course_code"
              type="text"
              value={course_code}
              onChange={(e) => setCourseCode(e.target.value)}
            />
          </div>

          {/* Image URL */}
          <div className="mb-4 flex flex-col sm:flex-row gap-5 items-center">
            <label
              className="block text-gray-700 text-xl mb-2 sm:mb-0 sm:w-1/4"
              htmlFor="img_url"
            >
              Cover image url:
            </label>
            <input
              className="shadow appearance-none border rounded w-full sm:w-3/4 py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="img_url"
              type="text"
              value={image_url}
              onChange={(e) => setImageUrl(e.target.value)} // เก็บค่า URL ของรูปภาพ
            />
          </div>

          {/* pick color for display */}
          <div className="mb-4 flex flex-col sm:flex-row gap-5 items-center">
            <label
              className="block text-gray-700 text-xl mb-2 sm:mb-0 sm:w-1/4"
              htmlFor="pickcolor"
            >
              Pick color:
            </label>
            <select
              className="shadow appearance-none border rounded w-full sm:w-3/4 py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="pickcolor"
              value={color}
              onChange={(e) => setColor(e.target.value)}
            >
              <option value="purple">Purple</option>
              <option value="yellow">Yellow</option>
              <option value="green">Green</option>
              <option value="red">Red</option>
              <option value="pink">Pink</option>
              <option value="blue">Blue</option>
              <option value="orange">Orange</option>
              <option value="brown">Brown</option>
            </select>
          </div>

          {/* Buttons */}
          <div className="flex flex-col sm:flex-row items-center justify-end gap-5">
            <button
              className="bg-R4 hover:bg-R4 hover:bg-opacity-80 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full sm:w-auto"
              type="submit"
            >
              Create
            </button>
          </div>
        </form>
      </div>

      <div className="absolute top-5 left-5">
        <a href="/insdash">
          <IoArrowBackCircle size={60} color="#344B59" />
        </a>
      </div>
    </div>
  );
}
