import axios from "axios";

function Main() {
  const handleLogin = () => {
    console.log("확인");
    axios.get("http://localhost:4000/sso/login").then((response) => {
      window.location.replace(response.request.responseURL);
      document.cookie = "isPKCE=true;";
    });
  };
  console.log(localStorage.getItem("userid") === "", "확인");
  return (
    <div>
      {!!localStorage.getItem("userid") === false ||
      localStorage.getItem("userid") === undefined ? (
        <button onClick={handleLogin}>통합로그인 하러가기</button>
      ) : (
        "로그인한 사람만 볼수있지롱~"
      )}
    </div>
  );
}

export default Main;
