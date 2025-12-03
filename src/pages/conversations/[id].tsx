import Chats from "@/components/ChatComp/Chats";
import Layout from "@/components/Layout";
import { useRouter } from "next/router";
import React from "react";

export default function ConversationId() {
  const router = useRouter();
  const { id } = router.query;
  return (
    <Layout>
      <Chats id={id as string} />
    </Layout>
  );
}
