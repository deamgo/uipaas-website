import React from "react";
//style
import './index.less'
import Guide from "./guide";
import Header from "./header";
import Content from "./content";
//


const ContentApp: React.FC = () => {

  const [fstGuide, setFstGuide] = React.useState(true)

  return (
    <>
      {fstGuide && (
        <>
          <div className="__capp_guide">
            <Guide handleClosed={() => { setFstGuide(false) }} />
          </div>
        </>
      )}
      <div className="__capp_header" style={fstGuide ? {
        top: '213rem'
      } : {}}>
        <Header />
      </div>
      <div className="__capp_content" style={fstGuide ? {
        top: '285rem',
        height: 984 - 285 + 72 + 'rem'
      } : {}}>
        <Content />
      </div>
    </>
  )
}

export default ContentApp