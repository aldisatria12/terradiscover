export type ResLogin = {
  Data: {
    token: string;
  }
  Msg: string;
};

export type ResRegister = {
  data: {
    id: number;
    username: string;
    email: string;
  };
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
