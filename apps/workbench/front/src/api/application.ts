import request from "@/utils/axios"
import { tokenStore } from "@/store/store"
import { currentWorkspaceStore } from "@/store/wsStore"

function createApplication(data: { name: string }) {
  return request({
    url: '/workspace/' + currentWorkspaceStore.getCurrentWorkspace().id + '/application',
    method: 'post',
    data,
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + tokenStore.getToken()
    }
  })
}

function getApplicationList() {
  return request({
    url: '/workspace/' + currentWorkspaceStore.getCurrentWorkspace().id + '/application',
    method: 'get',
    headers: {
      'Content-Type': 'application',
      'Authorization': 'Bearer ' + tokenStore.getToken()
    }
  })
}

export {
  createApplication,
  getApplicationList
}