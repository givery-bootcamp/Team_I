import React from "react";

interface ConfirmModalProps {
  message: string | null;
  modalRef: React.RefObject<HTMLDialogElement>;
  onConfirm: () => void;
  onCancel: () => void;
};

const ConfirmModal = ({ message, modalRef, onConfirm, onCancel }: ConfirmModalProps) => {

  return (
    <dialog className="fixed inset-0 items-center justify-center backdrop-blur-sm" ref={modalRef}>  
      <div className="bg-white p-6 rounded-lg shadow-lg">
        <p className="text-sm text-gray-700 mb-4">{message}</p>
        <div className="flex justify-center">
          <button className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 focus:outline-none mr-2" onClick={onConfirm}>はい</button>
          <button className="px-4 py-2 bg-gray-300 text-gray-700 rounded hover:bg-gray-400 focus:outline-none" onClick={onCancel}>いいえ</button>
        </div>
      </div>
    </dialog>
  );
};

export default ConfirmModal;
