import { useEffect, Dispatch, SetStateAction, useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

interface token {
  accessToken: string;
  refreshToken: string;
}

interface Props {
  token: token;
  setToken: Dispatch<SetStateAction<token>>;
}

const fetchToken = async () => {
  try {
    const url = new URL(window.location.href);
    const authorizationCode = url.searchParams.get("code");
    if (authorizationCode) {
      const resp = await axios.post(
        "http://localhost:4000/callback",
        { authorizationCode },
        { withCredentials: true }
      );
      const accessToken = resp.data.accessToken;
      const refreshToken = document.cookie
        .split(" ")
        .filter((item) => item.includes("vegasRefreshToken"))[0]
        .split("=")[1];
      return { accessToken, refreshToken };
    }
  } catch (e) {
    console.error(e);
  }
};

function Callback(props: Props) {
  const [tkn, setTkn] = useState<token>({ accessToken: "", refreshToken: "" });
  const navigate = useNavigate();
  useEffect(() => {
    const getToken = async () => {
      const token = await fetchToken();
      // if (token !== undefined) props.setToken(token);
      if (token !== undefined) setTkn(token);
    };
    getToken();
    navigate("/service", { replace: true });
  }, []);
  return (
    <>
      <div>서비스 리디렉션 중입니다...</div>
    </>
  );
}

export default Callback;

// try {
//   const url = new URL(window.location.href);
//   const authorizationCode = url.searchParams.get("code");
//   if (authorizationCode) {
//     const resp = await axios.post(
//       "http://localhost:4000/callback",
//       { authorizationCode },
//       { withCredentials: true }
//     );
//     const accessToken = resp.data.accessToken;
//     const refreshToken = document.cookie
//       .split(" ")
//       .filter((item) => item.includes("vegasRefreshToken"))[0]
//       .split("=")[1];
//     props.setToken({ accessToken, refreshToken });
//     window.localStorage.setItem("userid", resp.data.userid);
//   }
// } catch (e) {
//   console.error(e);
// }
