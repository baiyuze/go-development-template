import { Base } from '@/libs/Base/Base'
import {
  addMyEntityName,
  getMyEntityName,
  updateMyEntityName,
} from './Service/MyEntityNameDrawer'
import { useGlobalState } from '@/libs/Store/Store'

export class MyEntityNameDrawer extends Base<{ [key: string]: any }> {
  constructor() {
    super({
      data: [],
      myEntityName: {},
    })
  }

  /**
   * 添加
   * @param data
   */
  async addMyEntityName(data: Record<string, any>) {
    return addMyEntityName(data)
  }
  /**
   * 更新
   * @param data
   */
  async updateMyEntityName(id: string, data: Record<string, any>) {
    return updateMyEntityName(id, data)
  }

  /**
   * 获取详情
   */
  async getMyEntityNameDetail(current: any, id?: string) {
    return getMyEntityName(id || current?.id)
  }
}
