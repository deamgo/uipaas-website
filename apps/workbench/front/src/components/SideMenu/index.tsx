import React from "react"
//
import './index.less'
//
import MenuContent from '@components/MenuContent'
import { mcontent } from "@/interface/some"

type SideMenuProps = {
  title?: string,
  list: mcontent[]
}


const SideMenu: React.FC<SideMenuProps> = (props) => {
  return (
    <>
      <div className="__sidermenu_list">
        {props.title && (
          <div className="__sidermenu_list_title">
            {props.title}
          </div>
        )}
        {props.list && props.list.map(item => (
          <MenuContent
            key={item.id}
            id={item.id}
            title={item.title}
            path={item.path}
            matcher={item.matcher}
            index={item.index}
            icon={item.icon} />
        ))}
      </div>
    </>
  )
}

export default SideMenu