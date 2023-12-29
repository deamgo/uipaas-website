import React from 'react'
//
import './index.less'
import { Link, useLocation } from 'react-router-dom'
//

interface INavbarItem {
  path: string
  matcher: string
  name: string
  index: number
}

interface INavbarProps {
  items: INavbarItem[]
}

const Navbar: React.FC<INavbarProps> = (props) => {

  const [active, setActive] = React.useState('')
  const location = useLocation()

  React.useEffect(() => {
    const url = location.pathname
    let urlArr = url.split('/')
    console.log(urlArr);
    if (urlArr.length < 3) {
      console.log(urlArr[1]);

      setActive(urlArr[1])
    } else if (urlArr.length === 3) {
      setActive(urlArr[2])
    }
    // if (props.path && urlArr[props.index] === props.matcher) {
    //   setActive()
    // } else {
    //   setActive(false)
    // }
  }, [location.pathname])

  return (
    <>
      <ul className="__navbar">
        {props.items.map(item => (
          <li key={item.name} className={`__navbar_item ${active === item.matcher ? '__navbar_active' : ''}`}>
            <Link to={item.path}>{item.name}</Link>
          </li>
        ))}
      </ul>
    </>
  )
}

export default Navbar