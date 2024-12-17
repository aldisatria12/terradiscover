export type ResLogin = {
  Data: {
    token: string;
  };
  Msg: string;
};

export type ResRegister = {
  Data: {
    id: number;
    username: string;
    email: string;
  };
  Msg: string;
};

export type ResError = {
  message: string;
  errors: [
    {
      field: string;
      message: string;
    }
  ];
};
