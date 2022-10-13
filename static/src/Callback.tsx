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
    (async () => {
      await getToken();
      navigate("/", { replace: true });
    })();
  }, []);
  const getToken = () => {
    axios.defaults.withCredentials = true;
    const url = new URL(window.location.href);
    const authorizationCode = url.searchParams.get("code");
    if (authorizationCode) {
      return axios
        .post("http://localhost:4000/callback", { authorizationCode })
        .then((res) => {
          const accessToken = res.data.accessToken;
          const refreshToken = document.cookie
            .split(" ")
            .filter((cookie) => cookie.includes("vegas"))[0]
            .split("=")[1];
          console.log(accessToken, refreshToken);
          props.setToken({ accessToken, refreshToken });
          localStorage.setItem("userid", res.data.userid);
        })
        .catch((err) => {
          console.error(err);
        });
    }
  };
  return (
    <>
      <div>서비스 리디렉션 중입니다...</div>
    </>
  );
}

export default Callback;
