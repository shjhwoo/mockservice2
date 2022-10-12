import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import "./App.css";
import Callback from "./Callback";
import Main from "./Main";
import Service from "./Service";
import SingleLogOut from "./SingleLogOut"

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Main />} />
        <Route path="/callback" element={<Callback />} />
        <Route path="/service" element={<Service />} />
        <Route path="/slo" element={<SingleLogOut />} />
      </Routes>
    </Router>
  );
}

export default App;
