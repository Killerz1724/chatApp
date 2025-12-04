import LoginForm from "@/components/LoginComp/LoginForm";
import React from "react";

export default function LoginPage() {
  return (
    <section className="flex flex-col gap-8 items-center min-w-[30rem] bg-slate-50 rounded-xl px-24 py-12 shadow-lg">
      <div className="space-y-2">
        <h2 className="text-2xl font-black">
          Welcome to <span className="text-blue-400">Chatty</span>
        </h2>
        <p className="text-center text-sm font-semibold">
          Sign In to your chat world!
        </p>
      </div>
      <LoginForm />
    </section>
  );
}
