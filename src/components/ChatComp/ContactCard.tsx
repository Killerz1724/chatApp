import clsx from "clsx";
import Image from "next/image";

type contactCardType = {
  name: string;
  avatar: string;
  text?: string;
};

export default function ContactCard({
  name,
  avatar,
  text = "",
}: contactCardType) {
  return (
    <div
      className={clsx(
        "flex gap-6 items-center bg-gray-100 px-4 py-2 rounded-2xl w-full",
        "hover:bg-blue-300 hover:text-white hover:cursor-pointer   transition-all duration-200"
      )}
    >
      <div className="w-8 h-8 rounded-full relative object-cover overflow-hidden">
        <Image src={avatar} fill alt={`${name} profile picture`} />
      </div>
      <div className="space-y-2 w-full">
        <h4 className="text-sm font-semibold">{name}</h4>
        <p className="text-xs opacity-50">{text}</p>
      </div>
    </div>
  );
}
