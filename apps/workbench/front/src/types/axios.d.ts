import 'axios'

declare module 'axios' {
  interface AxiosResponse<T = any> extends Promise<T> {
    code: number;
    msg: string
    data: {
      Token: string
    } | {
      CodeKey: string
    } | string | null
  }
}