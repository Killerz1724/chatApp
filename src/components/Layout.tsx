import React from "react";
import Sidebar from "./Sidebar";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <section className="flex bg-slate-50  min-h-[35rem] min-w-[70rem] shadow-xl rounded-2xl">
      <Sidebar />
      <div className="flex pt-8 pb-4 px-4 w-full">{children}</div>
    </section>
  );
}
