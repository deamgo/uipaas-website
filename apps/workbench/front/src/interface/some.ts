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