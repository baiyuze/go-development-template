import MyPluginName from './Views/MyPluginName'
import Setting from '@/components/Setting/Setting'
import { provider } from '@/provider/index'
import p from '../../assets/svg/p.svg'

export default {
  is: 'MyPluginName',
  name: '${{widgetName}}',
  category: 'run',
  icon: p,
  authorizationRequired: false,
  canvasView: provider(MyPluginName),
  settingsView: Setting,
}
