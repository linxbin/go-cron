import request from '@/utils/request'

const api = {
  list: '/task/list',
  detail: '/task/detail',
  enable: '/task/enable',
  disable: '/task/disable',
  add: '/task/create',
  edit: '/task/update',
  delete: '/task/delete',
  logList: '/task-log/list',
  clearLogList: '/task-log/clear'
}

export function getTaskList (parameter) {
  return request({
    url: api.list,
    method: 'get',
    params: parameter
  })
}

export function getTaskDetail (id) {
  return request({
    url: api.detail + '/' + id,
    method: 'get'
  })
}

export function addTask (parameter) {
  return request({
    url: api.add,
    method: 'post',
    data: parameter
  })
}

export function editTask (parameter) {
  return request({
    url: api.edit,
    method: 'post',
    data: parameter
  })
}

export function enableTask (parameter) {
  return request({
    url: api.enable,
    method: 'post',
    data: parameter
  })
}

export function deleteTask (parameter) {
  return request({
    url: api.delete,
    method: 'post',
    data: parameter
  })
}

export function disableTask (parameter) {
  return request({
    url: api.disable,
    method: 'post',
    data: parameter
  })
}

export function taskLogList (id, parameter) {
  return request({
    url: api.logList + '/' + id,
    method: 'get',
    params: parameter
  })
}

export function clearLogList (id) {
  return request({
    url: api.clearLogList,
    method: 'post',
    params: { task_id: id }
  })
}
