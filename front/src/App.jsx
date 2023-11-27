// import "./App.css";
import zinclSvg from "./assets/zinclabsolution.svg";
import { Paper } from "./components/Paper";

import { useState, useEffect } from 'react'
import './App.less'
import featureSVG from './assets/feature.svg'
import logoSvg from './assets/logo.svg'
import { CollectInfo } from "./components/CollectInfo";
// import "./App.css";
import zinclSvg from "./assets/zinclabsolution.svg";
import { Paper } from "./components/Paper";

import { useState, useEffect } from 'react'
import './App.less'
import featureSVG from './assets/feature.svg'
import logoSvg from './assets/logo.svg'
import { CollectInfo } from "./components/CollectInfo";

function App() {

  // const [__HG, setHG] = useState(1440);
  
  useEffect(() => {
    let WD = window.screen.width
    let IwD = window.innerWidth
    let InH = window.innerHeight
    document.querySelectorAll('.container').forEach(el => {
      el.style.height = `${InH}px`
    })
    document.documentElement.style.fontSize = IwD/1440 + 'px'
    // setHG(HG)
  },[])


  // const [__HG, setHG] = useState(1440);
  
  useEffect(() => {
    let WD = window.screen.width
    let IwD = window.innerWidth
    let InH = window.innerHeight
    document.querySelectorAll('.container').forEach(el => {
      el.style.height = `${InH}px`
    })
    document.documentElement.style.fontSize = IwD/1440 + 'px'
    // setHG(HG)
  },[])


  return (
    <>
      <header>
        <div 
          className='head'>
            <img className='logo' src={logoSvg} alt="Logo" />
          </div>
        <div 
          className='head'>
            <img className='logo' src={logoSvg} alt="Logo" />
          </div>
      </header>
      <main>
        <div className="container">
          <div className='text_dsp_block'>
            <span className="title">
            AI原生
            <br />
            面向未来的可编排数字化平台
          </span>
            <span className="gary">软件开发团队的敏捷交付神器</span>
            <span className='mark'>即将上线...</span>
            <button className='reserveButton'>
              <span>立即预约体验</span>
            </button>
          </div>
          <div className='svg_block'>
              <img src={featureSVG} alt='feature' />
          </div>
        </div>
        <div className="container">
          <h1 className="text_block_headline">产品特性</h1>
          <div className="main_container">
            <div className="img_container">
              <img src={zinclSvg} className="zinclabsolution" />
            </div>
            <div className="text_dsp_block">
              <Paper title={"AI原生"} content={"打造应用开发新范式"} />
              <Paper
                title={"更易用的模型驱动型产品"}
                content={"专注业务，无需专业建模知识。"}
              />
              <Paper
                title={"应用独立部署"}
                content={"支持应用脱离平台独立运行，一键快速部署"}
              />
              <Paper
                title={"可拓展性强"}
                content={"支持自定义组件，可插拔式集成"}
              />
            </div>
          </div>
        </div>
        <div className="container"></div>
        <div className="tail"></div>
      </main>
      <footer>
        <div className="foot"></div>
        <div className="foot"></div>
      </footer>
    </>
  );
  );
}

export default App;
export default App;
