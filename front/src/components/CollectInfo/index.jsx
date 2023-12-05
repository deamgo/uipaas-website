import "./index.less";
import vectorIcon from "../../assets/Vector.svg";
import React, { useState, useEffect } from "react";
import $message from "../Msg";
import { saveCompInfo } from "../../api/comp_info";
import { validate } from "./validator";

export function CollectInfo({ onClose }) {

  const [formData, setFormData] = useState({
    companyname: "",
    companysize: "",
    name: "",
    businessemail: "",
    requirementdescription: "",
  });

  const [messages, setMessages] = useState('');

  const validateInput = () => {
    let error = validate(formData);
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
        $message.error(
          err.response.data.value
            ? err.response.data.value.msg
            : 'System Error, please contact the backend'
        );
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
