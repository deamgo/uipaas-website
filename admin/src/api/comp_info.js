import { request } from '../util/request'

const getToken = () => {
    return localStorage.getItem('jwtToken')
}

function getComponyList(params) {
    return request({
        url: '/company',
        method: 'get',
        params: params,
        Headers: {
            'Content-Type': 'application',
            'Authorization': `Bearer ${getToken()}`
        }
    })
}

const updateToken = (newToken) => {
    localStorage.setItem('jwtToken', newToken)
}

export {
    getComponyList,
    updateToken
}