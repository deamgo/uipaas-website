import 'axios'

declare module 'axios' {
  interface AxiosResponse<T = any> extends Promise<T> {
    value: {
      code: number;
      msg: string
      data: {
        token: string
        code_key: string
      }
    }
  }
}