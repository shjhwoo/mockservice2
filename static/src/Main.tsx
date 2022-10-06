import axios from "axios";
import { useState, useEffect } from "react";

function Main() {
  useEffect(() => {
    axios
      .get("http://localhost:4000/checksso")
      .then((res) => {
        //사용자를 바로 SSO쿠키가 있는지 없는지 판단하는 페이지로 보낸다.
        document.cookie = "isPKCE=true";
        window.location.replace(res.data.redirectionURL);
      })
      .catch((err) => {
        console.log(err, "에러 확인");
      });
  }, []);
  return <div>베가스 접속중...</div>;
}

export default Main;
