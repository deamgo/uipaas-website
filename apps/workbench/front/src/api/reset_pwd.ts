import request from "@/utils/axios"
import { IUsrAccount } from "./account"

function forgotVerify(data: IUsrAccount) {
  return request({
    url: '/forgot_verify',
    method: 'post',
    data: data,
    headers: { 'Content-Type': 'application/json' }
  })
}


function ResetPwd(data: IUsrAccount) {
  return request({
    url: '/reset_password',
    method: 'post',
    data: data,
    headers: { 'Content-Type': 'application/json' }
  })
}

export {
  ResetPwd,
  forgotVerify
}