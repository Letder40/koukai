import React from 'react';
import PersonAddIcon from '@mui/icons-material/PersonAdd';
import InsertCommentSharpIcon from '@mui/icons-material/InsertCommentSharp';
import LanguageIcon from '@mui/icons-material/Language';

function RoundButton({content: Content}: {content: React.ComponentType<any> }) {
  return (
    <button className="h-10 w-10 bg-white rounded-full"><Content color="action"/></button>
  )
}

function ActionButtons() {
  return (
    <div className="h-20 w-full bg-dblue flex justify-center items-center gap-5">
      <RoundButton content={PersonAddIcon}/>
      <RoundButton content={InsertCommentSharpIcon}/>
    </div>
  )
}

function Server() {
  return (
    <div className="flex justify-center items-center mt-5 w-full">
      <button id="ServerIcon" className="peer h-14 w-14 bg-dblack-Medium rounded-full transition-all delay-150 duration-100 ease-in-out hover:rounded-2xl hover:bg-dblue"><LanguageIcon htmlColor='white'/></button>
      <div id="ServerMarker" className="peer peer-hover:visible w-1 bg-white h-5 absolute left-0 rounded-lg invisible"></div>
    </div>
  )
}

function Servers() {
  return (
    <div className="flex flex-col items-center bg-dblack-Strong h-full w-20 overflow-scroll">
      <Server/>
    </div>
  )
}

export default function Navbar(): React.JSX.Element {
  return (
    <div className="bg-dblack-Light h-screen w-screen md:w-80 shadow-lg shadow-dblack-Light grid grid-cols-[5rem_1fr]">
      <Servers/>
      <ActionButtons/>
    </div>
  )
}
