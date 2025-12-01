import { Html, Head, Main, NextScript } from "next/document";

export default function Document() {
  return (
    <Html lang="en">
      <Head />
      <body>
        <main className="px-[20rem] py-[5rem] bg-gray-300 h-screen">
          <Main />
        </main>

        <NextScript />
      </body>
    </Html>
  );
}
