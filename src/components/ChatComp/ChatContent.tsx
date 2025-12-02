import React from "react";
import InputMessage from "./InputMessage";
import RecepientProfile from "./RecepientProfile";

export default function ChatContent() {
  return (
    <article className="flex flex-col w-full">
      <RecepientProfile />
      <div className="flex flex-col bg-blue-100 px-8 py-6 rounded-2xl h-full w-full">
        <div className="flex-1">Chat Content</div>
        <InputMessage />
      </div>
    </article>
  );
}
