// src/shared/context/AlertContext.tsx
import React, { createContext, useContext, useState, ReactNode } from 'react';
import Alert from './Alert';

interface AlertContextType {
  alertVisible: boolean;
  alertMessage: string;
  showAlert: (message: string) => void;
  hideAlert: () => void;
}

interface AlertProviderProps {
    children: ReactNode;
}

const AlertContext = createContext<AlertContextType | undefined>(undefined);

export const AlertProvider: React.FC<AlertProviderProps> = ({ children }) => {
  const [alertVisible, setAlertVisible] = useState(false);
  const [alertMessage, setAlertMessage] = useState('');

  const showAlert = (message: string) => {
    setAlertMessage(message);
    setAlertVisible(true);
  };

  const hideAlert = () => {
    setAlertVisible(false);
    setAlertMessage('');
  };

  return (
    <AlertContext.Provider value={{ alertVisible, alertMessage, showAlert, hideAlert }}>
      {children}
    </AlertContext.Provider>
  );
};

export const useAlert = () => {
  const context = useContext(AlertContext);
  if (!context) {
    throw new Error('useAlert must be used within an AlertProvider');
  }
  return context;
};

export const AlertHandler: React.FC= () => {
    const { alertVisible, alertMessage, hideAlert} = useAlert();
    return (
        <>
            {alertVisible && <Alert message={alertMessage} onClose={hideAlert} />}
        </>
    );
}
