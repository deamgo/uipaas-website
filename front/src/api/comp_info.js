import { request } from '../util/request'

function saveCompInfo(data) {
    return request({
        url: '/company',
        method: 'post',
        data: data,
        headers: { 'Content-Type': 'application/json' }
    })
}

// function test() {
//     return request({
//         url: '/',
//         method: 'get',
//         Headers: {'Content-Type': 'application'}
//     })
// }

export {
    saveCompInfo,
}