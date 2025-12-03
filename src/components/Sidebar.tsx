import clsx from "clsx";
import { Contact, MessageSquareText } from "lucide-react";
import { useRouter } from "next/router";

const sideBarItem = [
  {
    name: "Contact",
    icon: Contact,
    path: "/contacts",
  },
  {
    name: "Message",
    icon: MessageSquareText,
    path: "/conversations",
  },
];

export default function Sidebar() {
  const path = `/${useRouter().pathname.split("/").at(1)}`;

  return (
    <div id="sidebar" className="px-2 py-4 flex">
      <aside className="px-2 py-8 bg-slate-800 text-white rounded-2xl h-full flex-1">
        <ul className="flex flex-col justify-center gap-8">
          {sideBarItem.map((item, index) => (
            <li
              key={index}
              className={clsx(
                "flex gap-4 items-center text-slate-400 p-2 rounded-md",
                "hover:text-white hover:bg-blue-300 hover:cursor-pointer transition-all duration-200",
                path === `${item.path}` && "bg-blue-300 text-white"
              )}
            >
              <item.icon />
            </li>
          ))}
        </ul>
      </aside>
    </div>
  );
}
