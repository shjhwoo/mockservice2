import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import "./App.css";
import Callback from "./Callback";
import Main from "./Main";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Main />} />
        <Route path="/callback" element={<Callback />} />
      </Routes>
    </Router>
  );
}

export default App;
