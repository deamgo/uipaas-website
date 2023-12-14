import React from 'react'
import { resize } from '@utils/adapt'
//style
import './index.less'
//
import signLogoSvg from '@assets/sign/sign-logo.svg'



const Privacy: React.FC = () => {

  React.useEffect(() => {
    window.addEventListener('resize', resize)
    resize()
  })

  return (
    <>
      <div className="__privacy">
        <div className="__privacy_logo">
          <img src={signLogoSvg} alt="UIPaaS" />
        </div>
        <div className="__privacy_context">
          <h1 className="__privacy_context_title">
            Privacy Policy
          </h1>
          <span className="__privacy_context_date">
            Last Modified: February 27, 2023
          </span>
          <p className="__privacy_context_abstract">
            This Privacy Agreement ("Agreement") governs the use of our websites, mobile applications and other services. Please read the following carefully to understand how we collect, use, protect and disclose your personal information.
          </p>

          <h2 className="__privacy_context_title_sec">
            Information Collection
          </h2>
          <p className="__privacy_context_detail">
            We may collect personal information that you provide, including, but not limited to, name, email address, contact information, geographic location, IP address and device information. We will take reasonable steps to protect the security of the personal information you provide.
          </p>

          <h2 className="__privacy_context_title_sec">
            Use of Information
          </h2>
          <p className="__privacy_context_detail">
            We will use your personal information to provide and improve our services and to send you relevant notices. We may use your personal information for marketing and promotional activities, but we will respect your privacy and you have the right to opt out of receiving these notices.
          </p>

          <h2 className="__privacy_context_title_sec">
            Disclosure of Information
          </h2>
          <p className="__privacy_context_detail">
            We may share your personal information with third party partners to provide better services. We will take reasonable steps to protect your personal information and require third party partners to similarly comply with this agreement.
          </p>

          <h2 className="__privacy_context_title_sec">
            Cookies and Other Technologies
          </h2>
          <p className="__privacy_context_detail">
            We may use cookies and other similar technologies to collect and store information about you. These technologies help us provide a better user experience and are used to analyze and improve our services.
          </p>

          <h2 className="__privacy_context_title_sec">
            Privacy
          </h2>
          <p className="__privacy_context_detail">
            We will take reasonable steps to protect the security of your personal information from unauthorized access, use or disclosure. However, due to the open nature of the Internet, we cannot fully guarantee the absolute security of information.
          </p>

          <h2 className="__privacy_context_title_sec">
            Protection of Minors
          </h2>
          <p className="__privacy_context_detail">
            Our services do not apply to minors. If you are under the legal age, please use our services under the guidance of a parent or guardian.
          </p>

          <h2 className="__privacy_context_title_sec">
            Changes to our Privacy Policy
          </h2>
          <p className="__privacy_context_detail">
            We may modify this Privacy Policy from time to time and will post the updated version on our website or mobile application. Please check back periodically to keep up to date with the latest Privacy Policy.
          </p>

          <h2 className="__privacy_context_title_sec">
            Contact Us
          </h2>
          <p className="__privacy_context_detail">
            If you have any questions or comments about our Privacy Policy, please contact us via our website or mobile application.
          </p>

          <p className="__privacy_context_tail">
            Please read and understand this Privacy Agreement carefully before using our services. By continuing to use our services, you agree to all the terms and conditions of this agreement.
          </p>

        </div>
      </div>
    </>
  )
}

export default Privacy