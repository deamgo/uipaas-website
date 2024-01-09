import { createApplication, getApplicationList } from '@/api/application'
import { makeAutoObservable } from 'mobx'

interface IAppStats {
  shouldRefresh: boolean,
}

class Application {

  constructor() {
    makeAutoObservable(this)
  }

  AppStats: IAppStats = {
    shouldRefresh: false
  }

  getAppStats() {
    return this.AppStats
  }

  setAppStats() {
    this.AppStats.shouldRefresh = !this.AppStats.shouldRefresh
  }

  async createApp(name: string) {
    try {
      const { value } = await createApplication({
        name
      })
      if (value.code === 0) {
        return value.msg
      } else {
        return value.msg
      }
    } catch (error) {
      console.log(error);
    }
  }

  async getApp() {
    try {
      const { value } = await getApplicationList()
      if (value.code === 0) {
        if (value.data !== null) {
          return value.data
        }
        return []
      } else {
        return []
      }
    } catch (error) {
      console.log(error);
      return []
    }
  }
}

const applicationStore = new Application()

export {
  applicationStore
}