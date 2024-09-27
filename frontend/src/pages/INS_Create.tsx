import React, { useState } from "react";
import axios from "axios"; // เพิ่ม axios สำหรับการเรียก API
import Background from "../img/SignupBack.png";
import { IoArrowBackCircle } from "react-icons/io5";
import { useNavigate } from "react-router-dom"; // Import useNavigate

export default function INS_Create() {
  const [coursename, setCoursename] = useState("");
  const [coursecode, setCoursecode] = useState("");
  const [ImageUrl, setImageUrl] = useState("");
  const [color, setColor] = useState("Purple");
  const navigate = useNavigate(); // สร้าง navigate สำหรับ redirect

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();

    const formData = {
      coursename,
      coursecode,
      ImageUrl,
      color,
    };

    try {
      // ส่งข้อมูลไปที่ API
      // const response = await axios.post("/api/createcourse", formData);
      // console.log("Course created successfully:", response.data);
      console.log(formData);

      // ถ้าสำเร็จ ทำการ redirect ไปที่ /dashboard
      navigate("/insdash");
    } catch (error) {
      console.error("Error creating course:", error);
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
              htmlFor="coursename"
            >
              Course Name:
            </label>
            <input
              className="shadow appearance-none border rounded w-full sm:w-3/4 py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="coursename"
              type="text"
              value={coursename}
              onChange={(e) => setCoursename(e.target.value)} // เก็บค่าชื่อ course
            />
          </div>

          {/* Course Code */}
          <div className="mb-4 flex flex-col sm:flex-row gap-5 items-center">
            <label
              className="block text-gray-700 text-xl mb-2 sm:mb-0 sm:w-1/4"
              htmlFor="coursecode"
            >
              Course Code:
            </label>
            <input
              className="shadow appearance-none border rounded w-full sm:w-3/4 py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="coursecode"
              type="text"
              value={coursecode}
              onChange={(e) => setCoursecode(e.target.value)} // เก็บค่า course code
            />
          </div>

          {/* Image URL */}
          <div className="mb-4 flex flex-col sm:flex-row gap-5 items-center">
            <label
              className="block text-gray-700 text-xl mb-2 sm:mb-0 sm:w-1/4"
              htmlFor="imgurl"
            >
              Cover image url:
            </label>
            <input
              className="shadow appearance-none border rounded w-full sm:w-3/4 py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="imgurl"
              type="text"
              value={ImageUrl}
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
              value={color} // state สำหรับเก็บค่าที่ถูกเลือก
              onChange={(e) => setColor(e.target.value)} // ฟังก์ชันสำหรับอัพเดท state
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
