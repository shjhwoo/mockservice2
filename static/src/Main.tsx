import axios from "axios";
import { useState, useEffect } from "react";

function Main() {
  const [isLogin, setIsLogin] = useState<boolean>(false);
  console.log(document.cookie, "쿠키 정보?");
  useEffect(() => {
    axios
      .post("http://localhost:4000/checkcookie", {
        headers: {
          withCredentials: true,
        },
        cookie: document.cookie,
      })
      .then((res) => {
        //쿠키가 있으니까 원래 서비스로 되돌아온다.
        console.log(res, "응답 확인");
        //서비스 쿠키가 살아있는 경우
        if (res.data.message === "has login cookie") {
          setIsLogin(true);
        } else {
          //서비스 쿠키는 없으므로 SSO쿠키가 살아있는지 확인하러 간다.
          console.log("리디렉션 확인 얍!!");
          document.cookie = "isPKCE=true";
          window.location.replace(res.data.redirectionURL);
        }
      })
      .catch((err) => {
        //쿠키가 없으니까 로그인 버튼을 사용자에게 보여준다.
        console.log(err, "에러 확인");
        setIsLogin(false);
      });
  }, []);
  const handleLogin = () => {
    console.log("확인");
    axios.get("http://localhost:4000/sso/login").then((response) => {
      window.location.replace(response.request.responseURL);
      document.cookie = "isPKCE=true;";
    });
  };
  console.log(localStorage.getItem("userid"), "확인");
  return (
    <div>
      {isLogin ? (
        <>
          <div>모든 직원이 볼 수 있는 정보</div>
          <div>의사만 볼 수 있는 정보</div>
        </>
      ) : (
        <button onClick={handleLogin}>통합로그인 하러가기</button>
      )}
    </div>
  );
}

export default Main;
