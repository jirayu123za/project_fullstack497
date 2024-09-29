import book from "../img/Books.png";
import clock from "../img/Clock.png";
import calender from "../img/Calendar.png";

export default function AboutUsSec() {
  return (
    <section
      className="min-h-[800px] py-6 font-poppins text-E1"
      data-aos="fade-up"
      data-aos-delay="200"
    >
      <div className="container mx-auto border-8 border-B1 rounded-t-[200px] p-20">
        <div className="flex justify-center pt-10 font-medium text-5xl text-E1">
          About Us
        </div>
        <div className="flex justify-around pt-20 space-x-10 lg:flex-row flex-col">
          <div className="flex flex-col items-center">
            <img src={book} alt="" />
          </div>
          <div className="flex flex-col items-center">
            <img src={clock} alt="" />
          </div>
          <div className="flex flex-col items-center">
            <img src={calender} alt="" />
          </div>
        </div>
        <div>
          <p className="text-3xl mt-12 font-semibold text-center">
            can manage assignments efficiently by prioritizing tasks, setting
            clear deadlines, and using organizational tools to track progress."
          </p>
        </div>
      </div>
    </section>
  );
}
