import "./App.css";
import RightMain from "./components/RightMain";
import noticon from "./icons/bxs_bell.png";
import joinicon from "./icons/material-symbols_join.png";
import dashicon from "./icons/mdi_human-welcome.png";
import exiticon from "./icons/vaadin_exit-o.png";

function App() {
  const icons = [dashicon, noticon, joinicon, exiticon];
  const links = ["/dashboard", "/notifications", "/join", "/exit"];

  return (
    <div className="flex justify-center items-center">
      <RightMain icons={icons} links={links} />
    </div>
  );
}

export default App;
