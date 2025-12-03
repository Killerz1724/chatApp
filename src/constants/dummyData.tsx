export const users = [
  {
    id: "u1",
    name: "Alice Tan",
    avatar: "https://i.pravatar.cc/150?img=1",
  },
  {
    id: "u2",
    name: "Izzaz gooner",
    avatar: "https://i.pravatar.cc/150?img=5",
  },
  {
    id: "u3",
    name: "Paat pdf",
    avatar: "https://i.pravatar.cc/150?img=10",
  },
];

// Main logged-in user
export const currentUserId = "u1";

export const conversations = [
  {
    id: "c1",
    participants: ["u1", "u2"],
    lastMessageAt: "2024-01-21T15:32:10Z",
    messages: [
      {
        id: "m1",
        senderId: "u2",
        text: "Hey Alice! How's your day?",
        createdAt: "2024-01-21T14:00:00Z",
        status: "read", // sent | delivered | read
      },
      {
        id: "m2",
        senderId: "u1",
        text: "Pretty good! Just working on a project 😄",
        createdAt: "2024-01-21T14:01:10Z",
        status: "read",
      },
      {
        id: "m3",
        senderId: "u2",
        text: "Nice! Need help with anything?",
        createdAt: "2024-01-21T14:05:25Z",
        status: "read",
      },
      {
        id: "m4",
        senderId: "u1",
        text: "Nope, all good! Thanks!",
        createdAt: "2024-01-21T15:32:10Z",
        status: "delivered",
      },
      {
        id: "m3",
        senderId: "u2",
        text: "Nice! Need help with anything?",
        createdAt: "2024-01-21T14:05:25Z",
        status: "read",
      },
      {
        id: "m4",
        senderId: "u1",
        text: "Nope, all good! Thanks!",
        createdAt: "2024-01-21T15:32:10Z",
        status: "delivered",
      },
      {
        id: "m3",
        senderId: "u2",
        text: "Nice! Need help with anything?",
        createdAt: "2024-01-21T14:05:25Z",
        status: "read",
      },
      {
        id: "m4",
        senderId: "u1",
        text: "Nope, all good! Thanks!",
        createdAt: "2024-01-21T15:32:10Z",
        status: "delivered",
      },
    ],
  },

  {
    id: "c2",
    participants: ["u1", "u3"],
    lastMessageAt: "2024-01-20T18:45:12Z",
    messages: [
      {
        id: "m5",
        senderId: "u1",
        text: "Did you send the files?",
        createdAt: "2024-01-20T18:20:00Z",
        status: "read",
      },
      {
        id: "m6",
        senderId: "u3",
        text: "Yup, sending now!",
        createdAt: "2024-01-20T18:21:10Z",
        status: "read",
      },
      {
        id: "m7",
        senderId: "u3",
        attachment: {
          type: "file",
          name: "report.pdf",
          size: "1.2MB",
          url: "https://example.com/report.pdf",
        },
        createdAt: "2024-01-20T18:22:40Z",
        status: "read",
      },
      {
        id: "m8",
        senderId: "u1",
        text: "Got it, thanks!",
        createdAt: "2024-01-20T18:45:12Z",
        status: "delivered",
      },
    ],
  },
];

export const typingStatus = [
  // userId + conversationId
  {
    conversationId: "c1",
    userId: "u2",
    isTyping: true,
  },
];
