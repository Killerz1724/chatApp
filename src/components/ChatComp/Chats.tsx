import React from "react";
import ChatLists from "./ChatLists";
import ChatContent from "./ChatContent";

export default function Chats({ id }: { id: string }) {
  return (
    <>
      <ChatLists />
      <ChatContent id={id} />
    </>
  );
}
