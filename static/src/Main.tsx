import axios from "axios";
import { useState } from "react";

function Main() {
  const [isLogin, setIsLogin] = useState<boolean>(false);
  const handleLogin = () => {
    console.log("확인");
    axios.get("http://localhost:4000/sso/login").then((response) => {
      window.location.replace(response.request.responseURL);
      document.cookie = "isPKCE=true;";
    });
  };
  return (
    <div>
      {isLogin ? (
        "로그인한 사람만 볼수있지롱~"
      ) : (
        <button onClick={handleLogin}>통합로그인 하러가기</button>
      )}
    </div>
  );
}

export default Main;
