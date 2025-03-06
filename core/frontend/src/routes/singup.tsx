import React, { useState } from 'react';
import '../css/style.css';
import { Link } from 'react-router-dom';

const Singup = () => {
    const [singupData, setSingupData] = useState({
        username: '',
        email: '',
        password: '',
        role: '1',
    });

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const {name, value} = event.target;
        setSingupData ({
            ...singupData,
            [name]: value
        })
    };

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        try {
            const response = await fetch("/api/singup", {
                method: 'POST',
                headers: {
                    'Content-Type': 'Application/json'
                },
                body: JSON.stringify(singupData),
            });

            if (!response.ok) {
                throw new Error('singup error');
            }

            window.location.replace("/login")

        } catch (error) {
            console.log(error) 
        }
    }

    return (
    <React.StrictMode>
      <div className='w-screen h-screen flex items-center justify-center bg-dblue'>
        <form onSubmit={handleSubmit} className='w-5/6 max-w-screen-sm h-[600px] grid grid-rows-[180px_1fr_1fr_1fr_1fr] text-white bg-dblack-Medium'>
          <div className='self-center justify-self-center flex items-center flex-col'>
            <h1 className='self-center justify-self-center text-3xl'>Welcome to koukai</h1>
            <p>Singup, you won't regret it</p>
          </div>
          <div className='self-center justify-self-center w-11/12'>
            <label htmlFor='username'>Username</label>
            <input onChange={handleChange} name='username' type='text' className='w-full h-10 px-3 bg-dblack-Stronger' placeholder='Username'/>
          </div>
          <div className='self-center justify-self-center w-11/12'>
            <label htmlFor='email'>Email</label>
            <input onChange={handleChange} name='email' type='text' className='w-full h-10 px-3 bg-dblack-Stronger' placeholder='Email'/>
          </div>
          <div className='self-center justify-self-center w-11/12'>
            <label htmlFor='password'>Password</label>
            <input onChange={handleChange} name='password' type='password' className='w-full h-10 px-3 bg-dblack-Stronger' placeholder='Password'/>
          </div>
          <div className='self-center justify-self-center w-11/12'>
            <input type='submit' value='Singup' className='self-start justify-self-center w-full h-10 mb-1 bg-dblueDeep'/>
            <span className='text-[#a5a5a5]'>Do you have an account?</span>&nbsp;
            <Link to='/login' className='self-center justify-self-start text-[#00bbff]'>Login</Link>
          </div>
        </form>
      </div>
    </React.StrictMode>
  )
}

export default Singup;
