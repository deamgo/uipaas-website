import { makeAutoObservable, observable } from 'mobx'
import { IUserInfo } from '@api/account'
import Cookies from 'js-cookie'
import { getUserInfo } from '@/api/developer_profile'

class AppStore {

  constructor() {
    makeAutoObservable(this)
    this.loadUserInfo()
  }

  userInfo: IUserInfo = {
    id: '',
    username: '',
    email: '',
    avatar: '',
  }

  setUserInfo(user: typeof this.userInfo) {
    sessionStorage.setItem('userInfo', JSON.stringify(user))
    this.userInfo = user
  }

  getUserInfo() {
    return this.userInfo
  }

  loadUserInfo() {
    let userInfo = sessionStorage.getItem('userInfo')
    if (userInfo) {
      Object.assign(this, JSON.parse(userInfo))
    }
  }

  async updateUserInfo() {
    const { value } = await getUserInfo()
    sessionStorage.setItem('userInfo', JSON.stringify(value.data))
    this.setUserInfo(value.data)
  }

  resetUserInfo() {
    this.userInfo = {
      id: '',
      username: '',
      email: '',
      avatar: '',
    }
    console.log(this.userInfo);
  }
}

class TokenStore {

  constructor() {
    makeAutoObservable(this)
    this.loadToken()
  }

  @observable token: string = ''

  setToken(token: typeof this.token) {
    Cookies.set('token', token)
    this.token = token
  }

  getToken() {
    let token = Cookies.get('token')
    if (token) {
      this.setToken(token)
    }
    return this.token
  }

  loadToken() {
    let token = Cookies.get('token')
    if (token) {
      Object.assign(this, token)
    }
  }

  resetToken() {
    this.token = ''
  }
}

const appStore = new AppStore()
const tokenStore = new TokenStore()

export {
  appStore,
  tokenStore
}