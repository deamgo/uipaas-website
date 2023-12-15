import request from "@/utils/axios"
import { IUsrAccount } from "./account"


function usrSignUp(data: IUsrAccount) {
  return request({
    url: '/signup',
    method: 'post',
    data: data,
    headers: { 'Content-Type': 'application/json' }
  })
}

function usrSignUpVerify(data: IUsrAccount) {
  return request({
    url: '/signup_verify',
    method: 'post',
    data: data,
    headers: { 'Content-Type': 'application/json' }
  })
}

export {
  usrSignUp,
  usrSignUpVerify
}