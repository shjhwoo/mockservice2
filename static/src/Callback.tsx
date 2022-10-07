import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import { slientRefresh } from "./Timer";

function Callback() {
  const navigate = useNavigate();
  useEffect(() => {
    const url = new URL(window.location.href);
    const authorizationCode = url.searchParams.get("code");
    console.log(authorizationCode);
    if (authorizationCode) {
      axios
        .post(
          "http://localhost:4000/callback",
          { authorizationCode },
          { withCredentials: true }
        )
        .then((res) => {
          console.log(
            "서비스에서 사용할 수 있는 액세스 토큰이 발급되었습니다."
          );
          window.localStorage.setItem("userid", res.data.userid);
          //쿠키 유효시간 타이머 실행시키도록 할 것.
          //slientRefresh();
          navigate("/service", { replace: true });
        })
        .catch((err) => {
          console.log("catch", err);
        });
    }
  }, [navigate]);
  return (
    <>
      <div>서비스 리디렉션 중입니다...</div>
    </>
  );
}

export default Callback;
