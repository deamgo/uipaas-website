export interface mcontent {
  id: string
  title: string
  path?: string
  matcher?: string
  index?: number
  icon?: React.ReactElement
}

export interface IMultiplySelectorPropsItem {
  id: string
  text: string
  path: string
  type: 'normal' | 'error'
  method: () => void
  children?: React.ReactNode
}

export interface IWorkSpaceMate {
  id: string
  name: string
  logo: string
  lable: string
  description: string
  creator: string
  createTime?: string
  updateTime?: string
  status?: number
}

export interface IWorkspaceItemProps {
  id: string
  logo: string
  name: string
}

export interface IPaginationProps {
  pages: number
  total: number
  current: number
  onCurrentPageChange: (page: number) => void
}

export interface ISelectOption {
  id: string
  value: string
  label: string
}
export interface ISelectProps {
  id?: string
  default: string
  list: ISelectOption[]
  onChange?: (value: string, developer_id?: string) => void
  onOpen?: () => void
  children?: React.ReactNode
}

export interface IDownListProps {
  id?: string
  list: ISelectOption[]
  onChange?: (value: string, developer_id?: string) => void
  onOpen?: () => void
  children?: React.ReactNode
}