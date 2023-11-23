import "./App.css";
import zinclSvg from "./assets/zinclabsolution.svg";
import { Paper } from "./components/Paper";


function App() {
  return (
    <>
      <header>
        <div className="head"></div>
      </header>
      <main>
        <div className="container">
          <span className="title">
            AI原生
            <br />
            面向未来的可编排数字化平台
          </span>
          <span className="gary">软件开发团队的敏捷交付神器</span>
        </div>
        <div className="container">
          <h1 className="text_block_headline">产品特性</h1>
          <div className="main_container">
            <div className="img_container">
              <img src={zinclSvg} className="zinclabsolution" />
            </div>
            <div className="text_dsp_block">
              <Paper title={"AI原生"} content={"打造应用开发新范式"} style={{color:'red'}} />
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
      </footer>
    </>
  );
}

export default App;
