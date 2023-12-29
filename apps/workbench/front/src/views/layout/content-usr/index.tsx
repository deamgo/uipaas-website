import React from "react";
//style
import './index.less'
import Header from "./header";
import Content from "./content";
//


const ContentUsr: React.FC = () => {


  return (
    <>
      <div className="__cusr">
        <div className="__cusr_header">
          <Header />
        </div>
        <div className="__cusr_content">
          <Content />
        </div>
      </div>
    </>
  )
}

export default ContentUsr