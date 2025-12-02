import clsx from "clsx";
import React from "react";

export default function SearchBar() {
  return (
    <div
      className={clsx(
        "py-3 px-2 flex items-center gap-2 border border-gray-300 rounded-2xl",
        " focus-within:border-blue-400 transition-all duration-200"
      )}
    >
      <input
        id="search-contact"
        placeholder="Search"
        type="text"
        className={clsx("opacity-50 text-sm outline-none", "focus:opacity-100")}
      />
    </div>
  );
}
