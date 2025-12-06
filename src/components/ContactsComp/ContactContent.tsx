import { users } from "@/constants/dummyData";
import Image from "next/image";
import React from "react";
import Button from "../Button";

export default function ContactContent({ id }: { id?: string }) {
  const selectedUser = users.find((user) => user.id === id);
  return (
    <div className="flex items-center justify-center h-full w-full">
      {selectedUser ? (
        <div className="flex flex-col gap-4 items-center">
          <div className="w-24 h-24 relative rounded-full overflow-hidden">
            <Image
              src={
                selectedUser
                  ? selectedUser.avatar
                  : "/images/defaultProfPic.webp"
              }
              alt={"dummy"}
              fill
            />
          </div>
          <h4 className="font-bold text-lg">{selectedUser?.id}</h4>
          <Button>+ Add as contact</Button>
        </div>
      ) : (
        <h4>Click on a contact</h4>
      )}
    </div>
  );
}
