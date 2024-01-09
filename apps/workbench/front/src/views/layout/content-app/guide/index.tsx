import React from "react";
//
import './index.less'
//

interface GuideProps {
  handleClosed: () => void;
}

const Guide: React.FC<GuideProps> = (props) => {
  return (
    <>
      <div className="__guide">
        <h1 className="one">Step 1</h1>
        {/* <h1 className="two">Step 2</h1>
        <h1 className="thd">Step 3</h1> */}
        <span className="one">Create your worksapce</span>
        {/* <span className="two">Create your Applications</span>
        <span className="thd">Invite Co-developer</span> */}
        {/* <div className="__guide_svg one">
          <DArrowR />
        </div>
        <div className="__guide_svg two">
          <DArrowR />
        </div> */}
        <div className="__guide_closed" onClick={props.handleClosed}>
          <span>X</span>
        </div>
      </div>
    </>
  )
}

export default Guide