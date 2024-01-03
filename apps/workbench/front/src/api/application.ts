import request from "@/utils/axios"
import { tokenStore } from "@/store/store"

function createApplication() {
  return request({
    url: '/developer',
    method: 'get',
    headers: {
      'Content-Type': 'application',
      'Authorization': 'Bearer ' + tokenStore.getToken()
    }
  })
}

export {
  createApplication,
}