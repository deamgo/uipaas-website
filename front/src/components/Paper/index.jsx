import './index.less'

export function Paper({ title, content, ps }) {
    return (
      <div className="ctr" style={ps}>
        <span className="p__title">{title}</span>
        <br />
        <span className="p__content">{content}</span>
      </div>
    );
  } 