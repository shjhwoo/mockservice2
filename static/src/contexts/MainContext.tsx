import { createContext, useContext, Dispatch, useReducer } from "react";

export const SAVE = "SAVE";
export const REFRESH = "REFRESH";

type State = {
  accessToken: string;
};

type Action =
  | { type: "SAVE"; accessToken: string }
  | { type: "REFRESH"; accessToken: string };

type accessTokenDispatch = Dispatch<Action>;

// Context 만들기
const SampleStateContext = createContext<State | null>(null);
const SampleDispatchContext = createContext<accessTokenDispatch | null>(null);

//리듀서
export const reducer = (state: State, action: Action) => {
  switch (action.type) {
    case "SAVE":
      return { ...state, accessToken: action.accessToken };
    case "REFRESH":
      return { ...state, accessToken: action.accessToken };
    default:
      return state;
  }
};

export const TokenProvider = ({ children }: { children: React.ReactNode }) => {
  const [state, dispatch] = useReducer(reducer, { accessToken: "" });
  return (
    <SampleStateContext.Provider value={state}>
      <SampleDispatchContext.Provider value={dispatch}>
        {children}
      </SampleDispatchContext.Provider>
    </SampleStateContext.Provider>
  );
};

// state 와 dispatch 를 쉽게 사용하기 위한 커스텀 Hooks
export function useSampleState() {
  const state = useContext(SampleStateContext);
  if (!state) throw new Error("Cannot find SampleProvider"); // 유효하지 않을땐 에러를 발생
  return state;
}

export function useSampleDispatch() {
  const dispatch = useContext(SampleDispatchContext);
  if (!dispatch) throw new Error("Cannot find SampleProvider"); // 유효하지 않을땐 에러를 발생
  return dispatch;
}
