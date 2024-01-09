import request from "@/utils/axios"
import { IUsrAccount } from "./account"
import { tokenStore } from "@/store/store"

// manually refresh rather then auto updated
// const token = tokenStore.getToken()
// // request.defaults.headers.common['Authorization'] = 'Bearer ' + Cookies.get('token')
// request.defaults.headers.common['Authorization'] = 'Bearer ' + token
// // 'Authorization': 'Bearer ' + Cookies.get('token')



function getUserInfo() {
  return request({
    url: '/developer',
    method: 'get',
    headers: {
      'Content-Type': 'application',
      'Authorization': 'Bearer ' + tokenStore.getToken()
    }
  })
}

const updateUserName = (id: string, data: IUsrAccount) => {
  return request({
    url: '/developer/username/' + id,
    method: 'put',
    data: data,
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + tokenStore.getToken()
    }
  })
}

const verifiPwdEmail = (data: IUsrAccount) => {
  return request({
    url: '/developer/email/firststep',
    method: 'post',
    data: data,
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + tokenStore.getToken()
    }
  })
}

const verifiEmail = (data: IUsrAccount) => {
  return request({
    url: '/developer/email/secondstep',
    method: 'post',
    data: data,
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + tokenStore.getToken()
    }
  })
}

const verifiEmailCode = (data: IUsrAccount) => {
  return request({
    url: '/developer/email/thirdstep',
    method: 'put',
    data: data,
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + tokenStore.getToken()
    }
  })
}

const sendPwdEmailCode = (data: IUsrAccount) => {
  return request({
    url: '/developer/password/firststep',
    method: 'post',
    data: data,
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + tokenStore.getToken()
    }
  })
}

const verifiPwdEmailCode = (data: IUsrAccount) => {
  return request({
    url: '/developer/password/secondstep',
    method: 'post',
    data: data,
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + tokenStore.getToken()
    }
  })
}

const updatePwd = (data: IUsrAccount) => {
  return request({
    url: '/developer/password/thirdstep',
    method: 'put',
    data: data,
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + tokenStore.getToken()
    }
  })
}

export {
  getUserInfo,
  updateUserName,
  verifiPwdEmail,
  verifiEmail,
  verifiEmailCode,
  sendPwdEmailCode,
  verifiPwdEmailCode,
  updatePwd,
}