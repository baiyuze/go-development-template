import { Base } from '@/libs/Base/Base'
import {
  deleteMyEntityNames,
  addMyEntityName,
  cloneData,
} from './Service/MyEntityName'

export class MyEntityName extends Base<{ [key: string]: any }> {
  constructor() {
    super({
      data: [],
    })
  }
  onMounted() {}
  /**
   * 删除
   * @param id
   * @returns
   */
  async deleteMyEntityNames(ids: string[]) {
    return deleteMyEntityNames(ids)
  }

  /**
   * 添加数据
   * @param data
   * @returns
   */
  addMyEntityName(data: Record<string, any>) {
    return addMyEntityName(data)
  }

  /**
   * 克隆
   * @param ids
   * @returns
   */
  cloneData(ids: string[]) {
    return cloneData(ids)
  }
}
