import axios from "axios"
import { useEffect } from "react"

function SingleLogOut() {
    useEffect(() => {
        axios
        .post(
        "http://localhost:4000/slo",             
        { withCredentials: true },
        )
        .then((response)=> {
            console.log("서비스 쿠키 파괴 완료.")
            console.log(response)
        })
        .catch((err) => {
            console.log(err)
        })
    })

    return (
        <>
            <div>로그아웃 중...</div>
        </>
    )
}


export default SingleLogOut