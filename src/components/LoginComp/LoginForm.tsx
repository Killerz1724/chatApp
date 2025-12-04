import React from "react";
import Input from "../Input";
import Button from "../Button";

export default function LoginForm() {
  return (
    <form className="flex flex-col gap-4 w-full">
      <Input
        label="Email"
        type="email"
        name="email"
        id="email"
        placeholder="Enter your email"
      />

      <Input
        label="Password"
        type="password"
        name="password"
        id="password"
        placeholder="Enter your password"
      />

      <Button>Login</Button>
    </form>
  );
}
