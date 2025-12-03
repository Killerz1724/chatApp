import clsx from "clsx";

export default function ChatBubble({
  subject,
  text,
  createdAt,
  status,
}: {
  subject: string;
  text: string;
  createdAt: string;
  status: string;
}) {
  const date = new Date(createdAt).toLocaleTimeString("en-US", {
    hour: "numeric",
    minute: "numeric",
  });
  return (
    <div
      className={clsx(
        "flex w-full mb-2",
        subject === "sender" && "justify-end"
      )}
    >
      <div
        className={clsx(
          "flex flex-col gap-1 py-1 px-2 rounded-lg max-w-[30rem] ",
          subject === "sender" && "bg-blue-200 ",
          subject === "receiver" && "bg-gray-200 "
        )}
      >
        <p className="text-sm font-semibold">{text}</p>
        <div className="flex gap-1 items-center justify-end">
          <span className="text-xs opacity-50">{date}</span>
          <span className="text-xs opacity-50">{status}</span>
        </div>
      </div>
    </div>
  );
}
