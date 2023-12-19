import React from "react"
import { createRoot } from 'react-dom/client'
import { v1 } from 'uuid'
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
    setMsgList((pre) => {
      const obj = [...pre, option];
      setTimeout(() => {
        remove(option);
      }, 2000);
      return obj;
    });
  };

  React.useEffect(() => {
    if (msgList.length > 2) {
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
  return v1();
}

const createMessage = () => {
  let el = document.getElementById('#message-wrap');
  if (!el) {
    el = document.createElement('div');
    el.className = 'message-wrap';
    el.id = 'message-wrap';
    document.body.append(el);
  }
  const root = createRoot(el);
  root.render(<Message />);
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

