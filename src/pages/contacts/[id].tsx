import ChatLists from "@/components/ChatComp/ChatLists";
import ContactContent from "@/components/ContactsComp/ContactContent";
import Layout from "@/components/Layout";
import { useRouter } from "next/router";

export default function ContactsId() {
  const route = useRouter();
  const id = route.query.id as string;
  return (
    <Layout>
      <ChatLists />
      <ContactContent id={id} />
    </Layout>
  );
}
