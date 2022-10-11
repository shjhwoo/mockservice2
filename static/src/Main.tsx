import { useEffect, useState } from "react";
import api from "./Api";

function Main() {
  const [isLogin, setIsLogin] = useState<boolean>(false);
  useEffect(() => {
    //전역에 있는 액세스 토큰 가져와서 바로 서버에 검증 요청을 보냄
    const check = async () => {
      const resp = await api.checkServiceToken(); //전역에서 꺼내와~
      if (resp === undefined) return;
      //resp가 바로 인증받은 경우
      //refresh해서 인증받은 경우
      if (
        resp.data.message === "directly authorized" ||
        resp.data.message === "authorized with refreshed access token"
      ) {
        //서비스 접근 권한 존재함. 기본 메인화면 보여준다.
        setIsLogin(true);
      }
      if (resp.data.message === "have to check sso cookie") {
        //refresh마저 안되는 경우: redirectionURL을 받아 SSO확인하러간다.
        setIsLogin(false);
        document.cookie = "isPKCE=true";
        const redirectionURL = resp;
        location.replace(redirectionURL);
      }
    };
    check();
    // axios
    //   .get("http://localhost:4000/checksso")
    //   .then((res) => {
    //     //사용자를 바로 SSO쿠키가 있는지 없는지 판단하는 페이지로 보낸다.
    //     document.cookie = "isPKCE=true";
    //     window.location.replace(res.data.redirectionURL);
    //   })
    //   .catch((err) => {
    //     console.log(err, "에러 확인");
    //   });
  }, []);
  return <div>베가스 접속중...</div>;
}

export default Main;
