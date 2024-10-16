import React, { useEffect, useState } from "react";
import Background from "../img/SignupBack.png";
import Logo from "../img/Logo.png";
import { useLocation } from "react-router-dom";
import { jwtDecode } from "jwt-decode";

export default function SignUpPage() {
  const location = useLocation();
  const role = location.state?.role || localStorage.getItem("role") || "Role";
  const [email, setEmail] = useState("");
  const [user_name, setUsername] = useState("");
  const [first_name, setFirstName] = useState("");
  const [last_name, setLastName] = useState("");
  const [password, setPassword] = useState("");
  const [profile_image_url, setImg] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();

    if (password !== confirmPassword) {
      alert("Password and Confirm Password do not match!");
      return;
    }

    let roleValue = 0;
    console.log("Role received: ", role);

    if (role === "Student") {
      roleValue = 1;
    } else if (role === "Instructor") {
      roleValue = 2;
    }

    const formData = {
      email,
      user_name,
      first_name,
      last_name,
      password,
      profile_image_url,
      group_id: roleValue,
    };
    console.log(formData);

    try {
      const response = await fetch("/api/CreateUser", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        console.log("Form submitted successfully");
        window.location.href = "/landing";
      } else {
        console.error("Form submission failed");
      }
    } catch (error) {
      console.error("Error submitting form:", error);
    }
  };

  useEffect(() => {
    const urlParams = new URLSearchParams(location.search);
    const token = urlParams.get("token");
    if (role !== "Role") {
      localStorage.setItem("role", role);
    }

    if (token) {
      try {
        const decodedUser = jwtDecode(token);

        setEmail((decodedUser as { email?: string }).email || "");
        setFirstName((decodedUser as { firstName?: string}).firstName || "");
        setLastName((decodedUser as {lastName?: string }).lastName || "");
        setImg((decodedUser as {picture?: string}).picture || "");
        console.log("Decoded User:", decodedUser);
      } catch (error) {
        console.error("Failed to decode token:", error);
      }
    } else {
      console.error("Token is missing");
    }
  }, [location]);

  const handleClick = async () => {
    localStorage.setItem("role", role);
    window.location.href = "api/auth/google/login";
  };


  return (
    <div
      className="bg-white h-screen flex items-center font-poppins text-E1"
      style={{
        backgroundImage: `url(${Background})`,
        backgroundSize: "cover",
        backgroundRepeat: "no-repeat",
      }}
    >
      <div className="container mx-auto bg-white flex flex-col lg:flex-row  justify-center min-h-screen lg:min-h-[825px] rounded-3xl shadow-xl ">
        {/* left */}
        <div className="flex-auto flex justify-center items-center flex-col gap-10">
          <div className="font-medium text-5xl">Sign up</div>
          <div className="flex flex-col lg:flex-row items-center gap-3">
            {/* Google button Authen */}
            <div className="px-6 sm:px-0 max-w-sm">
              <button
                type="button"
                className="text-white w-full  bg-[#4285F4] hover:bg-[#4285F4]/90 focus:ring-4 focus:outline-none focus:ring-[#4285F4]/50 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center justify-between mr-2 mb-2"
                onClick={handleClick}
              >
                <svg
                  className="mr-2 -ml-1 w-4 h-4"
                  aria-hidden="true"
                  focusable="false"
                  data-prefix="fab"
                  data-icon="google"
                  role="img"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 488 512"
                >
                  <path
                    fill="currentColor"
                    d="M488 261.8C488 403.3 391.1 504 248 504 110.8 504 0 393.2 0 256S110.8 8 248 8c66.8 0 123 24.5 166.3 64.9l-67.5 64.9C258.5 52.6 94.3 116.6 94.3 256c0 86.5 69.1 156.6 153.7 156.6 98.2 0 135-70.4 140.8-106.9H248v-85.3h236.1c2.3 12.7 3.9 24.9 3.9 41.4z"
                  ></path>
                </svg>
                Sign up with Google<div></div>
              </button>
            </div>
            <div className="font-medium text-xl text-center">as a {role}</div>
          </div>
          <div>
            <form
              className="px-8 text-xl flex flex-col gap-5"
              onSubmit={handleSubmit}
            >
              <div>
                <label className="flex justify-start mb-2" htmlFor="email">
                  Email
                </label>
                <input
                  className="shadow border border-G1 rounded-lg w-full py-2 px-3 text-gray-600"
                  id="email"
                  type="email"
                  value={email}
                  readOnly
                  onChange={(e) => setEmail(e.target.value)}
                />
              </div>
              <div>
                <label className="flex justify-start mb-2" htmlFor="username">
                  Username
                </label>
                <input
                  className="shadow border border-G1 rounded-lg w-full py-2 px-3 text-gray-600"
                  id="username"
                  type="text"
                  onChange={(e) => setUsername(e.target.value)}
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
                    value={first_name}
                    onChange={(e) => setFirstName(e.target.value)}
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
                    value={last_name}
                    onChange={(e) => setLastName(e.target.value)}
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
                  type="password"
                  onChange={(e) => setPassword(e.target.value)}
                />
              </div>
              <div>
                <label className="flex justify-start mb-2" htmlFor="ConPass">
                  Confirm Password
                </label>
                <input
                  className="shadow border border-G1 rounded-lg w-full py-2 px-3 text-gray-600"
                  id="ConPass"
                  type="password"
                  onChange={(e) => setConfirmPassword(e.target.value)}
                />
              </div>
              {/* Submit button */}
              <div className="flex justify-center">
                <button className="w-[200px] h-[55px] bg-M1 text-white rounded-full hover:bg-blue-500 text-2xl">
                  Submit
                </button>
              </div>
            </form>
          </div>
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