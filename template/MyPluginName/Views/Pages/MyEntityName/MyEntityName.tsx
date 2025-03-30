import { defineComponent } from 'vue'
import type { Ref } from 'vue'
import BaseTable from '@/components/Table/Table'
import styles from './MyEntityName.module.scss'
import { useMyEntityName } from '../../../Controllers/MyEntityName'
import IconButton from '@/components/IconButton/IconButton'
import MyEntityNameDrawer from '../Dialog/MyEntityNameDrawer/MyEntityNameDrawer'
import Search from '@/components/Search/Search'
import { columns } from './Config'
import TdButton from '@/components/TdButton/TdButton'
import { vPermission } from '@/libs/Permission/Permission'

interface RenderTableType {
  url?: string
  dataSource: Ref<any[]>
  isDrag?: boolean
  isChecked?: boolean
  isHidePagination?: boolean
  params?: Record<string, any>
  autoHeight?: boolean
}

export default defineComponent({
  name: 'MyEntityName',
  directives: {
    permission: vPermission,
  },
  setup(props, ctx) {
    const {
      dataSource,
      contextMenu,
      dialogConfig,
      tableRef,
      current,
      search,
      sort,
      headers,
      onError,
      onSearch,
      onRowClick,
      onConfirmMyEntityName,
      onCheck,
      onAddMyEntityName,
      onExport,
      openDetail,
      onSuccess,
      onBeforeUpload,
    } = useMyEntityName(props, ctx)

    /**
     * @returns 表格
     */
    const RenderBaseTable = (props: RenderTableType) => {
      const {
        url,
        dataSource,
        isDrag,
        isChecked,
        isHidePagination,
        params,
        autoHeight,
      } = props

      return (
        <div
          class={{
            [styles.myEntityNameList]: true,
          }}
        >
          <BaseTable
            ref={tableRef}
            url={url}
            sortUrlTpl="/api/v1/myPluginName/myEntityName/{id}/adjustsort/{sort}"
            v-model:dataSource={dataSource.value}
            columns={columns}
            contextMenu={contextMenu}
            params={params}
            isDrag={isDrag}
            isChecked={isChecked}
            autoHeight={autoHeight}
            onCheck={onCheck}
            onRowClick={onRowClick}
            isHidePagination={isHidePagination}
            pageSize={50}
            v-slots={{
              name: ({ row }: any) => {
                return row?.name ? (
                  <TdButton
                    onClick={() => openDetail(row)}
                    text={<span style="color:#5a84ff">详情</span>}
                    icon="scale"
                    tip={row?.name}
                    hover
                  >
                    {row?.name}
                  </TdButton>
                ) : (
                  '-'
                )
              },
            }}
          ></BaseTable>
        </div>
      )
    }
    return () => {
      return (
        <div class={styles.myEntityNameContent}>
          {/* 添加/编辑 */}
          <MyEntityNameDrawer
            v-model={dialogConfig.visible}
            title={dialogConfig.title}
            row={current.value}
            sort={sort.value}
            onConfirm={onConfirmMyEntityName}
          />
          <div class={styles.headerContent}>
            <div class={styles.header}>
              <IconButton
                v-permission="myEntityName-add"
                icon="add-p"
                onClick={onAddMyEntityName}
                type="primary"
              >
                添加
              </IconButton>
              <el-divider direction="vertical" />
              <el-upload
                v-permission="myEntityName-import"
                name="file"
                accept=".xlsx,.xls,.csv"
                show-file-list={false}
                onError={onError}
                onSuccess={onSuccess}
                before-upload={onBeforeUpload}
                headers={headers.value}
                action="/api/v1/myPluginName/myEntityName/import"
              >
                <IconButton icon="in">导入</IconButton>
              </el-upload>

              <IconButton
                v-permission="myEntityName-output"
                icon="out"
                onClick={onExport}
              >
                导出
              </IconButton>
            </div>
            <Search
              placeholder="请输入关键字"
              v-model={search.value}
              onConfirm={onSearch}
              style={{ marginTop: '-1px' }}
            />
          </div>
          <RenderBaseTable
            url="/api/v1/myPluginName/myEntityName"
            dataSource={dataSource}
            isChecked={true}
            isDrag={true}
          />
        </div>
      )
    }
  },
})
