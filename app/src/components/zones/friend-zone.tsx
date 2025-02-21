import React from "react";

export default function FriendZone() {
  return (
    <div className="w-full h-full">
      <div className="w-full h-20 border-solid border-b-4 border-b-dblack-Strong grid grid-rows-[2rem_1fr]">
        <div className="w-full h-full flex gap-5 items-center text-white">
          <div className="ml-3.5 px-2 font-semibold shadow-sm shadow-dblack-Stronger">KOUKAI</div>
          <div className="hover:bg-dblack-Light hover:font-semibold px-2 cursor-pointer">Todos</div>
          <div className="hover:bg-dblack-Light hover:font-semibold px-2 cursor-pointer">Online</div>
          <div className="hover:bg-dblack-Light hover:font-semibold px-2 cursor-pointer">Solicitudes</div>
          <div className="hover:bg-dblack-Light hover:font-semibold px-2 cursor-pointer">Añadir Amigo</div>
        </div>
        <div className="w-full h-full flex items-center justify-center">
          <input placeholder="Buscar" className="h-7 w-11/12 bg-dblack-Stronger p-1 px-4 text-white rounded-lg focus:outline-none"/>
        </div>
      </div>
    </div>
  )
}
