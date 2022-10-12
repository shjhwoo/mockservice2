import { useState } from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import "./App.css";
import Callback from "./Callback";
import Main from "./Main";
import Service from "./Service";
import SingleLogOut from "./SingleLogOut";

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
        <Route path="/callback" element={<Callback token={token} setToken={setToken} />} />
        <Route path="/service" element={<Service />} />
        <Route path="/slo" element={<SingleLogOut />} />
      </Routes>
    </Router>
  );
}

export default App;
