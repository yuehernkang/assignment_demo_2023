"use client";

import { useEffect, useState } from "react";
import { Provider } from "react-redux";
import ChatList from "./components/chatList";
import MessageList from "./components/messageList/messagesList";
import ChatInput from "./components/messageList/senderArea";
import NameModal from "./components/nameModal";
import { store } from "./redux/stores/store";

export default function Home() {
  const [isModalOpen, setIsModalOpen] = useState(false);
  useEffect(() => {
    const user = localStorage.getItem("username");
    if (user == undefined) setIsModalOpen(true);
  }, []);
  return (
    <Provider store={store}>
      <NameModal isModalOpen={isModalOpen} setIsModalOpen={setIsModalOpen} />
      <main className="flex flex-row h-screen">
        <div className="flex-none w-1/3 border-r pr-4">
          <ChatList />
        </div>
        <div className="flex-grow pl-4 border-l pr-2">
          <div className="flex flex-grow flex-col h-full">
            <MessageList />
            <div className="flex-none">
              <ChatInput onSend={() => {}} />
            </div>
          </div>
        </div>
      </main>
    </Provider>
  );
}
