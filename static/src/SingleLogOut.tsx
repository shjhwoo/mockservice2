import axios from "axios";
import { useEffect } from "react";

interface token {
  accessToken: string;
  refreshToken: string;
}

interface Props {
  token: token;
}

axios.defaults.withCredentials = true;

function SingleLogOut(props: Props) {
  useEffect(() => {
    axios
      .post("http://localhost:4000/slo", {})
      .then((response) => {
        console.log("서비스 쿠키 파괴 완료.");
        console.log(response);
      })
      .catch((err) => {
        console.log(err);
      });
  });

  return (
    <>
      <div>로그아웃 중...</div>
    </>
  );
}

export default SingleLogOut;
