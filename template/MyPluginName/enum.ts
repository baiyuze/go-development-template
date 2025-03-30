export interface TabType {
  label: string
  name: string
  columns?: any[]
  data?: any[]
  isFooter: boolean
  [key: string]: any
}

export const permissionCodes = {
  'myEntityName-list': '列表-列表',
  'myEntityName-add': '列表-添加',
  'myEntityName-import': '列表-导入',
  'myEntityName-output': '列表-输出',
}
