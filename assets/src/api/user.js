import request from '@/utils/request'

const api = {
  list: '/user/list',
  add: '/user/add',
  delete: '/user/delete'
}

export function getUserList (parameter) {
  return request({
    url: api.list,
    method: 'get',
    params: parameter
  })
}

export function addUser (parameter) {
  return request({
    url: api.add,
    method: 'post',
    data: parameter
  })
}

export function deleteUser (parameter) {
  return request({
    url: api.delete,
    method: 'post',
    data: parameter
  })
}
