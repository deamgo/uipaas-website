import React from 'react'
//
import './index.less'
//

interface IDividerProps {
  ys?: React.CSSProperties
  children?: React.ReactNode
}

const Divider: React.FC<IDividerProps> = (props) => {
  return (
    <>
      <div style={props.ys} className="__divider"></div>
    </>
  )
}

export default Divider