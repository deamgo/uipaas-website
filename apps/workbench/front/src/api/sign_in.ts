import request from "@/utils/axios"
import { IUsrAccount } from "./account"


function usrSignIn(data: IUsrAccount) {
  return request({
    url: '/signin',
    method: 'post',
    data: data,
    headers: { 'Content-Type': 'application/json' }
  })
}

export {
  usrSignIn,
}