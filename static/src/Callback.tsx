import { useEffect, Dispatch, SetStateAction } from "react";
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

axios.defaults.withCredentials = true;

function Callback(props: Props) {
  const navigate = useNavigate();
  useEffect(() => {
    getToken();
    navigate("/service", { replace: true });
  }, []);
  const getToken = () => {
    axios.defaults.withCredentials = true;
    const url = new URL(window.location.href);
    const authorizationCode = url.searchParams.get("code");
    if (authorizationCode) {
      axios
        .post(
          "http://localhost:4000/callback",
          { authorizationCode },
          { headers: { withCredentials: true } }
        )
        .then((res) => {
          const accessToken = res.data.accessToken;
          const refreshToken = document.cookie
            .split(" ")
            .filter((cookie) => cookie.includes("vegas"))[0];
          console.log(accessToken, refreshToken);
          props.setToken({ accessToken, refreshToken });
        })
        .catch((err) => {
          console.error(err);
        });
    }
  };
  useEffect(() => {
    console.log(props.token);
  }, [props.token]);
  return (
    <>
      <div>서비스 리디렉션 중입니다...</div>
    </>
  );
}

export default Callback;
