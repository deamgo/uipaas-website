const emailReg = /^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/
const codeReg = /^[A-Za-z0-9]{6}$/
const passwordReg = /^[A-Za-z0-9]{8,20}$/
const usernameReg = /^[\u4e00-\u9fa5a-zA-Z0-9]{6,12}$/
const emailVerificationReg = /^[0-9]{4}$/

export {
  emailReg,
  codeReg,
  passwordReg,
  usernameReg,
  emailVerificationReg
}