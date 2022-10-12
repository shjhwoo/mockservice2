import { useEffect, Dispatch, SetStateAction } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

interface Props {
  accessToken: string;
  refreshToken: string;
  setAccessToken: Dispatch<SetStateAction<string>>;
  setRefreshToken: Dispatch<SetStateAction<string>>;
}

function Callback(props: Props) {
  const navigate = useNavigate();
  useEffect(() => {
    const url = new URL(window.location.href);
    const authorizationCode = url.searchParams.get("code");
    console.log(authorizationCode, props);
    if (authorizationCode) {
      axios
        .post(
          "http://localhost:4000/callback",
          { authorizationCode },
          { withCredentials: true }
        )
        .then((res) => {
          console.log(
            "서비스에서 사용할 수 있는 토큰이 발급되었습니다.",
            res.data
          );
          props.setAccessToken(res.data.accessToken);
          props.setRefreshToken(
            document.cookie
              .split(" ")
              .filter((item) => item.includes("vegasRefreshToken"))[0]
              .split("=")[1]
          );
          console.log(authorizationCode, props, "22");
          window.localStorage.setItem("userid", res.data.userid);
          //navigate("/service", { replace: true });
        })
        .catch((err) => {
          console.log("catch", err);
        });
    }
  }, [props.accessToken, props.refreshToken]);
  return (
    <>
      <div>서비스 리디렉션 중입니다...</div>
    </>
  );
}

export default Callback;
