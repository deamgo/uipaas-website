import './index.css'

export function Paper({ title, content }) {
    return (
      <div className="ctr">
        <span className="title_block">{title}</span>
        <br />
        <span className="content">{content}</span>
      </div>
    );
  } 