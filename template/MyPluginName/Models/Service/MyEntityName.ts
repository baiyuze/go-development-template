import { Base } from '@/libs/Base/Base'
const request = Base.request

/**
 * 添加
 * @returns
 */
export const addMyEntityName = (data: any) => {
  return request.post('/api/v1/myPluginName/myEntityName', data)
}

/**
 * 批量删除
 * @returns
 */
export const deleteMyEntityNames = (ids: string[]) => {
  return request({
    data: ids,
    url: '/api/v1/myPluginName/myEntityName',
    method: 'delete',
  })
}

/**
 * 克隆
 * @returns
 */
export const cloneData = (data: any) => {
  return request.post('/api/v1/myPluginName/myEntityName/clone', data)
}
