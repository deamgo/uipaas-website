export function validate(formData) {
  let error;

  const rdREG = /^[\u4e00-\u9fa5a-zA-Z0-9]{1,200}$/;
  if (!rdREG.test(formData.requirementdescription)) {
    error = 'Invalid requirement description';
  }
  if (!formData.requirementdescription.trim()) {
    error = "requirementdescription cant be empty";
  }

  const emailREG = /^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/;
  if (!emailREG.test(formData.businessemail)) {
    error = "Invalid business email";
  }
  if (!formData.businessemail.trim()) {
    error = 'businessemail cant be empty'
  }

  const nREG = /^[\u4e00-\u9fa5a-zA-Z0-9]{1,30}$/;
  if (!nREG.test(formData.name)) {
    error = 'Invalid name'
  }
  if (!formData.name.trim()) {
    error = "Name cant be empty";
  }

  const csREG = /^[0-9]+([.][0-9]{1,10})?$/;
  if (!csREG.test(formData.companysize)) {
    error = 'Invalid company size'
  }
  if (!formData.companysize.trim()) {
    error = "companysize cant be empty";
  }

  const cnREG = /^[\u4e00-\u9fa5a-zA-Z0-9]{1,30}$/;
  if (!cnREG.test(formData.companyname)) {
    error = 'Invalid company name'
  }
  if (!formData.companyname.trim()) {
    error = "companyname cant be empty";
  }

  return error;
}