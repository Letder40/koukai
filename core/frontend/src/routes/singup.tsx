import React from 'react';
import '../css/style.css';
import { Link } from 'react-router-dom';

const Singup = () => (
  <React.StrictMode>
    <div className='w-screen h-screen flex items-center justify-center bg-dblue'>
      <form className='w-5/6 max-w-screen-sm h-[600px] grid grid-rows-[180px_1fr_1fr_1fr_1fr] text-white bg-dblack-Medium'>
        <div className='self-center justify-self-center flex items-center flex-col'>
            <h1 className='self-center justify-self-center text-3xl'>Welcome to koukai</h1>
            <p>Singup, you won't regret it</p>
        </div>
        <div className='self-center justify-self-center w-11/12'>
            <label htmlFor='identifier'>Email</label>
            <input name='username' type='text' className='w-full h-10 px-3 bg-dblack-Stronger' placeholder='Username'/>
        </div>
        <div className='self-center justify-self-center w-11/12'>
            <label htmlFor='identifier'>Username</label>
            <input name='identifier' type='text' className='w-full h-10 px-3 bg-dblack-Stronger' placeholder='Email'/>
        </div>
        <div className='self-center justify-self-center w-11/12'>
            <label htmlFor='identifier'>Password</label>
            <input name='password' type='password' className='w-full h-10 px-3 bg-dblack-Stronger' placeholder='Password'/>
        </div>
        <div className='self-center justify-self-center w-11/12'>
            <input type='submit' value='Singup' className='self-start justify-self-center w-full h-10 mb-1 bg-dblueDeep'/>
            <span className='text-[#a5a5a5]'>Do you have an account?</span>&nbsp;
            <Link to='/login' className='self-center justify-self-start text-[#00bbff]'>Login</Link>
        </div>
      </form>
    </div>
  </React.StrictMode>
);

export default Singup;
