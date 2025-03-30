import { ref, onMounted, reactive, Ref, nextTick, computed } from 'vue'
import { injectModel } from '@/libs/Provider/Provider'
import { MyEntityName } from '../Models/MyEntityName'
import { ElMessage } from 'element-plus'
import { ConfirmBox } from '@/components/ConfirmBox/ConfirmBox'
import { useFile } from './File'

interface CurrentType {
  row: any
  index: number
}
export const useMyEntityName = (props: any, ctx?: any) => {
  const myEntityName = injectModel<MyEntityName>('myEntityName')
  const { exportFile } = useFile()
  /**
   * 头部配置
   */
  const headers = ref({})
  /**
   * 动态列配置
   */
  const myEntityNameColumns = ref<Record<string, any>>([])
  /**
   * 搜索值
   */
  const search = ref('')

  /**
   * 排序
   */
  const sort = ref(0)
  /**
   * 选择项
   */
  const selection = ref([])
  /**
   * 当前选中的行
   */
  const current = ref<any>(null)
  /**
   * 数据源
   */
  const dataSource: Ref<any[]> = ref([])

  /**
   * 表格
   */
  const tableRef = ref()
  const dialogConfig = reactive({
    visible: false,
    title: '',
    isAdd: false,
  })

  const dialogSettingConfig = reactive({
    visible: false,
    title: '',
  })

  /**
   * 分页数据
   */
  const paginationParams = ref({})

  /**
   * 打开详情
   * @param row
   */
  const openDetail = (row: any) => {
    current.value = row
    dialogConfig.visible = true
    dialogConfig.title = row.name
    dialogConfig.isAdd = false
    sort.value = row.sort
  }

  const contextMenu = [
    {
      label: '展开详情',
      fn: (c: CurrentType) => {
        current.value = null
        sort.value = c.row.sort
        nextTick(() => openDetail(c.row))
      },
      divided: true,
      icon: 'o',
    },
    {
      label: '向上添加',
      fn: (c: CurrentType, pageNum: number) => {
        current.value = null
        sort.value = c.index + 1 + (pageNum - 1) * 50
        dialogConfig.visible = true
        dialogConfig.title = '添加'
        dialogConfig.isAdd = false
      },
      divided: true,
      icon: 'up',
    },
    {
      label: '向下添加',
      fn: (c: CurrentType, pageNum: number) => {
        current.value = null
        sort.value = c.index + 2 + (pageNum - 1) * 50
        dialogConfig.visible = true
        dialogConfig.title = '添加'
        dialogConfig.isAdd = false
      },
      divided: true,
      icon: 'down',
    },
    {
      label: '创建副本',
      fn: async ({ row }: CurrentType) => {
        await myEntityName.cloneData([row.id])
        ElMessage.success('创建副本成功')
        tableRef.value?.getList()
      },
      divided: true,
      icon: 'copy',
    },
    {
      label: '删除',
      fn: async (c: CurrentType) => {
        const names = selection.value.map((item: { name: string }) => item.name)
        ConfirmBox(
          `是否删除${names.length ? names.join(',') : c.row.name}`
        ).then(async () => {
          const ids = selection.value.map((item: { id: string }) => item.id)
          await myEntityName.deleteMyEntityNames(ids.length ? ids : [c.row.id])
          ElMessage.success('删除成功')
          tableRef.value.getList()
        })
      },
      icon: 'close',
    },
  ]

  const onCheck = (records: any) => {
    selection.value = records
  }

  const onAddMyEntityName = () => {
    const params = tableRef.value?.getPaginationParams()
    current.value = null
    dialogConfig.visible = true
    dialogConfig.isAdd = true
    dialogConfig.title = '添加'
    sort.value = params.totalCount + 1
  }

  const onConfirmMyEntityName = async () => {
    dialogConfig.visible = false
    if (dialogConfig.isAdd) {
      tableRef.value?.scrollToRow({
        skip: true,
      })
    } else {
      await tableRef.value?.getList()
    }
  }
  /**
   * 行点击时更新current
   */
  const onRowClick = ({ row }: any) => {
    if (dialogConfig.visible && current.value) {
      current.value = row
    }
  }
  /**
   * 导出
   */
  const onExport = () => {
    const params = tableRef.value?.getParams()
    exportFile(
      '/api/v1/myPluginName/myEntityName/export',
      params,
      'myPluginName'
    )
  }

  /**
   * 关键字搜索
   */
  const onSearch = () => {
    tableRef.value?.getList({
      Filter: search.value,
    })
  }

  /**
   * 重置表格数据
   */
  const reloadList = () => {
    tableRef.value?.getList()
  }
  /**
   * 上传成功
   */
  const onSuccess = () => {
    tableRef.value?.getList()
    ElMessage.success('导入成功')
  }
  /**
   * 失败
   * @param err
   */
  const onError = (err: any) => {
    try {
      const message = JSON.parse(err.message)
      ElMessage.error(message.msg)
    } catch (error) {
      ElMessage.error('导入失败')
    }
  }
  /**
   * 上传钩子
   */
  const onBeforeUpload = (file: File) => {
    const format = ['xlsx', 'xls', 'csv']
    if (!format.includes(file.name.split('.')[1])) {
      ElMessage.error('导入文件格式不正确，请导入.xlsx/.xls与.csv格式的文件')
      return false
    }
    return true
  }

  onMounted(() => {
    headers.value = {
      Authorization: `Bearer ${sessionStorage.getItem('Token')}`,
      'X-Project': sessionStorage.getItem('X-Project'),
    }
  })

  ctx.expose({
    reloadList,
  })

  return {
    dataSource,
    contextMenu,
    dialogConfig,
    dialogSettingConfig,
    tableRef,
    current,
    search,
    sort,
    myEntityNameColumns,
    paginationParams,
    headers,
    onBeforeUpload,
    onError,
    onSuccess,
    openDetail,
    onSearch,
    onExport,
    onRowClick,
    onConfirmMyEntityName,
    onCheck,
    onAddMyEntityName,
  }
}
