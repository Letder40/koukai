import React, { useState } from 'react';
import '../css/style.css';
import { Link } from 'react-router-dom';

const Login = () => {
    const [formData, setFormData] = useState({
        identifier: '',
        password: '',
    });

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const {name, value} = event.target;
        setFormData ({
            ...formData,
            [name]: value
        })
    };

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        try {
            const response = await fetch("/api/login", {
                method: 'POST',
                headers: {
                    'Content-Type': 'Application/json'
                },
                body: JSON.stringify(formData),
            });

            if (!response.ok) {
                throw new Error('Login error');
            }

            const data = await response.json()
            console.log('Response:', data);
        } catch (error) {
            console.log("error: " + error) 
        }
    }

    return(
      <React.StrictMode>
        <div className='w-screen h-screen flex items-center justify-center bg-dblue'>
          <form className='w-5/6 max-w-screen-sm h-[500px] grid grid-rows-[180px_1fr_1fr_1fr] text-white bg-dblack-Medium' onSubmit={handleSubmit}>
            <div className='self-center justify-self-center flex items-center flex-col'>
                <h1 className='self-center justify-self-center text-3xl'>Welcome to koukai</h1>
                <p>Login to start chatting, you won't regret it</p>
            </div>
            <div className='self-center justify-self-center w-11/12'>
                <label htmlFor='identifier'>Email or username</label>
                <input onChange={handleChange} name='identifier' type='text' className='w-full h-10 px-3 bg-dblack-Stronger' placeholder='Email or username'/>
            </div>
            <div className='self-center justify-self-center w-11/12'>
                <label htmlFor='identifier'>Password</label>
                <input onChange={handleChange} name='password' type='password' className='w-full h-10 px-3 bg-dblack-Stronger' placeholder='Password'/>
            </div>
            <div className='self-center justify-self-center w-11/12'>
                <input type='submit' value='Login' className='self-start justify-self-center w-full h-10 mb-1 bg-dblueDeep'/>
                <span className='text-[#a5a5a5]'>Need an account?</span>&nbsp;
                <Link to='/singup' className='self-center justify-self-start text-[#00bbff]'>Register</Link>
            </div>
          </form>
        </div>
      </React.StrictMode>
    );
}
export default Login;
