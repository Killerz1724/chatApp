import { conversations, currentUserId, users } from "@/constants/dummyData";
import ContactCard from "./ContactCard";
import SearchBar from "./SearchBar";

export default function ChatLists() {
  const userConv = conversations.filter((conv) =>
    conv.participants.includes(currentUserId)
  );
  const contacts = userConv.map(
    (conv) => conv.participants.filter((id) => id !== currentUserId)[0]
  );

  return (
    <article className="flex flex-col gap-8 pl-2 pr-4 w-full max-w-[15rem]">
      <SearchBar />
      <ul className="flex flex-col gap-4">
        {users
          .filter((user) => contacts.includes(user.id))
          .map((user) => {
            const lastMesage = userConv
              .find((conv) => conv.participants.includes(user.id))
              ?.messages.at(-1);
            return (
              <li key={user.id}>
                <ContactCard
                  id={user.id}
                  name={user.name}
                  avatar={user.avatar}
                  text={lastMesage?.text}
                  isConv
                />
              </li>
            );
          })}
      </ul>
    </article>
  );
}
