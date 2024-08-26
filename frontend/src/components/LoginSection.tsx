import React from "react";

export default function LoginSection() {
  return (
    <section className="bg-white min-h-[800px] py-6 font-poppins mt-8 lg:mt-0">
      <div className="container mx-auto min-h-[800px] flex justify-center items-center">
        <div className="flex flex-col gap-y-8 justify-center items-center lg:flex-row lg:gap-x-8 lg:gap-y-0">
          {/* Login */}
          <div className="flex-1 relative">
            <div className="mt-16 size-[568px] border-8 border-B1 bg-white rounded-xl">
              <div className="bg-B1 w-[552px] h-24 text-M1 font-bold text-3xl flex justify-start items-center">
                <div className="ml-7">Login</div>
              </div>
              {/* Username and Input form */}
              <form className="px-8 pt-6 mb-4 text-M1 text-xl font-medium">
                <div className="mb-11 mt-10">
                  <label className="flex justify-start mb-2" htmlFor="username">
                    Username
                  </label>
                  <input
                    className="shadow border border-G1 rounded-lg w-full py-2 px-3 text-gray-600"
                    id="username"
                    type="text"
                  />
                </div>
                <div>
                  <label className="flex justify-start mb-2" htmlFor="password">
                    Password
                  </label>
                  <input
                    className="shadow border border-G1 rounded-lg w-full py-2 px-3 text-gray-600"
                    id="password"
                    type="text"
                  />
                </div>
                <div className="flex items-center justify-end mt-5">
                  <a
                    className="font-medium text-sm text-M1 hover:text-blue-500"
                    href="#"
                  >
                    Forgot Password?
                  </a>
                </div>
                <div className="flex items-center justify-center mt-16">
                  <button className="w-[233px] h-[43px] bg-M1 text-white rounded-full hover:bg-blue-500">
                    Login
                  </button>
                </div>
              </form>
            </div>
          </div>
          {/* Intro Text */}
          <div
            className="flex-1 text-center lg:text-left text-E1 "
            data-aos="fade-down"
            data-aos-delay="500"
          >
            <h1 className="font-bold text-5xl mb-4 ">
              We provide a comprehensive solution for managing your educational
              institution's needs.
            </h1>
            <p className="mb-4 ">
              Our platform is designed to streamline administrative tasks,
              enhance communication, and improve the overall learning
              experience.
            </p>
          </div>
        </div>
      </div>
    </section>
  );
}
