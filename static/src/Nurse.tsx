import { Link } from "react-router-dom";

interface token {
  accessToken: string;
  refreshToken: string;
}

interface Props {
  token: token;
}

function Nurse(props: Props) {
  console.log(props.token, "<Nurse/>");
  return (
    <>
      <div>간호사 페이지</div>
      <Link to="/service">서비스 페이지로 이동</Link>
    </>
  );
}

export default Nurse;
