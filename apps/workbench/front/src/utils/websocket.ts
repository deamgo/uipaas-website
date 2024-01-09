import { tokenStore } from "@/store/store"


let token = tokenStore.getToken().length > 0 ? tokenStore.getToken() : 'invalid'
export const socket = new WebSocket('ws://192.168.31.111:8989/v1/ws', [token])

socket.onopen = () => {
  console.log('websocket::open')
}

socket.onclose = () => {
  console.log('websocket::close')
}
