import request from "@/utils/axios"
import { IUsrAccount } from "./account"
import { tokenStore } from "@/store/store"
import Cookies from "js-cookie"

request.defaults.headers.common['Authorization'] = 'Bearer ' + tokenStore.getToken()

// 'Authorization': 'Bearer ' + Cookies.get('token')


function getUserInfo() {
  return request({
    url: '/developer',
    method: 'get',
    headers: {
      'Content-Type': 'application'
    }
  })
}

const updateUserName = (id: string, data: IUsrAccount) => {
  return request({
    url: '/developer/username/' + id,
    method: 'put',
    data: data,
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

const verifiPwdEmail = (data: IUsrAccount) => {
  return request({
    url: '/developer/email/firststep',
    method: 'post',
    data: data,
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

const verifiEmail = (data: IUsrAccount) => {
  return request({
    url: '/developer/email/secondstep',
    method: 'post',
    data: data,
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

const verifiEmailCode = (data: IUsrAccount) => {
  return request({
    url: '/developer/email/thirdstep',
    method: 'put',
    data: data,
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

const sendPwdEmailCode = (data: IUsrAccount) => {
  return request({
    url: '/developer/password/firststep',
    method: 'post',
    data: data,
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

const verifiPwdEmailCode = (data: IUsrAccount) => {
  return request({
    url: '/developer/password/secondstep',
    method: 'post',
    data: data,
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

const updatePwd = (data: IUsrAccount) => {
  return request({
    url: '/developer/password/thirdstep',
    method: 'put',
    data: data,
    headers: {
      'Content-Type': 'application/json'
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