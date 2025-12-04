import { conversations, currentUserId } from "@/constants/dummyData";
import InputMessage from "./InputMessage";
import RecepientProfile from "./RecepientProfile";
import ChatBubble from "../ChatBubble";
import clsx from "clsx";
import { useEffect, useRef } from "react";

export default function ChatContent({ id }: { id: string }) {
  const selectedConv = conversations.filter((val) =>
    val.participants.includes(id)
  )[0];

  const bottomRef = useRef<HTMLDivElement>(null);
  const scrollToBottom = () => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  };
  useEffect(() => {
    scrollToBottom();
  }, [id]);

  return (
    <article className="flex flex-col w-full h-full">
      {!id ? (
        <div className="flex h-full w-full items-center justify-center">
          Click on a conversation
        </div>
      ) : (
        <>
          <RecepientProfile id={id} />
          <div className="flex gap-4 flex-col bg-blue-100 px-8 py-6 rounded-2xl h-full w-full">
            <div
              className={clsx(
                "max-h-[20rem] pr-4  overflow-y-auto flex-1",
                "scrollbar scrollbar-thumb-rounded-full scrollbar-corner-transparent scrollbar-thumb-sky-700 scrollbar-track-transparent"
              )}
            >
              {selectedConv.messages.map((message, index) => (
                <ChatBubble
                  key={index}
                  subject={
                    message.senderId === currentUserId ? "sender" : "receiver"
                  }
                  text={message.text as string}
                  createdAt={message.createdAt}
                  status={message.status}
                />
              ))}
              <div ref={bottomRef} />
            </div>
            <InputMessage />
          </div>
        </>
      )}
    </article>
  );
}
