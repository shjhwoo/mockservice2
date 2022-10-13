import axios from "axios";

import { useState, useEffect } from "react";

import { Link } from "react-router-dom";
import api from "./Api";

interface token {
  accessToken: string;
  refreshToken: string;
}

interface Props {
  token: token;
}

function Service(props: Props) {
  const [chart, setChart] = useState("");
  const handleLogout = () => {
    console.log("로그아웃을 요청합니다");
    axios
      .post("http://localhost:4000/logout", {
        cookie: document.cookie
          .split(" ")
          .filter((item) => item.includes("vegasRefreshToken"))[0],
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
  console.log("<service/>", props.token);
  const handleChartRequest = () => {
    console.log(props.token.accessToken, "서비스 요청 보내기 직전에 토큰확인.");
    api.getChart(props.token.accessToken).then((res) => {
      console.log("<service/>", res, "차트 요청에 대한 응답입니다.");
      //만약에, 리프레시 토큰마저도 만료되었다면 SSO 세션을 확인하러 가야한다.
      if (typeof res === "string" && res.includes("http://")) {
        alert("서비스 토큰이 만료되었습니다. SSO세션을 확인합니다");
        document.cookie = "isPKCE=true;";
        window.location.replace(res);
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
