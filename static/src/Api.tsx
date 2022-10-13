import axios from "axios";
type header = { [k: string]: string | boolean };

interface option {
  method: string;
  url: string;
  data?: any;
  headers?: header;
}

interface token {
  accessToken: string;
  refreshToken: string;
}

class Api {
  constructor() {}
  async requestWithCookies(option: option) {
    const { accessToken } = option.data;
    const requestURL = option.url;
    try {
      console.log("여기로옴");
      //액세스 토큰만으로 시도한다.
      option.data = { accessToken };
      const response = await axios(option);
      //서비스 토큰이 정상인 경우
      console.log("액세스 토큰 살아있네 바로 쓰자");
      return response;
    } catch (e: any) {
      console.error(e.response, "에러메세지 확인");
      if (e.response.status === 401) {
        console.log("액세스 토큰 권한 없음", e.response);
        const refreshToken =
          document.cookie
            .split(" ")
            .filter((cookie) => cookie.includes("vegas"))[0] === undefined
            ? ""
            : document.cookie
                .split(" ")
                .filter((cookie) => cookie.includes("vegas"))[0]
                .split("=")[1]
                .replace(/;| /g, "");
        option.url = "http://localhost:4000/refresh";
        option.data = { accessToken, refreshToken };
        try {
          const refreshTokenResponse = await axios(option);
          console.log("리프레시 토큰 응답", refreshTokenResponse);
          if (refreshTokenResponse.status === 201) {
            //새로 발급받은액세스 토큰으로 요청 한번 더 보낸다.
            console.log("새로 받은 액세스토큰으로 재요청");
            option.url = requestURL;
            option.data.accessToken = refreshTokenResponse.data.accessToken;
            option.headers = { withCredentials: true };
            const response = await axios(option);
            console.log(response, "재요청에 대한 응답");
            return response;
          }
        } catch (e: any) {
          //리프레시 토큰마저도 무쓸모.. SSO가 있는지 확인하러 가야함
          console.log(e, "SSO확인하러 갑니다");
          option.method = "GET";
          option.url = "http://localhost:4000/checksso";
          option.data = null;
          const SSOresponse = await axios(option);

          const ssoCheckRedirectionURL = SSOresponse.data.redirectionURL;
          return ssoCheckRedirectionURL;
        }
      }
    }
  }

  //서비스 토큰의 유효성을 확인하는 요청
  async checkServiceToken(token: token) {
    try {
      const { accessToken, refreshToken } = token;
      const option = {
        method: "POST",
        url: "http://localhost:4000/checkservicetkn",
        data: { accessToken, refreshToken },
        headers: { withCredentials: true }, //액세스 토큰과 리프레시 토큰 모두 한꺼번에 보내서 검증한다.
      };
      const response = await axios(option);
      return response;
    } catch (e) {
      console.error(e);
    }
  }

  //예시 서비스 요청: 차트 불러오기
  async getChart(accessToken: string) {
    try {
      const option = {
        method: "POST",
        url: "http://localhost:4000/api/chart",
        data: { accessToken },
        headers: { withCredentials: true },
      };
      const resp = await this.requestWithCookies(option);
      return resp;
    } catch (e) {
      console.error(e);
    }
  }
}

const api = new Api();

export default api;
