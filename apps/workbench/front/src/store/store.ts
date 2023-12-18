import { makeAutoObservable, observable, configure, runInAction } from 'mobx'
import { IUsrAccount } from '@api/account'

class AppStore {

  constructor() {
    makeAutoObservable(this)
  }

  userInfo: IUsrAccount | null = null

  setUserInfo(user: typeof this.userInfo) {
    this.userInfo = user
  }
}

const appStore = new AppStore()

export default appStore