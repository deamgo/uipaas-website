import {request }from '../util/request'

function saveCompInfo(data){
    return request({
        url: '/comp_info',
        method: 'post',
        data:data,
        Headers: {'Content-Type': 'application'}
    })
}

export   {saveCompInfo}