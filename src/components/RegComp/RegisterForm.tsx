import React from "react";
import Input from "../Input";
import Button from "../Button";

export default function RegisterForm() {
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
        label="Nickname"
        type="text"
        name="nickname"
        id="nickname"
        placeholder="Enter your nickname for display"
      />

      <Input
        label="Password"
        type="password"
        name="password"
        id="password"
        placeholder="Enter your password"
      />

      <Button>Register</Button>
    </form>
  );
}
