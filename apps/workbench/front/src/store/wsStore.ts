import { makeAutoObservable } from 'mobx'
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


  setFirst(workspaceName: string) {
    if (workspaceName.length != 0) {
      let index = this.WsList.findIndex(element => element.name === workspaceName);
      let temp = this.WsList[index];
      this.WsList[index] = this.WsList[0];
      this.WsList[0] = temp;
      currentWorkspaceStore.setCurrentWorkspace(temp)
    }
  }


  getWsListFirstByWorkspace(workspaceName: string) {
    this.setFirst(workspaceName)
    return this.WsList
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
    window.sessionStorage.setItem('currentWorkspace', JSON.stringify(item))
    this.currentWorkspace = item
  }

  getCurrentWorkspace() {
    return this.currentWorkspace
  }

  resetCurrentWorkspace() {
    this.currentWorkspace = {} as IWorkspaceItemProps
  }

  loadCurrentWorkspace() {
    let currentWorkspace = window.sessionStorage.getItem('currentWorkspace')
    this.setCurrentWorkspace(JSON.parse(currentWorkspace as string))
  }

}

const wsStore = new WsStore()
const currentWorkspaceStore = new CurrentWorkspaceStore()

export {
  wsStore,
  currentWorkspaceStore
}