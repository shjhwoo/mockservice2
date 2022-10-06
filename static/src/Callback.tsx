import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

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
          //이제 서비스 전용으로만 쓸 수 있는 액세스 토큰을 받았음.
          console.log(res.data, "답답해");
          window.localStorage.setItem("userid", res.data.userid);
          navigate("/", { replace: true });
          window.location.reload();
          console.log("성공");
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
