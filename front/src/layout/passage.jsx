import zinclSvg from "./assets/zinclabsolution.svg";
import { Paper } from "./components/Paper";
import { Divider } from "./components/Divider";

import { useState, useEffect } from 'react'
import './App.less'
import featureSVG from './assets/feature.svg'
import logoSvg from './assets/logo.svg'
import aboutSvg from './assets/aboutushero.svg'


function Main() {

  return (
    <>
      <header  style={btnS}>
        <div 
          className='head'>
            <img className='logo' src={logoSvg} alt="Logo" />
            <div className="reserve" id="headBtn">
                <button className='reserveButton'>
                    <span>Book demo</span>
                </button>
            </div>
        </div>
      </header>
      <main>
        <div className="container">
          <div className="title">
            <span>
              AI-base, future-oriented, 
              <br />
              drag-and-drop  development platform.
            </span>
          </div>
          <div className="dsp">
            <span>
              The agile delivery tool for software development teams
              <br/>
              Delivering software promptly with a business-oriented approach.
            </span>
          </div>
          <div className="reserve" id="homeBtn">
            <button className='reserveButton'>
                <span>Book demo</span>
            </button>
          </div>
          <div className='svg_block'>
              <img src={featureSVG} alt='feature' />
          </div>
          <div className="gary">
            <span>Coming soon...</span>
          </div>
        </div>
        <Divider margin={{margin: '0 100rem'}}/>
        <div className="container">
          <div className="sec_title">
            <span>Feature</span>
          </div>
          <div className="sec_svg_block">
              <img src={zinclSvg} className="zinclabsolution" />
          </div>
          <Paper 
            title={"AI base"} 
            content={"Creating a New Paradigm for Application Development"} 
            ps={{
              left: '737rem',
              top: '214rem',
            }}
          />
          <Paper
            title={"Easier-to-use model-driven products"}
            content={"Focus on business, no modeling expertise required"}
            ps={{
              left: '737rem',
              top: '346rem',
            }}
          />
          <Paper
            title={"Application-independent deployment"}
            content={"Support applications to run independently from the platform, one-click rapid deployment"}
            ps={{
              left: '737rem',
              top: '477rem',
            }}
          />
          <Paper
            title={"Highly expandable"}
            content={"Support for custom components, pluggable integration"}
            ps={{
              left: '737rem',
              top: '644rem',
            }}
          />
        </div>
        <Divider margin={{margin: '0 100rem'}}/>
        <div className="container">
          <div className="thd_title">
            <span>About us</span>
          </div>
          <Paper
              title={"A team of senior technical experts in the low-code field"}
              ps={{
                left: '88rem',
                top: '356rem',
                width: '592rem',
                height: '104rem',
              }}
            />
          <Paper
              content={"Our Vision is to provide development teams with a new development tool that can be used to improve the efficiency of delivering"}
              ps={{
                left: '88rem',
                top: '478rem',
                width: '823rem',
                height: '70rem',
              }}
            />
          <div className="thd_svg_block">
            <img src={aboutSvg} alt="About" />
          </div>
        </div>
        <Divider margin={{margin: '0 100rem'}}/>
        <div className="tail">
          <span className="title">
            AI-base, future-oriented,
            <br />
            drag-and-drop  development platform.
          </span>
          <button className="reserveButton">
            <span>Book demo</span>
          </button>
        </div>
      </main>
      <footer>
        <div className="foot">
          <span>Copyright © Deamoy Technology</span>
          <span>浙ICP备2021001545号-4</span>
        </div>
      </footer>
    </>
  );
}

export default Main;