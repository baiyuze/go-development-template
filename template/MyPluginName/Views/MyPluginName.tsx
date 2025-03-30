import {
  Component,
  DefineComponent,
  defineComponent,
  markRaw,
  ref,
  SetupContext,
  onMounted,
} from 'vue'
import styles from './MyPluginName.module.scss'
// import MyEntityName from './Pages/MyEntityName/MyEntityName'
import Tab from '@/components/Tab/Tab'
import { useProvideModels } from '@/libs/Provider/app'
import { usePermission } from '@/libs/Permission/Permission'
import { permissionCodes } from '../enum'
import { ModuleType, TabItem } from '../type/Type'
import { getEntityNames } from '@/hooks/hook'
import TabPane from '@/components/Tab/TabPane'

const Models: ModuleType = import.meta.glob('./config/*.json', {
  eager: true,
})

const entityNames = getEntityNames(Models)

export default defineComponent({
  name: 'MyPluginName',

  setup(props, ctx: SetupContext) {
    useProvideModels()
    usePermission(props, permissionCodes)

    const rf = ref<{
      [key: string]: any
    }>({})

    const tabData = ref<TabItem[]>([])

    const onTabChange = (v: string) => {
      rf.value?.[v]?.reloadList()
    }

    const initTableData = async () => {
      for (const i in entityNames) {
        const name = entityNames[i]
        const module = await import(`./Pages/${name}/${name}.tsx`)
        const MyEntityName = markRaw(module.default)
        tabData.value.push({
          label: name,
          name,
          component: MyEntityName,
        })
      }
    }

    initTableData()

    return () => {
      return (
        <div class={styles.MyPluginName}>
          <Tab data={tabData.value} type="list" onTab={onTabChange}>
            {tabData.value.map((widgetInfo) => {
              const Widget: any = widgetInfo.component
              return (
                <TabPane label={widgetInfo.label} name={widgetInfo.name}>
                  <Widget
                    ref={(r: any) => (rf.value['MyEntityName'] = r)}
                  ></Widget>
                </TabPane>
              )
            })}
          </Tab>
        </div>
      )
    }
  },
})
