import request from "@/utils/axios"
import { tokenStore } from "@/store/store"
import Cookies from "js-cookie"
import { currentWorkspaceStore } from "@/store/wsStore"

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

const getDevelopers = (currentPage: number) => {
  return request({
    url: '/workspace/' + currentWorkspaceStore.getCurrentWorkspace().id + '/developer?pageNum=' + currentPage,
    method: 'get',
    headers: {
      'Content-Type': 'application',
      'Authorization': 'Bearer ' + tokenStore.getToken()
    }
  })
}

const findDeveloper = (currentPage: number, queryParam: string) => {
  return request({
    url: '/workspace/' + currentWorkspaceStore.getCurrentWorkspace().id + '/developer?q=' + queryParam + '&pageNum=' + currentPage,
    method: 'get',
    headers: {
      'Content-Type': 'application',
      'Authorization': 'Bearer ' + tokenStore.getToken()
    }
  })
}
interface IDeveloper {
  developer_id?: string
  role?: string
  email?: string
}

const removeDeveloper = (data: IDeveloper) => {
  return request({
    url: '/workspace/' + currentWorkspaceStore.getCurrentWorkspace().id + '/developer',
    method: 'delete',
    data,
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + tokenStore.getToken()
    }
  })
}

const editDeveloperPermission = (data: IDeveloper) => {
  return request({
    url: '/workspace/' + currentWorkspaceStore.getCurrentWorkspace().id + '/developer',
    method: 'put',
    data,
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + tokenStore.getToken()
    }
  })
}

const inviteByEmail = (data: IDeveloper) => {
  return request({
    url: '/workspace/' + currentWorkspaceStore.getCurrentWorkspace().id + '/developer/invite',
    method: 'post',
    data,
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + tokenStore.getToken()
    }
  })
}





export {
  deleteWorkspace,
  getDevelopers,
  findDeveloper,
  removeDeveloper,
  editDeveloperPermission,
  inviteByEmail
}