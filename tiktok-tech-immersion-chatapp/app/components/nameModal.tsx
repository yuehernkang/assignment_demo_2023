import { Input, Modal } from "antd";
import React from "react";

interface NameModalProps {
  isModalOpen: boolean;
  setIsModalOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

const NameModal: React.FC<NameModalProps> = ({
  isModalOpen,
  setIsModalOpen,
}) => {
  const showModal = () => {
    setIsModalOpen(true);
  };

  const handleOk = () => {
    setIsModalOpen(false);
  };

  const handleCancel = () => {
    setIsModalOpen(false);
  };

  return (
    <>
      <Modal
        title="Enter your username!"
        open={isModalOpen}
        onOk={handleOk}
        onCancel={handleCancel}
        maskClosable={false}
        closable={false}
      >
        <Input placeholder="Username" />
      </Modal>
    </>
  );
};

export default NameModal;
