import React from "react"
import { createPortal } from 'react-dom'
//style 
import './index.less'
//
import { MessageApi, IMsgList } from './config'
//
import MsgShower from "./MsgShower"

let add: (option: IMsgList) => void

const Message: React.FC = () => {
  const [msgList, setMsgList] = React.useState<IMsgList[]>([])
  const remove = (L: IMsgList) => {
    const { key } = L
    setMsgList((pre) => pre.filter((item: IMsgList) => key !== item.key))
  }

  add = (option: IMsgList) => {
    console.log(option);
    setMsgList((pre) => {
      const obj = [...pre, option];
      setTimeout(() => {
        remove(option);
      }, 3000);
      return obj;
    });
  };

  React.useEffect(() => {
    if (msgList.length > 10) {
      msgList.shift();
    }
  }, [msgList]);

  return (
    <>
      {msgList.map(({ text, key, type }) => (
        <MsgShower key={key} type={type} text={text} />
      ))}
    </>
  );
}

const getId = (): string => {
  return ((Math.random() * 1000).toFixed());
}

const createMessage = () => {
  let el = document.getElementById('#message-wrap');
  if (!el) {
    el = document.createElement('div');
    el.className = 'message-wrap';
    el.id = 'message-wrap';
    document.body.append(el);
  }
  createPortal(<Message />, el);
};

createMessage();

const $message: MessageApi = {
  info: (text) => {
    add({
      text,
      key: getId(),
      type: 'info'
    })
  },
  success: (text) => {
    add({
      text,
      key: getId(),
      type: 'success'
    })
  },
  warning: (text) => {
    add({
      text,
      key: getId(),
      type: 'warning'
    })
  },
  error: (text) => {
    add({
      text,
      key: getId(),
      type: 'error'
    })
  }
}

export default $message

