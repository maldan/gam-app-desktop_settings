import Axios from 'axios';

const API_URL = process.env.VUE_APP_API_URL || `${window.location.origin}/api`;

export const RestApi = {
  application: {
    async getList(): Promise<any[]> {
      return (await Axios.get(`${API_URL}/application/list`)).data.response;
    },
  },
  main: {
    async uploadBackground(id: number, file: File) {
      const formData = new FormData();
      formData.append(`image_0`, file, file.name);
      return (
        await Axios.post(`${API_URL}/main/background?id=${id}`, formData, {
          headers: {
            'Content-Type': 'multipart/form-data',
          },
        })
      ).data.response;
    },
  },
};
