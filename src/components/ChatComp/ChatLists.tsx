import { currentUserId, users } from "@/constants/dummyData";
import ContactCard from "./ContactCard";

export default function ChatLists() {
  return (
    <article className="flex flex-col gap-8 pl-2 pr-16 w-full max-w-[13rem]">
      <div>Search Bar</div>
      <ul className="flex flex-col gap-4">
        {users
          .filter((user) => user.id !== currentUserId)
          .map((user) => (
            <li key={user.id}>
              <ContactCard name={user.name} avatar={user.avatar} />
            </li>
          ))}
      </ul>
    </article>
  );
}
