import "./index.less";
import vectorIcon from "../../assets/Vector.svg";
import React, { useState, useEffect } from "react";
import $message from "../Msg";
import { saveCompInfo, test } from "../../api/comp_info";
export function CollectInfo({onClose}) {
  // 状态
  const [formData, setFormData] = useState({
    companyName: "",
    companySize: "",
    name: "",
    businessEmail: "",
    requirementDescription: "",
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
    if (!formData.companyName.trim()) {
      error = "CompanyName cant be empty";
    }
    if (!formData.companySize.trim()) {
      error = "CompanySize cant be empty";
    }
    if (!formData.name.trim()) {
      error = "Name cant be empty";
    }
    const emailREG = /^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/;
    if (!emailREG.test(formData.businessEmail)) {
      error = "Invalid business email";
    }
    if (!formData.requirementDescription.trim()) {
      error = "RequirementDescription cant be empty";
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
    // saveCompInfo(formData);
    console.log(formData);
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
              <label htmlFor="companyName">
                <span>Company</span><span>*</span>
              </label>
              <input
                type="text"
                name="companyName"
                id="companyName"
                value={formData.companyName}
                onChange={handleChange}
              />
            </div>
            <div>
              <label htmlFor="companySize">
                <span>Team Size</span><span>*</span>
              </label>
              <input
                type="text"
                name="companySize"
                id="companySize"
                value={formData.companySize}
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
              <label htmlFor="businessEmail">
                <span>Business email</span><span>*</span>
              </label>
              <input
                type="text"
                name="businessEmail"
                id="businessEmail"
                value={formData.businessEmail}
                onChange={handleChange}
              />
            </div>
            <div>
              <label htmlFor="requirementDescription">
                <span>Description of Requirements</span><span>*</span>
              </label>
              <textarea
                name="requirementDescription"
                id="requirementDescription"
                cols="30"
                rows="5"
                className="desc_textarea"
                value={formData.requirementDescription}
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
