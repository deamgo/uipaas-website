import { tokenStore } from "@/store/store";
import request from "@/utils/axios"
import axios from "axios";
import Cookies from "js-cookie";

export type IUsrWorkspace = {
    name?: string
    logo?: string | null
}

request.defaults.headers.common['Authorization'] = 'Bearer ' + Cookies.get('token')

function workspaceList() {
    return request({
        url: '/workspace/list',
        method: 'get',
        headers: {
            'Content-Type': 'application',
            // 'Authorization': 'Bearer '+ sessionStorage.getItem('token')
        }
    })
}


function workspaceCreate(data: IUsrWorkspace) {
    return request({
        url: '/workspace/create',
        method: 'post',
        data: data,
        headers: {
            'Content-Type': 'application/json',
            // 'Authorization': 'Bearer '+ sessionStorage.getItem('token')
        }
    })
}


function workspaceLogo(formData: FormData) {
    return request({
        url: '/workspace/logo',
        method: 'post',
        data: formData,
        headers: {
            'Content-Type': 'multipart/form-data',
            // 'Authorization': 'Bearer '+ sessionStorage.getItem('token')
        }
    })

}

export {
    workspaceCreate,
    workspaceList,
    workspaceLogo
}