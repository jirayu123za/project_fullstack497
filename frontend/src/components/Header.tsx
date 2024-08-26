import React from "react";
import { useState, useEffect } from "react";
import Logo from "../img/Logo.png";

export default function Header() {
  const [isActive, setActive] = useState(false);

  useEffect(() => {
    const handleScroll = () => {
      setActive(window.scrollY > 60);
    };

    window.addEventListener("scroll", handleScroll);

    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, []);

  return (
    <div
      className={`${
        isActive ? "shadow-md" : ""
      } bg-white text-center py-5 fixed w-full transition-all z-10`}
    >
      <div className="container mx-auto flex justify-between items-center">
        <a href="#">
          <img src={Logo} alt="" />
        </a>
        <div className="flex font-semibold gap-4 md:gap-7 text-xl items-center">
          <a
            href="#"
            className="text-M1 flex items-center justify-center text-center"
          >
            About Us
          </a>
          <button className="w-full md:w-[163px] h-[57px] bg-E1 text-white rounded-full flex items-center justify-center hover:bg-blue-500">
            Sign Up
          </button>
        </div>
      </div>
    </div>
  );
}
