import React from "react";
import Sidebar from "./Sidebar";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <section className="flex bg-slate-50 shadow-xl rounded-2xl">
      <Sidebar />
      <div className="flex py-10 px-4 w-full">{children}</div>
    </section>
  );
}
