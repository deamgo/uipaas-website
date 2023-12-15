import { makeAutoObservable } from 'mobx'
import { IUsrAccount } from '@api/account'

class AppStore {
  userInfo: IUsrAccount | null = null;

  constructor() {
    makeAutoObservable(this)
  }

  setUserInfo(userInfo: typeof this.userInfo) {
    this.userInfo = userInfo
  }
}

const appStore = new AppStore()

export default appStore