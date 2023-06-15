import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { ChatMessage } from "../types/Message";

// Define a service using a base URL and expected endpoints
export const messagesApi = createApi({
  reducerPath: "messagesApi",
  baseQuery: fetchBaseQuery({
    baseUrl: "http://localhost:8080/api",
  }),
  endpoints: (builder) => ({
    getMessagesByRoomID: builder.query<ChatMessage[], string>({
      query: (name) => `pull?chat=${name}`,
    }),
  }),
});

// Export hooks for usage in functional components, which are
// auto-generated based on the defined endpoints
export const { useGetMessagesByRoomIDQuery } = messagesApi;
