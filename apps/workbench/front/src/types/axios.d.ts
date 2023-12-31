import { IWorkspaceItemProps } from '@/interface/some';
import 'axios'

declare module 'axios' {
  interface AxiosResponse<T = any> extends Promise<T> {
    value: {
      code: number;
      msg: string
      data: {
        token: string
        code_key: string
        id: string
        username: string
        email: string
        avatar: string
      } |
      IWorkspaceItemProps[] |
      IWorkspaceItemProps |
      string |
      {
        records: [],
        total: number
      }
    }
    code: number;
    msg: string;
  }
}