import { useRef, useState } from "react";

export const useConfirmModal = () => {
    const modalRef = useRef<HTMLDialogElement>(null);
    const [confirmMessage, setConfirmMessage] = useState<string | null>(null);
    const [resolveCallback, setResolveCallback] = useState<(value: boolean) => void | null>();

    const customConfirm = async (message: string) => {
        setConfirmMessage(message);
        modalRef.current?.showModal();
        return new Promise<boolean>((resolve) => {
            setResolveCallback(() => resolve);
        });
    }

    const onConfirm = () => {
        if (resolveCallback) resolveCallback(true);
        modalRef.current?.close();
    }

    const onCancel = () => {
        if (resolveCallback) resolveCallback(false);
        modalRef.current?.close();
    }
    
    return { modalRef, confirmMessage, onConfirm, onCancel, customConfirm};
};
