import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import "./App.css";
import Callback from "./Callback";
import Main from "./Main";
import Service from "./Service";
import SingleLogOut from "./SingleLogOut";
import Nurse from "./Nurse";

import { TokenProvider } from "./contexts/MainContext"; //반드시 컴포넌트명은 대문자!

const App: React.FC = () => {
  return (
    <Router>
      <TokenProvider>
        <Routes>
          <Route path="/" element={<Main />} />
          <Route path="/callback" element={<Callback />} />
          <Route path="/service" element={<Service />} />
          <Route path="/nurse" element={<Nurse />} />
          <Route path="/slo" element={<SingleLogOut />} />
        </Routes>
      </TokenProvider>
    </Router>
  );
};

export default App;
