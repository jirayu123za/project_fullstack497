import React, { useState } from "react";
import Background from "../img/SignupBack.png";

export default function INS_Create() {
  const [course_name, setCoursename] = useState("");
  const [course_code, setCoursecode] = useState("");
  const [term, setTerm] = useState("");

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();

    const formData = {
      course_name,
      course_code,
      term,
    };
    console.log(formData);

    try {
      const response = await fetch("api/api/CreateCourse", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ course_name, course_code, term }),
      });
      console.log(course_name, course_code, term);
      if (response.ok) {
        const result = await response.json();
        console.log("POST API success: ", result);
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
      <div className="w-full max-w-3xl mx-auto px-4">
        <form
          className="bg-white border-4 border-B1 shadow-lg rounded-lg px-6 pt-10 pb-8 mb-4"
          onSubmit={handleSubmit}
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
              onChange={(e) => setCoursename(e.target.value)}
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
              onChange={(e) => setCoursecode(e.target.value)}
            />
          </div>

          {/* Join Code */}
          <div className="mb-4 flex flex-col sm:flex-row gap-5 items-center">
            <label
              className="block text-gray-700 text-xl mb-2 sm:mb-0 sm:w-1/4"
              htmlFor="Term"
            >
              Term:
            </label>
            <input
              className="shadow appearance-none border rounded w-full sm:w-3/4 py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="Term"
              type="text"
              onChange={(e) => setTerm(e.target.value)}
            />
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
    </div>
  );
}
