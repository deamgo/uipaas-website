import { makeAutoObservable, observable, configure, runInAction } from 'mobx'
import Cookies from 'js-cookie'
import { IWorkspaceItemProps } from '@/interface/some'

class WsStore {

  constructor() {
    makeAutoObservable(this)
    // this.loadWsList()
  }

  WsList: IWorkspaceItemProps[] = []

  setWsList(wsList: typeof this.WsList) {
    this.WsList = wsList
  }

  getWsList() {
    return this.WsList
  }

  loadWsList() {
    // let WsList = sessionStorage.getItem('WsList')
    // if (ws) {
    //   Object.assign(this, JSON.parse(WsList))
    // }
  }

  resetWsList() {
    this.WsList = []
    console.log(this.WsList);

  }
}

class CurrentWorkspaceStore {
  constructor() {
    makeAutoObservable(this)
  }

  currentWorkspace: IWorkspaceItemProps = {
    id: '',
    name: '',
    logo: ''
  }

  setCurrentWorkspace(item: typeof this.currentWorkspace) {
    this.currentWorkspace = item
  }

  getCurrentWorkspace() {
    return this.currentWorkspace
  }

  resetCurrentWorkspace() {
    this.currentWorkspace = {
      id: '',
      name: '',
      logo: ''
    }
  }
}

const wsStore = new WsStore()
const currentWorkspaceStore = new CurrentWorkspaceStore()

export {
  wsStore,
  currentWorkspaceStore
}