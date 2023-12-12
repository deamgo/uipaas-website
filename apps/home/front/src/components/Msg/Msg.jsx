const Msg = ({ type, text }) => {
    return (
      <div className={`message ${type}`}>
        <span className='icon' />
        <span>{text}</span>
      </div>
    );
  };
  
  export default Msg;
  