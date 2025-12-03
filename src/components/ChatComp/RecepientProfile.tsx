import { users } from "@/constants/dummyData";
import Image from "next/image";

export default function RecepientProfile({ id }: { id: string }) {
  const user = users.find((user) => user.id === id);
  return (
    <div className="flex items-center  justify-between pt-2 pb-8 px-4">
      <div className="flex items-center gap-4">
        <div className="w-10 h-10 border border-gray-400 rounded-full relative object-cover overflow-hidden">
          <Image
            src={user ? user.avatar : "/images/defaultProfPic.webp"}
            alt={"dummy"}
            fill
          />
        </div>
        <h4 className="font-bold">{user?.id}</h4>
      </div>
    </div>
  );
}
