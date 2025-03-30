import { ref, onMounted, reactive, computed, Ref, watch } from 'vue'
import { injectModel } from '@/libs/Provider/Provider'
import { MyEntityNameDrawer } from '../Models/MyEntityNameDrawer'
import { ElMessage } from 'element-plus'
import isEqual from 'lodash/isEqual'
import { ConfirmBox } from '@/components/ConfirmBox/ConfirmBox'
import { cloneDeep } from 'lodash'

export const useMyEntityNameDrawer = (props: any, ctx?: any) => {
  const myEntityNameDrawer =
    injectModel<MyEntityNameDrawer>('myEntityNameDrawer')
  /**
   * 用来对比的初始化数据
   */
  const initiateData: Ref<Record<string, any>> = ref({})
  const formData = ref<Record<string, any>>({})
  // ref
  const formRef = ref()

  const current = computed(() => {
    return props.row || null
  })
  const visible = computed({
    get() {
      return props.modelValue
    },
    set(val) {
      ctx.emit('update:modelValue', val)
    },
  })
  /**
   * 添加的form字段
   */
  const formItems = reactive([
    {
      label: '名称',
      prop: 'name',
      el: 'input',
      placeholder: '请输入名称',
      rules: [{ required: true, message: '名称', trigger: 'blur' }],
    },
    {
      label: '编号',
      prop: 'code',
      el: 'input',
      placeholder: '请输入编号',
      rules: [{ required: true, message: '编号', trigger: 'blur' }],
    },
    {
      label: '备注',
      prop: 'remark',
      el: 'input',
      placeholder: '请输入备注',
    },
  ])
  /**
   * 校验是否有数据变化
   */
  const checkIsEqualObject = () => {
    const data = {
      formData: formData.value,
    }
    const check = isEqual(initiateData.value, data)
    return check
  }

  const onClose = (done: () => void) => {
    if (visible.value) {
      if (checkIsEqualObject()) {
        visible.value = false
        done && done()
      } else {
        ConfirmBox('是否保存设置？')
          .then(() => {
            onConfirm()
          })
          .catch(() => {
            visible.value = false
            done && done()
          })
      }
    }
  }
  /**
   * 保存
   */
  const onConfirm = async () => {
    await formRef.value?.validate()
    const data = {
      name: formData.value.name,
      code: formData.value.code,
      remark: formData.value.remark,
      sort: props.sort,
    }
    if (!current.value) {
      await myEntityNameDrawer.addMyEntityName(data)
    } else {
      const id = current.value.id
      await myEntityNameDrawer.updateMyEntityName(id, data)
    }
    ElMessage.success('保存成功')
    ctx.emit('confirm')
  }

  const updateCheckData = () => {
    initiateData.value = {
      formData: {
        ...formData.value,
      },
    }
  }
  /**
   * 弹窗打开获取详情
   */
  const onOpen = async () => {
    if (current.value) {
      const res = await myEntityNameDrawer.getMyEntityNameDetail(current.value)

      formData.value = {
        name: res.name,
        code: res.code,
        remark: res.remark,
        id: res.id,
      }
      updateCheckData()
    } else {
      formData.value = {}
      updateCheckData()
    }
  }

  watch(() => current.value, onOpen)

  return {
    formItems,
    formData,
    visible,
    formRef,
    onOpen,
    onClose,
    onConfirm,
  }
}
