import React from "react";
import { useState } from "react"

import FriendZone from "./zones/friend-zone"

type Element = React.JSX.Element;

const zoneMap = new Map<String, Element>([
  ["friends", <FriendZone/>],
]);

export default function DynamicZone() {
  const [currentComponent, setCurrentComponent] = useState("friends");
  return (
    <div className="w-full h-screen bg-dblack-Medium">
      {zoneMap.get(currentComponent)}
    </div>
  )
}
