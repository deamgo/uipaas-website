import React from 'react';
//

interface MsgShowerProps {
  text: string
  type: string
}

const MsgShower: React.FC<MsgShowerProps> = (props) => {
  return (
    <>
      <div className={`message ${props.type}`}>
        <span className='icon' />
        <span>{props.text}</span>
      </div>
    </>
  )
}

export default MsgShower