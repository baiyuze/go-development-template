import { Component } from 'vue'

export interface DataItemType {
  id?: string
  name?: string
  code?: string
  description?: string
  label?: string
  value?: string | number
}

export interface MyEntityNameBaseType {
  id?: string
  name?: string
  code?: string
  value?: number
  description?: string
  options?: Array<DataItemType>
  abilityValue?: number | string
  data?: DataItemType
  defaultValue?: string | number
  flow: string
}

export interface FlowDefinitionType {
  id?: string
  name?: string
  code?: string
  description?: string
}

export type ModuleType = Record<
  string,
  {
    default: Record<string, string>
    name: string
  }
>

export interface TabItem {
  name: string
  label: string
  component: Component
  hidden?: boolean
}
