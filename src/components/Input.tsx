import React from "react";

type InputType = {
  label: string;
} & React.InputHTMLAttributes<HTMLInputElement>;

export default function Input({ label, ...props }: InputType) {
  return (
    <div className="flex flex-col gap-2">
      <label className="text-sm font-semibold" htmlFor={props.id}>
        {label}
      </label>{" "}
      <input className="text-xs border rounded-lg px-4 py-2" {...props} />
    </div>
  );
}
