import Image from "next/image";

export default function RecepientProfile() {
  return (
    <div className="flex items-center  justify-between pt-2 pb-8 px-4">
      <div className="flex items-center gap-4">
        <div className="w-10 h-10 border border-gray-400 rounded-full relative object-cover overflow-hidden">
          <Image src={"/images/defaultProfPic.webp"} alt={"dummy"} fill />
        </div>
        <h4 className="font-bold">John Doe</h4>
      </div>
    </div>
  );
}
