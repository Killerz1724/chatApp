import React from "react";
import { Popover, PopoverContent, PopoverTrigger } from "../LibComp/Popover";
import SidebarProfile from "./SidebarProfile";
import clsx from "clsx";

export default function ProfilePopover() {
  return (
    <Popover>
      <PopoverTrigger>
        <SidebarProfile />
      </PopoverTrigger>
      <PopoverContent align="start">
        <ul className="space-y-4 text-sm">
          <li
            className={clsx(
              "font-semibold",
              "hover:cursor-pointer hover:text-blue-400"
            )}
          >
            See Profile
          </li>
          <li
            className={clsx(
              "text-red-600 font-bold",
              "hover:cursor-pointer hover:text-red-400"
            )}
          >
            Sign Out
          </li>
        </ul>
      </PopoverContent>
    </Popover>
  );
}
