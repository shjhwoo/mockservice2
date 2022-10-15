import { Link } from "react-router-dom";

function Nurse() {
  return (
    <>
      <div>간호사 페이지</div>
      <Link to="/service">서비스 페이지로 이동</Link>
    </>
  );
}

export default Nurse;
