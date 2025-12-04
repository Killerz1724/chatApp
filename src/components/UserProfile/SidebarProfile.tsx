import { currentUserId, users } from "@/constants/dummyData";
import clsx from "clsx";
import Image from "next/image";

export default function SidebarProfile() {
  const currentUser = users.find((user) => user.id === currentUserId);
  return (
    <div
      className={clsx(
        "w-8 h-8 rounded-full relative overflow-hidden",
        "hover:cursor-pointer hover:opacity-80 transition-all duration-200"
      )}
    >
      <Image src={currentUser?.avatar as string} fill alt={"dummy"} />
    </div>
  );
}
