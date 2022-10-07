//import React from "react";

import axios from "axios";

//서비스 전역적으로 사용할 쿠키 타이머 함수입니다.
export const slientRefresh = () => {
  console.log("쿠키 타이머가 작동을 시작하였습니다");
  console.log(`현재 페이지의 위치는${window.location.href}입니다.`);
  setTimeout(() => {
    axios
      .get("http://localhost:4000/checksso")
      .then((response) => {
        console.log(
          "서비스 쿠키 만료 1분전입니다. 자동으로 SSO쿠키 여부를 확인하러 갑니다..."
        );
        window.location.replace(response.data.redirectionURL);
      })
      .catch((err) => {
        console.log(err, "SSO 쿠키 확인 url 신청에 문제가 발생했습니다.");
      });
  }, 1000 * 60);
};
