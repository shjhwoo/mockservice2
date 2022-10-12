import axios from "axios";
import React, { useState, useEffect } from "react";

function Service() {
  const [isDoctor, setIsDoctor] = useState<boolean>(false);
  const handleLogout = () => {
    console.log("로그아웃을 요청합니다");
    axios
      .post("http://localhost:4000/logout", {
        cookie: document.cookie,
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
  useEffect(() => {
    axios
      .post("http://localhost:4000/api/chart", {
        headers: {
          withCredentials: true,
        },
        cookie: document.cookie,
      })
      .then((res) => {
        if (res.data.redirectionURL) {
          //액세스 토큰이 유효하지 않은 경우 SSO 토큰 유무를 확인하는 페이지로 이동함:: relaying state 어떻게 기억하지?
          window.location.replace(res.data.redirectionURL);
        } else if (res.data.isDoctor) {
          setIsDoctor(true);
        } else {
          setIsDoctor(false);
        }
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);
  return (
    <>
      {localStorage.getItem("userid") ? (
        <div>
          <h1>로그인한 사용자만 볼 수 있는 페이지</h1>
          {isDoctor ? (
            <div>의사만 볼 수 있음</div>
          ) : (
            <div>로그인한 누구나 볼수있음</div>
          )}
          <button onClick={handleLogout}>로그아웃</button>
        </div>
      ) : (
        <div>로그인이 필요한 페이지입니다.</div>
      )}
    </>
  );
}

export default Service;

//만약 페이지 자체의 접근 권한을 완전 엄격하게 막고 싶다면 useEffect 써서
//서비스 쿠키를 보내고,
//일단 위의 부분은 이따가 생각해ㅠㅠ

//
