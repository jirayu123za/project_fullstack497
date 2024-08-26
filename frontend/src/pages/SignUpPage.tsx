import React from "react";
import Background from "../img/SignupBack.png";
import Logo from "../img/Logo.png";
import { useLocation } from "react-router-dom";

export default function SignUpPage() {
  const location = useLocation();
  const role = location.state?.role || "Role";

  return (
    <div
      className="bg-white h-screen flex items-center font-poppins text-E1"
      style={{
        backgroundImage: `url(${Background})`,
        backgroundSize: "cover",
        backgroundRepeat: "no-repeat",
      }}
    >
      <div className="container mx-auto bg-white flex flex-col lg:flex-row  justify-center min-h-[825px] rounded-3xl shadow-xl">
        {/* left */}
        <div className="flex-auto flex justify-center items-center flex-col gap-10">
          <div className="font-medium text-5xl text-center">Sign Up</div>
          <div className="bg-B1 text-white font-medium text-2xl  min-w-60 lg:min-w-72 min-h-14  rounded text-center flex justify-center items-center">
            {role}
          </div>
          <div>
            <form className="px-8 pt-6 mb-4 text-xl flex flex-col gap-5">
              <div>
                <label className="flex justify-start mb-2" htmlFor="username">
                  Username
                </label>
                <input
                  className="shadow border border-G1 rounded-lg w-full py-2 px-3 text-gray-600"
                  id="username"
                  type="text"
                />
              </div>
              <div className="flex gap-2">
                <div>
                  <label
                    className="flex justify-start mb-2"
                    htmlFor="Firstname"
                  >
                    First Name
                  </label>
                  <input
                    className="shadow border border-G1 rounded-lg w-full py-2 px-3 text-gray-600"
                    id="Firstname"
                    type="text"
                  />
                </div>
                <div>
                  <label className="flex justify-start mb-2" htmlFor="Lastname">
                    Last Name
                  </label>
                  <input
                    className="shadow border border-G1 rounded-lg w-full py-2 px-3 text-gray-600"
                    id="Lastname"
                    type="text"
                  />
                </div>
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
              <div>
                <label className="flex justify-start mb-2" htmlFor="ConPass">
                  Confirm Password
                </label>
                <input
                  className="shadow border border-G1 rounded-lg w-full py-2 px-3 text-gray-600"
                  id="ConPass"
                  type="text"
                />
              </div>
            </form>
          </div>
          <button className="w-[200px] h-[55px] bg-M1 text-white rounded-full hover:bg-blue-500 text-2xl">
            Sign Up
          </button>
        </div>
        {/* right */}
        <div className="flex-auto bg-[#F8F7F7] flex flex-col justify-center items-center rounded-b-3xl lg:rounded-r-3xl lg:rounded-l-none mt-8 pt-8 lg:mt-0 lg:pt-0">
          <div>
            <img src={Logo} alt="" className="max-w-[300px]" />
          </div>
          <div className="max-w-[300px] mt-12 flex flex-col gap-10">
            <div>
              Your password should be at least 8 characters long and include a
              combination of uppercase and lowercase letters, numbers, and
              special characters . (e.g., !, @, #, $).
            </div>
            <div>
              Please create a password that is at least 8 characters long to
              enhance security.
            </div>
            <div>
              Your username can include letters, numbers, and special characters
              (-, _, .), but it cannot start or end with a special character.
            </div>
            <div>
              The email must be valid as you will need to verify it during the
              registration process.
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
