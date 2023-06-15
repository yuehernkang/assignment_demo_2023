"use client";
import { Button, Input } from "antd";
import React, { useState } from "react";

interface ChatInputProps {
  onSend: (message: string) => void;
}

const ChatInput: React.FC<ChatInputProps> = ({ onSend }) => {
  const [message, setMessage] = useState("");

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setMessage(e.target.value);
  };

  const handleSendClick = () => {
    onSend(message);
    setMessage("");
  };
  const isSendButtonDisabled = message === "";

  return (
    <div className="flex flex-row gap-2">
      <Input
        value={message}
        onChange={handleInputChange}
        placeholder="Type a message"
      />
      <Button onClick={handleSendClick} disabled={isSendButtonDisabled}>
        Send
      </Button>
    </div>
  );
};

export default ChatInput;
