import axios from "axios";
import React, { useState, useEffect } from "react";

function Service() {
  const [isDoctor, setIsDoctor] = useState<boolean>(false);
  useEffect(() => {
    axios
      .get("https://localhost:4000/api/chart")
      .then((res) => {
        if (res.data.isDoctor) {
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
