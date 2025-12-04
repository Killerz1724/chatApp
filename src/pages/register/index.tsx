import RegisterForm from "@/components/RegComp/RegisterForm";
import Link from "next/link";
import React from "react";

export default function index() {
  return (
    <section className="flex flex-col gap-8 items-center min-w-[30rem] bg-slate-50 rounded-xl px-24 py-12 shadow-lg">
      <div className="space-y-2">
        <h2 className="text-2xl font-black">
          Welcome to <span className="text-blue-400">Chatty</span>
        </h2>
        <p className="text-center text-sm font-semibold">
          Sign Up to access your chat world!
        </p>
      </div>
      <div className="space-y-2">
        <RegisterForm />
        <p className="text-xs">
          Already have an account?{" "}
          <Link className="text-blue-400 text-sm " href={"/login"}>
            Login Here!
          </Link>
        </p>
      </div>
    </section>
  );
}
