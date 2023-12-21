import React from "react"
//
import './index.less'
//
import MenuContent from '@components/MenuContent'

interface mcontent {
  id: number;
  title: string;
  path: string;
  icon: React.ReactElement;
}

type SideMenuProps = {
  list: mcontent[]
}

const SideMenu: React.FC<SideMenuProps> = (props) => {
  return (
    <>
      <div className="__smenu">
        {props.list
          ? props.list.map(item => (
            <MenuContent id={item.id} title={item.title} icon={item.icon} path={item.path} />
          ))
          : (<></>)}
      </div>
    </>
  )
}

export default SideMenu