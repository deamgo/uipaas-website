import { useEffect } from 'react'
import '../App.less'
import './pp.less'
import logoSvg from '../assets/logo.svg'
import PrivacySvg from '../assets/privacy.svg'

const PriPolicy = () => {

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
            <header>
                <div 
                className='head'>
                    <img className='logo' src={logoSvg} alt="Logo" />
                </div>
            </header>
            <main>
                {/* <div className="content">
                    <span className='title'>Privacy Policy</span>
                    <span className='time'>Last Modified: February 27, 2023</span>
                    <p className='passage'>This Privacy Agreement (&quot;Agreement&quot;) governs the use of our websites, mobile applications and other services. Please read the following carefully to understand how we collect, use, protect and disclose your personal information.</p>
                </div> */}
                <div className="content">
                    <img src={PrivacySvg} alt="Privacy-Policy" />
                </div>
            </main>
        </>
    )
}

export default PriPolicy