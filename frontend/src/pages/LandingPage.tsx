import React from "react";
import LoginSection from "../components/LoginSection";
import Aos from "aos";
import "aos/dist/aos.css";
import SelectRoleSignup from "../components/SelectRoleSignup";
import AboutUsSec from "../components/AboutUsSec";
import { useState, useEffect } from "react";
import Logo from "../img/Logo.png";

export default function LandingPage() {
  Aos.init({
    duration: 1800,
    offset: 100,
  });

  const scrollToSection = (sectionId: string) => {
    const section = document.getElementById(sectionId);
    if (section) {
      section.scrollIntoView({ behavior: "smooth" });
    }
  };

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
    <div className="overflow-hidden">
      <div
        className={`${
          isActive ? "shadow-lg" : ""
        } bg-white text-center py-5 fixed w-full transition-all z-10`}
      >
        <div className="container mx-auto flex lg:justify-between items-center justify-center">
          <a href="#">
            <img src={Logo} alt="" />
          </a>
          <div className="lg:flex font-semibold gap-4 md:gap-7 text-xl items-center hidden">
            <a
              href="#"
              className="text-M1 flex items-center justify-center text-center"
              onClick={() => scrollToSection("section3")}
            >
              About Us
            </a>
            <button
              className="w-full md:w-[163px] h-[57px] bg-E1 text-white rounded-full flex items-center justify-center hover:bg-blue-500"
              onClick={() => scrollToSection("section2")}
            >
              Sign Up
            </button>
          </div>
        </div>
      </div>
      <LoginSection />
      <div id="section2">
        <SelectRoleSignup />
      </div>
      <div id="section3">
        <AboutUsSec />
      </div>
    </div>
  );
}
