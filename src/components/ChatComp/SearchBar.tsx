import clsx from "clsx";
import { Search } from "lucide-react";
import React from "react";

export default function SearchBar() {
  return (
    <div
      className={clsx(
        "py-3 px-2 flex items-center gap-2 border border-gray-300 rounded-2xl",
        " focus-within:border-blue-400 transition-all duration-200"
      )}
    >
      <Search className="text-gray-400" />
      <input
        id="search-contact"
        placeholder="Search"
        type="text"
        className={clsx(
          "opacity-50 text-sm outline-none bg-transparent",
          "focus:opacity-100"
        )}
      />
    </div>
  );
}
