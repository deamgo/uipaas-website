import request from "@/utils/axios"
import { tokenStore } from "@/store/store"
import Cookies from "js-cookie"

// request.defaults.headers.common['Authorization'] = 'Bearer ' + tokenStore.getToken()
// // request.defaults.headers.common['Authorization'] = 'Bearer ' + Cookies.get('token')

const deleteWorkspace = (id: number | string) => {
  return request({
    url: '/workspace/' + id,
    method: 'delete',
    headers: {
      'Content-Type': 'application',
      'Authorization': 'Bearer ' + tokenStore.getToken()
    }
  })
}

export {
  deleteWorkspace
}