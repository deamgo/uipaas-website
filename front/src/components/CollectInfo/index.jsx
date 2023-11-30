import "./index.less";
import vectorIcon from "../../assets/Vector.svg";
import React, { useState, useEffect } from "react";
import $message from "../Msg";
import { saveCompInfo } from "../../api/comp_info";
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
      error = "公司名称不能为空";
    }
    if (!formData.companySize.trim()) {
      error = "公司规模不能为空";
    }
    if (!formData.name.trim()) {
      error = "姓名不能为空";
    }
    const emailREG = /^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/;
    if (!emailREG.test(formData.businessEmail)) {
      error = "邮箱不合法";
    }
    if (!formData.requirementDescription.trim()) {
      error = "需求描述不能为空";
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
      return;
    }
    saveCompInfo(formData);
  };

  return (
    <>
      {(
        <div className="info_collect_card">
          <div className="vector_icon">
            <img src={vectorIcon} onClick={onClose} />
          </div>
          <div className="title">Hello, future partners!</div>
          <form action="" onSubmit={handleSubmit}>
            <div>
              <label htmlFor="companyName">
                <span>*</span>Company
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
                <span>*</span>Company Size
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
                <span>*</span>Name
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
                <span>*</span>Business email
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
                <span>*</span>Requirement description
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
            <span style={{ color: "#ea000" }}>*</span> By submitting, you agree
            to the <a href="privacy">privacy policy</a>.
          </span>
        </div>
      )}
    </>
  );
}
