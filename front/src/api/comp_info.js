import {request }from '../util/request'

function saveCompInfo(data){
    return request({
        url: '/company',
        method: 'post',
        data:data,
        Headers: {'Content-Type': 'application'}
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