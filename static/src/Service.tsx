import axios from "axios";
import { useState } from "react";
import {
  REFRESH,
  useSampleState,
  useSampleDispatch,
} from "./contexts/MainContext";
import { Link } from "react-router-dom";
import api from "./Api";

function Service() {
  const state = useSampleState();
  const dispatch = useSampleDispatch();
  const [chart, setChart] = useState("");
  const handleLogout = () => {
    console.log("로그아웃을 요청합니다");
    axios
      .post("http://localhost:5000/logout", {
        cookie: document.cookie
          .split(" ")
          .filter((item) => item.includes("hanchartRefreshToken"))[0],
      })
      .then((response) => {
        console.log(response, "로그아웃 성공 시 돌아오는 응답입니다");
        document.cookie = "isPKCE=true";
        console.log(
          "SSO 쿠키를 파괴하러 갑니다...",
          response.data.redirectionURL
        );
        window.location.assign(response.data.redirectionURL);
      });
  };
  const handleChartRequest = () => {
    api.getChart(state.accessToken).then((res) => {
      console.log("<service/>", res, "차트 요청에 대한 응답입니다.");
      //만약에 응답이 오긴 했는데 액세스 토큰과 같이 돌아왔으면,토큰 갱신해서 저장!
      if (res.accessToken) {
        console.log("액세스 토큰 리프레시된..");
        dispatch({ type: REFRESH, accessToken: res.accessToken });
        setChart(res.response.data.message);
        return;
      }
      //만약에, 리프레시 토큰마저도 만료되었다면 SSO 세션을 확인하러 가야한다.
      if (typeof res === "string" && res.includes("http://")) {
        alert("서비스 토큰이 만료되었습니다. SSO세션을 확인합니다");
        document.cookie = "isPKCE=true;";
        window.location.replace(res);
        return;
      } else {
        setChart(res.data.message);
      }
    });
  };
  return (
    <>
      {localStorage.getItem("userid") ? (
        <div>
          <h1>로그인한 사용자만 볼 수 있는 페이지</h1>
          <section>{chart}</section>
          <div>{`액세스 토큰 보기::${state.accessToken}`}</div>
          <Link to="/nurse">간호사 페이지로 이동</Link>
          <button onClick={handleChartRequest}>차트정보 받아오기</button>
          <button onClick={handleLogout}>로그아웃</button>
        </div>
      ) : (
        <div>로그인이 필요한 페이지입니다.</div>
      )}
    </>
  );
}

export default Service;
