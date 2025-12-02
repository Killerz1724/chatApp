import React from "react";

export default function Sidebar() {
  return (
    <div id="sidebar" className="px-2 py-4 flex">
      <aside className="px-4 py-8 bg-slate-800 text-white rounded-2xl h-full flex-1">
        <ul className="flex flex-col gap-12">
          <li>Sidebar</li>
          <li>Sidebar</li>
          <li>Sidebar</li>
          <li>Sidebar</li>
        </ul>
      </aside>
    </div>
  );
}
