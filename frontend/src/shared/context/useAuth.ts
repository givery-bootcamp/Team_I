import {useContext} from "react";
import {AuthContext} from "./AuthContext.tsx";
import {IAuthContext} from "./AuthContext.types.ts";

export const useAuth = (): IAuthContext => {
    const context = useContext(AuthContext);
    if (context === undefined) {
        throw new Error("useAuth must be used within an AuthProvider");
    }
    return context;
}
