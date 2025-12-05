import React from "react";
import { Dialog, DialogContent, DialogTrigger } from "../LibComp/Dialog";
import { users } from "@/constants/dummyData";
import Image from "next/image";

export default function ProfileDialog({
  id,
  children,
}: {
  id: string;
  children: React.ReactNode;
}) {
  const selectedUser = users.find((val) => val.id === id);
  return (
    <Dialog>
      <DialogTrigger>{children}</DialogTrigger>
      <DialogContent>
        <div className="flex flex-col gap-6 items-center">
          <div className="w-24 h-24 relative rounded-full overflow-hidden">
            <Image src={selectedUser?.avatar as string} fill alt={"dummy"} />
          </div>
          <div className="space-y-2">
            <h4 className="font-bold text-lg">{selectedUser?.name}</h4>
            <p className="text-sm text-center">Bio</p>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  );
}
