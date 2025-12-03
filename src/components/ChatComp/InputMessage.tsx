import clsx from "clsx";

export default function InputMessage() {
  return (
    <div className="flex gap-8 w-full">
      <form className="border-gray-300 bg-white border-2 rounded-lg px-4 py-2 w-full flex gap-4">
        <input
          type="text"
          placeholder="type your message here"
          className="bg-none flex-1 outline-none text-sm"
        />
        <button
          className={clsx(
            "px-4 py-1 bg-gray-100 opacity-50 rounded-xl",
            "hover:bg-blue-400 hover:opacity-100 hover:text-white transition-all duration-200"
          )}
          type="submit"
        >
          -&gt;
        </button>
      </form>

      <button className="px-4 py-2 rounded-lg bg-blue-600 text-center text-white  ">
        +
      </button>
    </div>
  );
}
