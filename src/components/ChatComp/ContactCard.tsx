import Image from "next/image";

type contactCardType = {
  name: string;
  avatar: string;
};

export default function ContactCard({ name, avatar }: contactCardType) {
  return (
    <div className="flex gap-2 items-center">
      <div className="w-8 h-8 rounded-full relative object-cover overflow-hidden">
        <Image src={avatar} fill alt={`${name} profile picture`} />
      </div>
      <h4 className="text-sm font-semibold">{name}</h4>
    </div>
  );
}
