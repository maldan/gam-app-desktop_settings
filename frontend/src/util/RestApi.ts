import Axios from 'axios';

const API_URL = process.env.VUE_APP_API_URL || `${window.location.origin}/api`;

export const RestApi = {
  application: {
    async getList(): Promise<any[]> {
      return (await Axios.get(`${API_URL}/application/list`)).data.response;
    },
  },
};
