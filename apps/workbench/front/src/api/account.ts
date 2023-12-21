export type IUsrAccount = {
  invitation_code?: string
  id?: string | null
  username?: string | null
  email?: string | null
  old_email?: string | null
  password?: string | null
  code?: number
  code_key?: string | null
  avatar?: string | null
}

export type IUserInfo = {
  id: string
  username: string
  email: string
  avatar: string
}