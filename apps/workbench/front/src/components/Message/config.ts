export type MessageApi = {
  info: (text: string) => void,
  success: (text: string) => void,
  warning: (text: string) => void,
  error: (text: string) => void,
};

export interface IMsgList {
  key: string
  text: string
  type: string
}
