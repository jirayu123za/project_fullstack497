import React, { useState } from "react";
import axios from "axios";
import Background from "../img/SignupBack.png";
import { IoArrowBackCircle } from "react-icons/io5";
import { useNavigate } from "react-router-dom";

export default function STD_Join() {
  const [join_code, setJoin] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();

    const code = { join_code,
    };

    try {
      const resp = await axios.post("/api/JoinCourse",{
        method: "POST",
        headers: {
        "Content-Type": "application/json",
      },
        body: JSON.stringify({ join_code }),
      });
      console.log("You call func join course with code: ",code);

      if (resp.status === 200) {
        const result = resp.data;
        console.log("POST API success: ", result);
        navigate("/stddash");
      } else {
        console.log("POST API failed: ", resp);
      }
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
          onSubmit={handleSubmit}
        >
          {/* Course Name */}
          <div className="mb-4 flex flex-col sm:flex-row gap-5 items-center">
            <label
              className="block text-gray-700 text-xl mb-2 sm:mb-0 "
              htmlFor="join"
            >
              Join code :
            </label>
            <input
              className="shadow appearance-none border rounded w-full sm:w-3/4 py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="join"
              type="text"
              onChange={(e) => setJoin(e.target.value)}
            />
          </div>

          {/* Buttons */}
          <div className="flex flex-col sm:flex-row items-center justify-end gap-5">
            <button
              className="bg-R4 hover:bg-R4 hover:bg-opacity-80 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full sm:w-auto"
              type="submit"
            >
              Join
            </button>
          </div>
        </form>
      </div>

      <div className="absolute top-5 left-5">
        <a href="/stddash">
          <IoArrowBackCircle size={60} color="#344B59" />
        </a>
      </div>
    </div>
  );
}
