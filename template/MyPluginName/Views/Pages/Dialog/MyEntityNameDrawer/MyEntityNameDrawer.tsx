import { SetupContext, defineComponent } from 'vue'
import BaseDrawer from '@/components/BaseDrawer/BaseDrawer'
import styles from './MyEntityNameDrawer.module.scss'
import { useMyEntityNameDrawer } from '../../../../Controllers/MyEntityNameDrawer'
import DyForm from '@/components/DyForm/DyForm'

// @ts-ignore
export default defineComponent<{
  [key: string]: any
}>({
  name: '弹窗',
  props: {
    modelValue: {
      type: Boolean,
      default: false,
    },
    title: {
      type: String,
      default: '',
    },
    row: {
      type: Object,
    },
    sort: {
      type: Number,
      default: 0,
    },
  },
  emits: ['update:modelValue', 'close', 'submit', 'confirm'],
  setup(props: Record<string, any>, ctx: SetupContext) {
    const {
      onClose,
      onConfirm,
      onOpen,
      formRef,
      visible,
      formItems,
      formData,
    } = useMyEntityNameDrawer(props, ctx)
    return () => (
      <BaseDrawer
        class={styles.drawer}
        size="800px"
        title={props.title || '添加'}
        v-model={visible.value}
        close-on-click-modal={true}
        onConfirm={onConfirm}
        onOpen={onOpen}
        before-close={onClose}
        onClose={onClose}
      >
        <DyForm
          ref={formRef}
          formData={formData.value}
          labelWidth="106px"
          formItemProps={formItems}
        ></DyForm>
      </BaseDrawer>
    )
  },
})
