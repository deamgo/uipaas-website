import { useEffect } from "react"
import { BrowserRouter as Routes, Route } from 'react-router-dom'
import { Navigate } from 'react-router-dom'
import './App.less'
import { getComponyList } from './api/comp_info'

//layout
import StickyHeadTable from './layout/Table'

//svg
import logo from './assets/logov1.svg'

const Admin = () => {

  // const getCompanys = () => {
  //   getComponyList({
  //     pageSize: 10,
  //     pageNum: 1
  //   }).then(res => {
  //     console.log(res);
  //   })
  // }

  // useEffect(() => {
  //   window.addEventListener('resize', resize)
  //   resize()
  // },[])

  // const resize = () => {
  //   let IwD = window.innerWidth
  //   document.documentElement.style.fontSize = IwD/1440 + 'px'
  // }

  const isAuthenticated = localStorage.getItem('jwtToken') !== null

  return (
    <>
      <div className="siderbar">
        <div className="logo-wapper">
          <img src={logo} alt="UIPaaS-Logo" />
        </div>
        <div className="sider-menu">
          <div className="guide">
            <span>内容管理</span>
          </div>
        </div>
      </div>
      <div className="head">
        <div className="cicle"></div>
        <span>管理员</span>
      </div>
      <div className="main">
        <StickyHeadTable />
      </div>
      {/* <button onClick={getCompanys}>BUTTON</button> */}
      {!isAuthenticated && <Navigate to="/logo" />}
    </>
  )
}

export default Admin