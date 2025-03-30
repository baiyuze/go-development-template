import { importFileToService, exportFileToClient } from '@/api/file'
import { ElMessage } from 'element-plus'
import { downloadFile } from '@/utils'
import dayjs from 'dayjs'

export const useFile = () => {
  /**
   * 导入文件
   * @param url
   */
  const importFile = async (url: string, file: File) => {
    const formData = new FormData()
    formData.append('file', file)
    await importFileToService(url, formData)
    ElMessage('导入成功')
  }
  /**
   * 导出文件
   * @param url
   */
  const exportFile = async (url: string, params: any, name: string) => {
    const res = await exportFileToClient(url, params)
    downloadFile(res, `${name}_${dayjs().format('YYYYMMDDHHMMss')}.xlsx`)
    ElMessage.success('导出成功')
  }

  return {
    importFile,
    exportFile,
  }
}
