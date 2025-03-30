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
 * 获取详情
 * @returns
 */
export const getMyEntityName = (id: string) => {
  return request.get(`/api/v1/myPluginName/myEntityName/${id}`)
}

/**
 * 更新
 * @returns
 */
export const updateMyEntityName = (id: string, data: Record<string, any>) => {
  return request.put(`/api/v1/myPluginName/myEntityName/${id}`, data)
}
