import clsx from "clsx";
import Image from "next/image";
import { useRouter } from "next/navigation";

type contactCardType = {
  id: string;
  name: string;
  avatar: string;
  text?: string;
};

export default function ContactCard({
  id,
  name,
  avatar,
  text = "",
}: contactCardType) {
  const route = useRouter();
  return (
    <div
      className={clsx(
        "flex gap-6 items-center bg-gray-200 px-4 py-2 rounded-2xl w-full group",
        "hover:bg-blue-300 hover:text-white  hover:cursor-pointer   transition-all duration-200"
      )}
      onClick={() => route.push(`/conversations/${id}`)}
    >
      <div className="w-8 h-8 rounded-full relative object-cover overflow-hidden">
        <Image src={avatar} fill alt={`${name} profile picture`} />
      </div>
      <div className="space-y-2 w-full">
        <h4 className="text-sm font-semibold">{name}</h4>
        <p className="text-xs group-hover:text-white group-hover:opacity-100 opacity-50 transition-all duration-200">
          {text}
        </p>
      </div>
    </div>
  );
}
