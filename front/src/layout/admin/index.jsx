import { useEffect } from "react" 

const Admin = () => {

  useEffect(() => {
    window.addEventListener('resize', resize)
    resize()
  },[])

  const resize = () => {
    let IwD = window.innerWidth
    document.documentElement.style.fontSize = IwD/1440 + 'px'
  }

  return (
    <>
      <div className="siderbar">
        <div className="logo-wapper"></div>
        <div className="sider-menu"></div>
      </div>
      <div className="head"></div>
      <div className="main"></div>
    </>
  )
}

export default Admin