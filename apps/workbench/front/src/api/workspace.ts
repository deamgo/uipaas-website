import request from "@/utils/axios"

export type IUsrWorkspace = {
    name?: string
    logo?: string | null
}

function workspaceList(){
    return request({
        url: '/workspace/list',
        method: 'post',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer '+ sessionStorage.getItem('token')
        }
    })
}

function workspaceCreate(data:IUsrWorkspace){
    return request({
        url: '/workspace/create',
        method: 'post',
        data: data,
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer '+ sessionStorage.getItem('token')
        }
    })
}

export {
    workspaceCreate,
    workspaceList
}