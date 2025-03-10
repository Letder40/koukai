import React from "react";
import { useState } from "react"

import FriendZone from "./zones/friend-zone"
import PublicServerZone from './zones/public-server';

type Element = React.JSX.Element;

const zoneMap = new Map<String, Element>([
    ["friends", <FriendZone />],
    ["publicServer", <PublicServerZone />],
]);

export default function DynamicZone() {
    const [currentComponent, setCurrentComponent] = useState("publicServer");
    return (
        <div className="w-full h-screen bg-dblack-Medium">
            {zoneMap.get(currentComponent)}
        </div>
    )
}
