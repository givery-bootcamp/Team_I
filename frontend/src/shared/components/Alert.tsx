import React from 'react';

interface AlertProps {
  message: string;
  onClose: () => void;
}

const Alert: React.FC<AlertProps> = ({ message, onClose }) => {
  return (
    <div className="fixed bottom-4 right-0 transform bg-blue-500 text-white mx-3 px-6 py-3 rounded shadow-lg">
      <div className="flex justify-between items-center">
        <span>{message}</span>
        <button onClick={onClose} className="text-white ml-4">
          &times;
        </button>
      </div>
    </div>
  );
};

export default Alert;