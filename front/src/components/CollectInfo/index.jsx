import "./index.less";
import vectorIcon from "../../assets/Vector.svg";
import React, { useState, useEffect } from "react";
import $message from "../Msg";
import { saveCompInfo } from "../../api/comp_info";
export function CollectInfo({ onClose }) {
  // 状态
  const [formData, setFormData] = useState({
    companyname: "",
    companysize: "",
    name: "",
    businessemail: "",
    requirementdescription: "",
  });
  //   样式
  // const [isShow, setIsShow] = useState(true);
  //   消息提醒
  const [messages, setMessages] = useState([]);

  // const handleClose = () => {
  //   console.log("close");
  //   setIsShow(false);
  // };

  //   input输入框验证
  const validateInput = () => {
    let error = "";

    // 根据字段名进行不同的验证
    if (!formData.companyname.trim()) {
      error = "companyname cant be empty";
    }
    const cnREG = /^[\u4e00-\u9fa5a-zA-Z0-9]{1,30}$/;
    if (!cnREG.test(formData.companyname)) {
      error = 'Invalid company name'
    }
    if (!formData.companysize.trim()) {
      error = "companysize cant be empty";
    }
    const csREG = /^[0-9]+([.][0-9]{1,10})?$/;
    if (!csREG.test(formData.companysize)) {
      error = 'Invalid company size'
    }
    if (!formData.name.trim()) {
      error = "Name cant be empty";
    }
    const nREG = /^[\u4e00-\u9fa5a-zA-Z0-9]{1,30}$/;
    if (!nREG.test(formData.name)) {
      error = 'Invalid name'
    }
    if (!formData.businessemail.trim()) {
      error = 'businessemail cant be empty'
    }
    const emailREG = /^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/;
    if (!emailREG.test(formData.businessemail)) {
      error = "Invalid business email";
    }
    if (!formData.requirementdescription.trim()) {
      error = "requirementdescription cant be empty";
    }
    const rdREG = /^[\u4e00-\u9fa5a-zA-Z0-9]{1,200}$/;
    if (!rdREG.test(formData.requirementdescription)) {
      error = 'Invalid requirement description';
    }
    setMessages(error);
    return error;
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    const error = validateInput();
    if (error) {
      $message.error(error);
      return
    }
    saveCompInfo(formData).then((res) => {
      console.log(res);
      if (res.value.code === 0) {
        onClose();
        $message.success(res.value.msg);
      } else if (res.value.code === -1) {
        $message.error(res.value.msg);
      }
    })
      .catch((err) => {
        console.log(err);
        $message.error(err.response.data.value.msg);
      })
  };

  return (
    <>
      {(
        <div className="info_collect_card">
          <div className="vector_icon" onClick={onClose} >
            <img src={vectorIcon} />
          </div>
          <div className="title">Hello, future partners!</div>
          <form action="" onSubmit={handleSubmit}>
            <div>
              <label htmlFor="companyname">
                <span>Company</span><span>*</span>
              </label>
              <input
                type="text"
                name="companyname"
                id="companyname"
                value={formData.companyname}
                onChange={handleChange}
              />
            </div>
            <div>
              <label htmlFor="companysize">
                <span>Team Size</span><span>*</span>
              </label>
              <input
                type="text"
                name="companysize"
                id="companysize"
                value={formData.companysize}
                onChange={handleChange}
              />
            </div>
            <div>
              <label htmlFor="name">
                <span>Name</span><span>*</span>
              </label>
              <input
                type="text"
                name="name"
                id="name"
                value={formData.name}
                onChange={handleChange}
              />
            </div>
            <div>
              <label htmlFor="businessemail">
                <span>Business email</span><span>*</span>
              </label>
              <input
                type="text"
                name="businessemail"
                id="businessemail"
                value={formData.businessemail}
                onChange={handleChange}
              />
            </div>
            <div>
              <label htmlFor="requirementdescription">
                <span>Description of Requirements</span><span>*</span>
              </label>
              <textarea
                name="requirementdescription"
                id="requirementdescription"
                cols="30"
                rows="5"
                className="desc_textarea"
                value={formData.requirementdescription}
                onChange={handleChange}
              ></textarea>
            </div>
            <div>
              <button
                type="submit"
                className="sub_btn"
                onClick={handleSubmit}>
                Submit
              </button>
            </div>
          </form>
          <span className="foot_text">
            By submitting, you agree
            to the <a href="privacy">privacy policy</a>.
          </span>
        </div>
      )}
    </>
  );
}
