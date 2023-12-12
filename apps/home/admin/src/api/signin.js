import { request } from '../util/request'

function login(data) {
  return request({
    url: '/login',
    method: 'post',
    data: data,
    Headers: { 'Content-Type': 'application' }
  })
}

export {
  login,
}