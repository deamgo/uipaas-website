import { useState, useEffect } from 'react'
import './App.less'
import featureSVG from './assets/feature.svg'
import logoSvg from './assets/logo.svg'

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


  return (
    <>
      <header>
        <div 
          className='head'>
            <img className='logo' src={logoSvg} alt="Logo" />
          </div>
      </header>
      <main>
        <div className='container'>
          <div className='text_dsp_block'>
            <span className='title'>AI原生<br/>面向未来的可编排数字化平台</span>
            <span className='gary'>软件开发团队的敏捷交付神器</span>
            <span className='mark'>即将上线...</span>
            <button className='reserveButton'>
              <span>立即预约体验</span>
            </button>
          </div>
          <div className='svg_block'>
              <img src={featureSVG} alt='feature' />
          </div>
        </div>
        <div className='container'></div>
        <div className='container'></div>
        <div className='tail'></div>
      </main>
      <footer>
        <div className='foot'></div>
      </footer>
    </>
  )
}

export default App
