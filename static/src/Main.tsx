import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useSampleState } from "./contexts/MainContext";
import api from "./Api";

function Main() {
  const state = useSampleState();
  const navigate = useNavigate();
  useEffect(() => {
    //전역에 있는 액세스 토큰 가져와서 바로 서버에 검증 요청을 보냄
    const check = async () => {
      try {
        const resp = await api.checkServiceToken({
          accessToken: state.accessToken,
          refreshToken:
            document.cookie
              .split(" ")
              .filter((cookie) => cookie.includes("vegas"))[0] === undefined
              ? ""
              : document.cookie
                  .split(" ")
                  .filter((cookie) => cookie.includes("vegas"))[0]
                  .split("=")[1]
                  .replace(/;| /g, ""),
        });
        console.log("<Main/>", resp);
        if (resp === undefined) return;
        if (resp.data.message === "SSO 쿠키를 확인합니다") {
          //SSO 세션이 있는지 바로 확인하러 간다. 서비스 토큰이 하나도 없다
          document.cookie = "isPKCE=true;";
          window.location.replace(resp.data.redirectionURL);
          return;
        }
        if (
          resp.data.message === "유효한 액세스 토큰입니다" ||
          resp.data.message === "액세스 토큰이 만료되어 새로 발급했습니다"
        ) {
          console.log("*", resp.data.message);
          navigate("/service", { replace: true });
        }
      } catch (e) {
        console.error(e);
      }
    };
    check();
  }, []);
  return <div>접속중...</div>;
}

export default Main;
