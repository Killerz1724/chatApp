import { conversations, currentUserId, users } from "@/constants/dummyData";
import ContactCard from "./ContactCard";

export default function ChatLists() {
  const userConv = conversations.filter((conv) =>
    conv.participants.includes(currentUserId)
  );
  const contacts = userConv.map(
    (conv) => conv.participants.filter((id) => id !== currentUserId)[0]
  );

  return (
    <article className="flex flex-col gap-8 pl-2 pr-4 w-full max-w-[15rem]">
      <div>Search Bar</div>
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
                  name={user.name}
                  avatar={user.avatar}
                  text={lastMesage?.text}
                />
              </li>
            );
          })}
      </ul>
    </article>
  );
}
