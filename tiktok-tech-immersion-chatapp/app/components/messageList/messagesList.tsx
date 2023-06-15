"use client";

import { useGetMessagesByRoomIDQuery } from "@/app/services/messages";
import { ChatMessage } from "@/app/types/Message";
import { Avatar, List } from "antd";
import { FC, useEffect } from "react";

const data: ChatMessage[] = [
  {
    chat: "Hello, how are you?",
    sender: "Jack",
  },
  {
    chat: "I'm good, thanks! How about you?",
    sender: "Zack",
  },
  {
    chat: "I'm doing great. Did you watch the game last night?",
    sender: "Jack",
  },
  {
    chat: "Yes, it was amazing! The team played really well.",
    sender: "Zack",
  },
];

interface MessageItemProps {
  item: ChatMessage;
}

const MessageList = () => {
  const { data, error, isLoading } = useGetMessagesByRoomIDQuery("tom:jane");
  const currentUser = "Jack";

  useEffect(() => {
    if (data) console.log(data);
    if (error) console.log(error);
  }, [data, error]);

  return (
    <List
      itemLayout="horizontal"
      dataSource={data}
      renderItem={(item, index) => (
        <List.Item>
          <List.Item.Meta
            title={<a href="https://ant.design"></a>}
            description={
              item.sender == currentUser ? (
                <MessageItemAuthor item={item} />
              ) : (
                <MessageItemSender item={item} />
              )
            }
          />
        </List.Item>
      )}
    />
  );
};

const MessageItemAuthor: FC<MessageItemProps> = ({ item }) => {
  return (
    <div className="flex flex-row">
      <div className="rounded-2xl bg-slate-300 p-3">{item.chat}</div>
      <Avatar src={`https://xsgames.co/randomusers/avatar.php?g=pixel&key=1`} />
    </div>
  );
};

const MessageItemSender: FC<MessageItemProps> = ({ item }) => {
  return (
    <div className="flex flex-row">
      <Avatar src={`https://xsgames.co/randomusers/avatar.php?g=pixel&key=2`} />
      <div className="rounded-2xl bg-slate-300 p-3">{item.chat}</div>
    </div>
  );
};

export default MessageList;
