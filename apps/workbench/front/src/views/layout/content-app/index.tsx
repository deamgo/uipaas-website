import React from "react";
//style
import './index.less'
import Guide from "./guide";
import Header from "./header";
import Content from "./content";
import { observer } from "mobx-react-lite";
//


const ContentApp: React.FC = () => {

  const [fstGuide, setFstGuide] = React.useState(true)


  return (
    <>
      <div className="__capp">
        {fstGuide && (
          <>
            <div className="__capp_guide">
              <Guide handleClosed={() => { setFstGuide(false) }} />
            </div>
          </>
        )}
        <div className="__capp_header">
          <Header />
        </div>
        <div className="__capp_content">
          <Content />
        </div>
      </div>
    </>
  )
}

export default observer(ContentApp)