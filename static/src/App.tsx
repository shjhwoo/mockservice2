import { useState } from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import "./App.css";
import Callback from "./Callback";
import Main from "./Main";
import Service from "./Service";
import SingleLogOut from "./SingleLogOut";
import api from "./Api";
import Nurse from "./Nurse";

interface token {
  accessToken: string;
  refreshToken: string;
}

function App() {
  const [token, setToken] = useState<token>({
    accessToken: "",
    refreshToken: "",
  });
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Main token={token} />} />
        <Route
          path="/callback"
          element={<Callback token={token} setToken={setToken} />}
        />
        <Route path="/service" element={<Service token={token} />} />
        <Route path="/nurse" element={<Nurse token={token} />} />
        <Route path="/slo" element={<SingleLogOut token={token} />} />
      </Routes>
    </Router>
  );
}

export default App;
