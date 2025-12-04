import clsx from "clsx";
import React from "react";

type ButtonType = {
  children: React.ReactNode;
} & React.HTMLAttributes<HTMLButtonElement>;

export default function Button({ children, ...props }: ButtonType) {
  return (
    <button
      className={clsx(
        "font-bold px-4 py-2 mt-4 rounded-lg bg-blue-600 text-white",
        "hover:bg-blue-400 "
      )}
      {...props}
    >
      {children}
    </button>
  );
}
