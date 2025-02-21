import React from 'react';
import ReactDOM from 'react-dom/client';
import './css/style.css';

import Navbar from "./components/navbar";
import DynamicZone from "./components/dynamic-zone";

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

root.render(
  <React.StrictMode>
    <div className='flex'>
      <Navbar/> 
      <DynamicZone/>
    </div>
  </React.StrictMode>
);
