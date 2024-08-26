import React from "react";
import img from "../../src/img/BrazucaPlanning.png";
import { useNavigate } from "react-router-dom";

export default function SelectRoleSignup() {
  const navigate = useNavigate();
  const SignUpAsIns = () => {
    const role = "Instructor";
    navigate("/signup", { state: { role } });
  };
  const SignUpAsStd = () => {
    const role = "Student";
    navigate("/signup", { state: { role } });
  };

  return (
    <section className="bg-white min-h-[800px] py-6 font-poppins">
      <div className="container mx-auto min-h-[800px] flex justify-center items-center">
        <div className="flex flex-col gap-y-8 justify-center items-center lg:flex-row lg:gap-x-10 lg:gap-y-0">
          {/* Select Role */}
          <div
            className="flex-1 font-medium text-5xl text-E1 min-w-[752px]"
            data-aos="fade-right"
            data-aos-delay="200"
          >
            <div className="my-5 flex justify-center text-center font-bold">
              Sign Up
            </div>
            <div className="text-2xl flex justify-center text-center mb-5">
              What is your role ?
            </div>
            <div className="space-y-3 flex flex-col justify-center items-center lg:flex-row lg:space-y-0 lg:space-x-5">
              <button
                className="bg-B1 hover:bg-blue-500 text-white font-medium text-2xl w-72 h-14 rounded"
                onClick={SignUpAsIns}
              >
                Instructor
              </button>
              <button
                className="bg-B1 hover:bg-blue-500 text-white font-medium text-2xl w-72 h-14 rounded"
                onClick={SignUpAsStd}
              >
                Student
              </button>
            </div>
          </div>
          {/* img */}
          <div className="flex-1 min-w-[752px] flex justify-center lg:justify-end">
            <img
              src={img}
              alt=""
              className="h-[568px]"
              data-aos="fade-left"
              data-aos-delay="200"
            />
          </div>
        </div>
      </div>
    </section>
  );
}
