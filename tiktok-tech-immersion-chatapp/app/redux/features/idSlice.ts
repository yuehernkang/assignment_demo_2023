import { createSlice } from "@reduxjs/toolkit";

export interface IDState {
  value: string;
}

const initialState: IDState = {
  value: "",
};

export const idSlice = createSlice({
  name: "id",
  initialState,
  reducers: {
    setId: (state, action) => {
      state.value = action.payload;
    },
  },
});

export const { setId } = idSlice.actions;

export default idSlice.reducer;
