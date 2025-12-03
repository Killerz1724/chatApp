import InputMessage from "./InputMessage";
import RecepientProfile from "./RecepientProfile";

export default function ChatContent({ id }: { id: string }) {
  return (
    <article className="flex flex-col w-full">
      {!id ? (
        <div className="flex h-full w-full items-center justify-center">
          Click on a conversation
        </div>
      ) : (
        <>
          <RecepientProfile id={id} />
          <div className="flex flex-col bg-blue-100 px-8 py-6 rounded-2xl h-full w-full">
            <div className="flex-1">Chat Content</div>
            <InputMessage />
          </div>
        </>
      )}
    </article>
  );
}
