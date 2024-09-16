import "./App.css";
import RightMain from "./components/RightMain";
import noticon from "./icons/bxs_bell.png";
import joinicon from "./icons/material-symbols_join.png";
import dashicon from "./icons/mdi_human-welcome.png";
import exiticon from "./icons/vaadin_exit-o.png";
import InstructorAssignment from "./components/InstructorAssignment.tsx";
import AssigmentSubmit from "./components/AssignmentSubmit.tsx";
import AssignmentButton from "./components/AssignmentButton.tsx";
import AssignmentDetail from "./components/AssignmentDetail.tsx";
import AssignmentSubmitted from "./components/AssignmentSubmitted.tsx";
import UpcomingAssignment from "./components/UpcomingAssignment.tsx";
import UpcomingElement from "./components/UpcomingElement.tsx";

function App() {
  const icons = [dashicon, noticon, joinicon, exiticon];
  const links = ["/dashboard", "/notifications", "/join", "/exit"];

  const ConfigAssignment = [
    { StdCode: "640610629", Status: "#E61616" },
    { StdCode: "640629042", Status: "#E61616" },
    { StdCode: "633934788", Status: "#E61616" },
    { StdCode: "634124894", Status: "#E61616" },
    { StdCode: "123408895", Status: "#E61616" },
    { StdCode: "640610629", Status: "#E61616" },
    { StdCode: "640629042", Status: "#E61616" },
    { StdCode: "633934788", Status: "#E61616" },
    { StdCode: "634124894", Status: "#E61616" },
    { StdCode: "123408895", Status: "#E61616" },
    { StdCode: "640610629", Status: "#E61616" },
    { StdCode: "640629042", Status: "#E61616" },
    { StdCode: "633934788", Status: "#E61616" },
    { StdCode: "634124894", Status: "#E61616" },
    { StdCode: "123408895", Status: "#E61616" },
  ];

  return (
    <div>
      <div className="flex justify-center items-center">
        <RightMain icons={icons} links={links} />
      </div>

      <div>
        <InstructorAssignment filename="File" />
      </div>

      <div>
        <AssigmentSubmit StdCode="640610629" Status="#E61616" />
      </div>

      <div>
        <AssignmentButton color="yellow" text="Submit" />
      </div>

      <div>
        <AssignmentDetail role={"Instructor"} />
      </div>

      <div>
        <AssignmentSubmitted ConfigAssignment={ConfigAssignment} />
      </div>

      <div>
        <UpcomingAssignment percentage="30%" color="purple" timeleft={3} />
      </div>

      {/* <div>
        <UpcomingElement />
      </div> */}
    </div>
  );
}

export default App;
