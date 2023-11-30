//components
import { Paper } from './components/Paper'
import { CollectInfo } from './components/CollectInfo'
//
import { useState, useEffect } from 'react'
//style
import './App.less'
//svg
import logoSvg from './assets/logov1.svg'
import bannerSVG from './assets/ab/banner-pic.svg'
import aboutSvg from './assets/ab/aboutus.svg'
import aiiSvg from './assets/ab/ai-icon.svg'
import modiSvg from './assets/ab/model-icon.svg'
import serviSvg from './assets/ab/server-icon.svg'
import sectiSvg from './assets/ab/section-icon.svg'

let visAble = false

function App() {

  const [btnS, setbtnS] = useState({
    position: 'absolute',
  });

  const [btnVis, setbtnVis] = useState({
    visibility: 'hidden',
  })

  const [tabVis, settabVis] = useState(visAble)
  
  useEffect(() => {
    window.addEventListener('resize', resize)
    resize()
    scrollDetector()
  },[])

  const resize = () => {
    let IwD = window.innerWidth
    document.documentElement.style.fontSize = IwD/1920 + 'px'
  }
  
  const scrollDetector = () => {
    window.addEventListener('scroll', () => {
      let Hbtn = document.querySelector('#hBtn')
      let Hbtn_T = Hbtn.offsetTop
      let Hbtn_H = Hbtn.offsetHeight
      let tTH = Hbtn_H + Hbtn_T
      if (window.scrollY >= tTH) {
        setbtnS({
          position: 'fixed'
        })
        setbtnVis({
          display: 'block'
        })
      } else {
        setbtnS({
          position: 'absolute'
        })
        setbtnVis({
          display: 'none'
        })
      }
    })
  }

  const handleClick = () => {
    visAble = !visAble
    settabVis(visAble)
    document.body.style.overflow = visAble ? 'hidden' : 'auto'
    document.body.style.position = visAble ? 'fixed' : ''
  }

  return (
    <>
      <header  style={btnS}>
        <div 
          className='head'>
            <img className='logo' src={logoSvg} alt="Logo" />
              <button 
                className='reserveButton' 
                style={btnVis}
                onClick={handleClick}>
                  <span>Book demo</span>
              </button>
            {/* <div className="reserve" id="headBtn" style={btnVis}>
            </div> */}
        </div>
      </header>
      <main>
        <div className="container">
          <div className="title">
            <span>
              AI-base, future-oriented, 
              <br />
              drag-and-drop
              <br />
              development platform.
            </span>
          </div>
          <div className="dsp">
            <span>
              The agile delivery tool for software development teams.
              <br/>
              Delivering software promptly with a business-oriented approach.
            </span>
          </div>
            <button
              className='reserveButton' 
              id="hBtn"
              onClick={handleClick}>
                <span>Book demo</span>
            </button>
          {/* <div className="reserve" id="homeBtn">
          </div> */}
          <div className='svg_block'>
              <img src={bannerSVG} alt='feature' />
          </div>
          <div className="sec_title">
            <span>Feature</span>
          </div>
          <div id="icons">
            <img src={aiiSvg} alt="AI base" />
            <img src={modiSvg} alt="Easier use" />
            <img src={serviSvg} alt="deploy" />
            <img src={sectiSvg} alt="highly" />
          </div>
          <Paper 
            title={"AI base"} 
            content={"Creating a New Paradigm for Application Development"} 
            ps={{
              left: '395rem',
              top: '1065rem',
            }}
          />
          <Paper
            title={"Easier-to-use model-driven products"}
            content={"Focus on business, no modeling expertise required"}
            ps={{
              left: '1202rem',
              top: '1065rem',
            }}
          />
          <Paper
            title={"Application-independent deployment"}
            content={"Support applications to run independently from the platform, one-click rapid deployment"}
            ps={{
              left: '395rem',
              top: '1269rem',
            }}
          />
          <Paper
            title={"Highly expandable"}
            content={"Support for custom components, pluggable integration"}
            ps={{
              left: '1202rem',
              top: '1269rem',
            }}
          />
          <div className="thd_title">
            <span>About us</span>
          </div>
          <div className="thd_svg_block">
            <img src={aboutSvg} alt="About" />
            <span className="thd_dsp_one">
              A team of senior technical experts 
              <br />
              in the low-code field
            </span>
            <span className="thd_dsp_two">
              Our Vision is to provide development teams with a 
              <br />
              new development tool that can be used to improve 
              <br />
              the efficiency of delivering</span>
          </div>
        </div>
        <div className="tail">
          <span className="title">
            AI-base, future-oriented,
            <br />
            drag-and-drop  development platform.
          </span>
          <button 
            className="reserveButton"
            onClick={handleClick}>
            <span>Book demo</span>
          </button>
        </div>
      </main>
      <footer>
        <div className="foot">
          <span className="cpr">Copyright © Deamoy Technology</span>
          <span className="ba">浙ICP备2021001545号-4</span>
        </div>
      </footer>
      {tabVis && <CollectInfo onClose={handleClick}/>}
    </>
  ); 
}

export default App;
