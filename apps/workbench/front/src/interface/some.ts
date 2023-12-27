export interface mcontent {
  id: string;
  title: string;
  path: string;
  matcher: string;
  index: number;
  icon: React.ReactElement;
}

export interface IMultiplySelectorPropsItem {
  id: number
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
  id: string;
  logo: string;
  name: string;
}