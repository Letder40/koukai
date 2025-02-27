import React from 'react';
import './css/style.css';

import Navbar from "./components/navbar";
import DynamicZone from "./components/dynamic-zone";

const Main = () => (
  <React.StrictMode>
    <div className='flex'>
      <Navbar/> 
      <DynamicZone/>
    </div>
  </React.StrictMode>
);

export default Main;
