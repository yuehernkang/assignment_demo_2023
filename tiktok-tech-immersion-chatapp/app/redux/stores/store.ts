import { messagesApi } from "@/app/services/messages";
import { configureStore } from "@reduxjs/toolkit";
import idReducer from "../features/idSlice";

export const store = configureStore({
  reducer: { name: idReducer, [messagesApi.reducerPath]: messagesApi.reducer },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(messagesApi.middleware),
});

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;
