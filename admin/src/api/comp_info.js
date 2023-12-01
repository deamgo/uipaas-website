import {request }from '../util/request'

function getComponyList(params) {
    return request({
        url: '/company',
        method: 'get',
        params: params,
        Headers: {'Content-Type': 'application'}
    })
}

export {
    getComponyList,
}