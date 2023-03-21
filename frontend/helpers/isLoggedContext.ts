import { createContext } from "react";

interface IsLoggedContextType {
  isLogged: boolean;
  setIsLogged: (isLogged: boolean) => void;
}

export const IsLoggedContext = createContext<IsLoggedContextType>({
  isLogged: false,
  setIsLogged: () => {},
});
