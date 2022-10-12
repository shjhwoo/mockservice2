import { useState } from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import "./App.css";
import Callback from "./Callback";
import Main from "./Main";
import Service from "./Service";

function App() {
  const [accessToken, setAccessToken] = useState<string>("");
  const [refreshToken, setRefreshToken] = useState<string>("");
  return (
    <Router>
      <Routes>
        <Route
          path="/"
          element={
            <Main accessToken={accessToken} refreshToken={refreshToken} />
          }
        />
        <Route
          path="/callback"
          element={
            <Callback
              accessToken={accessToken}
              refreshToken={refreshToken}
              setAccessToken={setAccessToken}
              setRefreshToken={setRefreshToken}
            />
          }
        />
        <Route path="/service" element={<Service />} />
      </Routes>
    </Router>
  );
}

export default App;
