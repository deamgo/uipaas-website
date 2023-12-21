import React from "react";
//style
import './index.less'
import Header from "./header";
import Content from "./content";
//


const ContentWorkSpace: React.FC = () => {


  return (
    <>
      <div className="__cwss_header">
        <Header />
      </div>
      <div className="__cwss_content">
        <Content />
      </div>
    </>
  )
}

export default ContentWorkSpace