import React from "react";
import { Popover, PopoverContent, PopoverTrigger } from "../LibComp/Popover";
import SidebarProfile from "./SidebarProfile";
import clsx from "clsx";
import ProfileDialog from "./ProfileDialog";

export default function ProfilePopover({ id }: { id: string }) {
  return (
    <Popover>
      <PopoverTrigger>
        <SidebarProfile />
      </PopoverTrigger>
      <PopoverContent align="start">
        <ul className="space-y-4 text-sm">
          <li
            className={clsx(
              "font-semibold text-white",
              "hover:cursor-pointer hover:text-blue-400"
            )}
          >
            <ProfileDialog id={id}>See Profile</ProfileDialog>
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
