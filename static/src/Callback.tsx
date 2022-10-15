import { useEffect } from "react";
import {
  SAVE,
  useSampleState,
  useSampleDispatch,
} from "./contexts/MainContext";
import { useNavigate } from "react-router-dom";
import axios from "axios";

axios.defaults.withCredentials = true;

function Callback() {
  const dispatch = useSampleDispatch();
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
        .post("http://localhost:5000/callback", { authorizationCode })
        .then((res) => {
          const accessToken = res.data.accessToken;
          const refreshToken = document.cookie
            .split(" ")
            .filter((cookie) => cookie.includes("hanchart"))[0]
            .split("=")[1];
          console.log(accessToken, refreshToken);
          //여기서 액세스 토큰의 상태를 갱신하여, 전역에 저장한다
          dispatch({ type: SAVE, accessToken });
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
