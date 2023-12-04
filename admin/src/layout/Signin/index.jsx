import React, { useState } from 'react';
import axios from 'axios';
import { Navigate } from 'react-router-dom'
import { login } from '../../api/signin.js'
import { updateToken } from '../../api/comp_info.js'
import './index.less'

const LoginForm = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = async () => {
    try {
      // 发送登录请求，获取 JWT 令牌
      const response = await login({
        username: username,
        password: password
      });

      // 从响应中获取 JWT 令牌
      const jwtToken = response.data.token;

      // 将 JWT 令牌保存到本地存储或者其他地方（例如，Cookie）
      if (localStorage.getItem('jwtToken') !== jwtToken) {
        updateToken(jwtToken);
      }
      localStorage.setItem('jwtToken', jwtToken);

      console.log('Login successful');
      window.location.pathname = '/'
    } catch (error) {
      console.error('Login failed', error);
      alert('logo failed')
    }
  };
  // const isAuthenticated = localStorage.getItem('jwtToken') !== null


  return (
    <>
      <div className='logoCard'>
        <label>
          <input type="text" placeholder='   ' value={username} onChange={(e) => setUsername(e.target.value)} />
          <span className='ulabel'>Username</span>
        </label>
        <br />
        <label>
          <input type="password" placeholder='   ' value={password} onChange={(e) => setPassword(e.target.value)} />
          <span className='plabel'>Password</span>
        </label>
        <br />
        <button onClick={handleLogin}>Login</button>
      </div>
      {/* {isAuthenticated && <Navigate to="/" />} */}
    </>

  );
};

export default LoginForm;
